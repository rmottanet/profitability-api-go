package calculus_test

import (
	"testing"

	"profitability/cli/pkg/calculus"
)

func TestPos(t *testing.T) {
	rate := 10.0
	term := 721

	result, err := calculus.Pos(rate, term)
	if err != nil {
		t.Errorf("Error calculating Pos: %v", err)
	}

	expectedResult := 1.00

	if result.ResultPos != expectedResult {
		t.Errorf("Expected Result %v, got %v", expectedResult, result.ResultPos)
	}
}
