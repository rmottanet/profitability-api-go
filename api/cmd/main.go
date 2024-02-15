package main

import (
	"os"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/gorilla/mux"

	"profitability/api/pkg/config"
	"profitability/api/pkg/cache"
	"profitability/api/pkg/integrations"
	"profitability/api/pkg/middleware"
	"profitability/api/pkg/controllers"
	"profitability/api/pkg/services"	
	"profitability/api/pkg/html"
)

func main() {

	config.LoadEnv()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	corsHandler := cors.New(cors.Options{
	    AllowedOrigins:   []string{"*"},
	    AllowedMethods:   []string{"GET"},
	    AllowedHeaders:   []string{"Content-Type"},
	    AllowCredentials: true,
	})

	ratesCache := cache.NewRatesCache()
	
	router := mux.NewRouter()

	router.Use(corsHandler.Handler)
	
	// Services
	ratesService := services.NewRatesService(ratesCache)
	
	propService := services.NewPropService(ratesCache)
	preService := services.NewPreService()
	posService := services.NewPosService(ratesCache)
	ipcaService := services.NewIpcaService(ratesCache)
	
	// Controllers
	ratesController := controllers.NewRatesController(ratesService)	
	
	propController := controllers.NewPropController(propService)
	preController := controllers.NewPreController(preService)
	posController := controllers.NewPosController(posService)	
	ipcaController := controllers.NewIpcaController(ipcaService)

	// Routes
	router.HandleFunc("/api/rates", ratesController.GetRates).Methods("GET")

    router.Handle("/api/prop", middleware.ValidateInput(middleware.ProfitHandler(propController.ProfitProp)))
    router.Handle("/api/pre", middleware.ValidateInput(middleware.ProfitHandler(preController.ProfitPre)))
    router.Handle("/api/pos", middleware.ValidateInput(middleware.ProfitHandler(posController.ProfitPos)))
    router.Handle("/api/ipca", middleware.ValidateInput(middleware.ProfitHandler(ipcaController.ProfitIpca)))
  
	// Ingetregrations
	if _, err := integrations.GetSELICData(ratesCache); err != nil {
		log.Fatalf("Error retrieving exchange rate data: %v", err)
	}

	if _, err := integrations.GetIPCAData(ratesCache); err != nil {
		log.Fatalf("Error retrieving exchange rate data: %v", err)
	}
	
	
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        welcomeHTML := html.WelcomePageHTML()

        w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(welcomeHTML))
    })
    
	log.Printf("Server started on port %s\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
