package util

import (
	"fmt"
	"errors"
	"math"
	"strconv"
)

func IsValidFunctionality(functionality string) bool {
	validFunctionalitys := map[string]bool{
		"pre":  true,
		"pos":  true,
		"ipca": true,
		"prop": true,
	}

	return validFunctionalitys[functionality]
}

func IsValidRate(value string) (float64, error) {
    parsedValue, err := strconv.ParseFloat(value, 64)
    if err != nil {
        return 0, fmt.Errorf("valor inválido para float: %v", err)
    }

	if parsedValue <= 0 || parsedValue > 399.99 {
		return 0, errors.New("taxa deve ser entre 1 e 399.99")
	}
	    
    return parsedValue, nil
}

func IsValidTerm(value string) (int, error) {
	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("valor inválido para int: %v", err)
	}
	
	if parsedValue <= 0 {
		return 0, errors.New("prazo deve ser maior que zero")
	}

    if parsedValue > math.MaxInt32 {
		return 0, errors.New("guardando dinheiro para as futuras gerações?")
    }
    
	return int(parsedValue), nil
}
