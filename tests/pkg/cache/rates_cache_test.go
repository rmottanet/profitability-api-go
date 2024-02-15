package cache_test

import (
    "testing"
    "time"

    "profitability/api/pkg/cache"
)


func TestRatesCache_SetAndGet(t *testing.T) {
	cache := cache.NewRatesCache()

	// Definindo uma chave e um valor para serem armazenados no cache
	key := "SELIC"
	value := 4.5
	expiration := 24 * time.Hour

	// Inserindo o valor no cache
	cache.Set(key, value, expiration)

	// Verificando se o valor inserido pode ser recuperado do cache
	retrievedValue, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected key %s to exist in cache", key)
	}
	if retrievedValue != value {
		t.Errorf("Expected value %f for key %s, got %f", value, key, retrievedValue)
	}
}

func TestRatesCache_GetAll(t *testing.T) {
	cache := cache.NewRatesCache()

	// Definindo alguns valores para serem armazenados no cache
	cache.Set("SELIC", 4.5, 24*time.Hour)
	cache.Set("IPCA", 3.5, 24*time.Hour)

	// Recuperando todos os valores do cache
	rates, _, err := cache.GetAll()
	if err != nil {
		t.Errorf("Error retrieving all cache values: %v", err)
	}

	// Verificando se os valores recuperados estão corretos
	expectedRates := map[string]float64{"SELIC": 4.5, "IPCA": 3.5}
	for key, value := range expectedRates {
		if retrievedValue, ok := rates[key]; !ok || retrievedValue != value {
			t.Errorf("Incorrect cache value for key %s: expected %f, got %f", key, value, retrievedValue)
		}
	}
}

func TestRatesCache_GetTimestamp(t *testing.T) {
	cache := cache.NewRatesCache()

	// Verificando o timestamp antes de definir um valor no cache
	initialTimestamp := cache.GetTimestamp()

	// Definindo um valor no cache
	cache.Set("SELIC", 4.5, 24*time.Hour)

	// Verificando o timestamp após definir um valor no cache
	newTimestamp := cache.GetTimestamp()

	// Verificando se o novo timestamp é posterior ao timestamp inicial
	if newTimestamp.Before(initialTimestamp) {
		t.Error("The new timestamp must be after the initial timestamp")
	}
}
