package util_test

import (
//	"errors"
	"testing"

	"profitability/cli/pkg/util"
)

func TestAliquot(t *testing.T) {
	tests := []struct {
		term           int
		expectedRate   float64
		expectedErrMsg string
	}{
		{term: 100, expectedRate: 0.225, expectedErrMsg: ""},
		{term: 180, expectedRate: 0.225, expectedErrMsg: ""},
		{term: 181, expectedRate: 0.20, expectedErrMsg: ""},
		{term: 360, expectedRate: 0.20, expectedErrMsg: ""},
		{term: 361, expectedRate: 0.175, expectedErrMsg: ""},
		{term: 720, expectedRate: 0.175, expectedErrMsg: ""},
		{term: 721, expectedRate: 0.15, expectedErrMsg: ""},
		{term: 1000, expectedRate: 0.15, expectedErrMsg: ""},
	}

	for _, test := range tests {
		rate, err := util.Aliquot(test.term)

		if rate != test.expectedRate {
			t.Errorf("For term %d, expected tax rate %.3f but got %.3f", test.term, test.expectedRate, rate)
		}

		if (err != nil && test.expectedErrMsg == "") || (err == nil && test.expectedErrMsg != "") {
			t.Errorf("For term %d, expected error '%s' but got error: %v", test.term, test.expectedErrMsg, err)
		} else if err != nil && err.Error() != test.expectedErrMsg {
			t.Errorf("For term %d, expected error '%s' but got error '%v'", test.term, test.expectedErrMsg, err)
		}
	}
}

