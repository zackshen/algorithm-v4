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
}
