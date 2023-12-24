package util

import (
	"fmt"
	"strconv"
)

func ParseFloat(s string) (float64, error) {
	valor, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("erro ao converter string para float: %v", err)
	}
	return valor, nil
}

func ParseIPCA(data []map[string]interface{}) (float64, error) {
	var ipca float64

	for _, entry := range data {
		valorStr, ok := entry["valor"].(string)
		if !ok {
			return 0, fmt.Errorf("valor não é uma string")
		}

		valor, err := ParseFloat(valorStr)
		if err != nil {
			return 0, fmt.Errorf("erro ao converter valor IPCA: %v", err)
		}

		ipca += valor
	}

	return ipca, nil
}


func ParseSelic(data []map[string]interface{}) (float64, error) {

	if len(data) == 0 {
		return 0, fmt.Errorf("nenhum dado retornado")
	}
	
	selicStr, ok := data[0]["valor"].(string)
	if !ok {
		return 0, fmt.Errorf("valor não é uma string")
	}

	selic, err := ParseFloat(selicStr)
	if err != nil {
		return 0, fmt.Errorf("erro ao processar dados SELIC: %v", err)
	}
	
	return selic, nil
}
