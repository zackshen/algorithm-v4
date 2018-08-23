package main

import (
	"testing"
)

func TestExprEvaluate(t *testing.T) {
	val := ExprEvaluate("(1 + 100)")
	if val != 101 {
		t.Error("(1+100) evaluate fail")
	}

	val = ExprEvaluate("((1 + 100) + 100)")
	if val != 201 {
		t.Error("((1+100) + 100) evaluate fail")
	}

	val = ExprEvaluate("((1 - 100) + 100)")
	if val != 1 {
		t.Error("((1-100) + 100) evaluate fail")
	}

	val = ExprEvaluate("((2 * 3) + 4)")
	if val != 10 {
		t.Error("((2*3) + 4) evaluate fail")
	}

	val = ExprEvaluate("((4 / 2) + 5)")
	if val != 7 {
		t.Error("((4/2) + 5) evaluate fail")
	}
}
