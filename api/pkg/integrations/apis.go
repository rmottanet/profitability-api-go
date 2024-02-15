package integrations


// URLs da API
var urls = map[string]string{
	"ipca":  "https://api.bcb.gov.br/dados/serie/bcdata.sgs.433/dados/ultimos/12?formato=json",
	"selic": "https://api.bcb.gov.br/dados/serie/bcdata.sgs.432/dados/ultimos/1?formato=json",
}

// GetUrl retorna a URL correspondente ao índice fornecido.
func GetUrl(index string) string {
	return urls[index]
}

