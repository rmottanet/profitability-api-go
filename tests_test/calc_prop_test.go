package calculus_test

import (
	"testing"

	"profitability/cli/pkg/calculus"
	"profitability/cli/pkg/fetcher"
)

func TestProp(t *testing.T) {
	rate := 10.0
	selic, err := fetcher.FetchSelic()
	if err != nil {
		t.Errorf("error fetching selic: %v", err)
	}

	result, err := calculus.Prop(rate)
	if err != nil {
		t.Errorf("error calculating prop: %v", err)
	}

	expected := rate * selic / 100
	if result.ResultProp != expected {
		t.Errorf("expected %v, got %v", expected, result.ResultProp)
	}
}
