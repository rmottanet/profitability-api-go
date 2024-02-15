package services_test

//import (
	//"errors"
	//"testing"

	//"profitability/api/pkg/cache"
	//"profitability/api/pkg/integrations"
	//"profitability/api/pkg/services"
//)

//func TestIpcaService_Ipca(t *testing.T) {
	//// Criando um cache mock para o teste
	//mockCache := &cache.MockCache{
		//Data: map[string]float64{"IPCA": 3.5},
	//}

	//// Criando uma instância do serviço IpcaService com o cache mock
	//ipcaService := services.NewIpcaService(mockCache)

	//// Definindo um valor para a taxa do contrato pós-fixado
	//rate := 5.0

	//// Definindo um termo para o contrato
	//term := 12

	//// Definindo o comportamento esperado da função integrations.GetIPCAData()
	//// Aqui usamos um mock da função GetIPCAData, que sempre retorna nil (sem erro)
	//integrations.GetIPCAData = func(cache.Cache) (float64, error) {
		//return 3.5, nil
	//}

	//// Testando a função Ipca com valores esperados
	//expectedResult := (rate + 3.5/100) * (1 - 0.225)
	//expectedTaxRatePercentage := 22.5
	//resultIpca, taxRatePercentage, err := ipcaService.Ipca(rate, term)

	//// Verificando se o resultado retornado pela função está de acordo com o esperado
	//if err != nil {
		//t.Errorf("Unexpected error: %v", err)
	//}
	//if resultIpca != expectedResult {
		//t.Errorf("Expected IPCA result: %f, got: %f", expectedResult, resultIpca)
	//}
	//if taxRatePercentage != expectedTaxRatePercentage {
		//t.Errorf("Expected tax rate percentage: %f, got: %f", expectedTaxRatePercentage, taxRatePercentage)
	//}

	//// Testando o caso em que a função integrations.GetIPCAData() falha
	//integrations.GetIPCAData = func(cache.Cache) (float64, error) {
		//return 0, errors.New("Error getting IPCA data")
	//}
	//_, _, err = ipcaService.Ipca(rate, term)
	//if err == nil {
		//t.Error("Expected error, got nil")
	//}
//}

