package jexl_test

import (
	"testing"

	"github.com/manfromth3m0oN/jexl-go"
)

func TestJexl(t *testing.T) {
	testExpr := "6 * 12 + 5 / 2.6"
	res, err := jexl.Eval(testExpr)
	if err != nil {
		t.FailNow()
	}
	if res != "73.92307692307692" {
		t.FailNow()
	}
}
