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

func TestEngine(t *testing.T) {
	testContext := map[string]interface{}{"a": map[string]interface{}{"b": 2.0}}
	engine, err := jexl.NewEngine(testContext, "a.b")
	defer engine.Free()
	if err != nil {
		t.Logf("failed building engine: %v", err)
		t.FailNow()
	}
	res := engine.Run()
	if res != "2" {
		t.Logf("Result is %v, not 2", res)
		t.FailNow()
	}
}
