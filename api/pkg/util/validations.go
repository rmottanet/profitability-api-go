package util

import (
    "regexp"
    "strconv"
)


func ValidateRate(rate string) bool {
    // valida formato
    regex := regexp.MustCompile(`^\d{1,3}(\.\d{1,2})?$`)
    if !regex.MatchString(rate) {
        return false
    }

    // valida valor
    rateValue, err := strconv.ParseFloat(rate, 64)
    if err != nil || rateValue <= 0 {
        return false
    }

    return true
}


func ValidateTerm(term string) bool {
    // valida formato
    regex := regexp.MustCompile(`^\d{1,4}$`)
    if !regex.MatchString(term) {
        return false
    }

    // valida valor
    termValue, err := strconv.Atoi(term)
    if err != nil || termValue <= 0 || termValue > 9999 {
        return false
    }

    return true
}
