package main

import (
	"testing"
)

func TestExprEvaluate(t *testing.T) {
	val, err := ExprEvaluate("(1 + 100)")
	if err != nil {
		t.Error(err)
	}
	if val != 101 {
		t.Errorf("(1+100) evaluate fail: %v\n", val)
	}

	val, err = ExprEvaluate("((1 + 100) + 100)")
	if err != nil {
		t.Error(err)
	}
	if val != 201 {
		t.Errorf("((1+100) + 100) evaluate fail: %v\n", val)
	}

	val, err = ExprEvaluate("((1 - 100) + 100)")
	if err != nil {
		t.Error(err)
	}
	if val != 1 {
		t.Errorf("((1-100) + 100) evaluate fail: %v\n", val)
	}

	val, err = ExprEvaluate("((2 * 3) + 4)")
	if err != nil {
		t.Error(err)
	}
	if val != 10 {
		t.Errorf("((2*3) + 4) evaluate fail: %v\n", val)
	}

	val, err = ExprEvaluate("((4 / 2) + 5)")
	if err != nil {
		t.Error(err)
	}
	if val != 7 {
		t.Errorf("((4/2) + 5) evaluate fail: %v\n", val)
	}

	val, err = ExprEvaluate("((3 / 2) + 5)")
	if err != nil {
		t.Error(err)
	}
	if val != 6.5 {
		t.Errorf("((3/2) + 5) evaluate fail: %v\n", val)
	}

	val, err = ExprEvaluate("((1 / 3) + 1)")
	if err != nil {
		t.Error(err)
	}
	if val != 1.333333 {
		t.Errorf("((1/3) + 1) evaluate fail: %v\n", val)
	}

	val, err = ExprEvaluate("((a / 2) + 5)")
	if err == nil {
		t.Error(err)
	}

	val, err = ExprEvaluate("((100 / 0) + 5)")
	if err == nil {
		t.Error(err)
	}

}
