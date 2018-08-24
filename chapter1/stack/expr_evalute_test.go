package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExprEvaluate(t *testing.T) {
	val := ExprEvaluate("(1 + 100)")
	if val != 101 {
		t.Errorf("(1+100) evaluate fail: %v\n", val)
	}

	val = ExprEvaluate("((1 + 100) + 100)")
	if val != 201 {
		t.Errorf("((1+100) + 100) evaluate fail: %v\n", val)
	}

	val = ExprEvaluate("((1 - 100) + 100)")
	if val != 1 {
		t.Errorf("((1-100) + 100) evaluate fail: %v\n", val)
	}

	val = ExprEvaluate("((2 * 3) + 4)")
	if val != 10 {
		t.Errorf("((2*3) + 4) evaluate fail: %v\n", val)
	}

	val = ExprEvaluate("((4 / 2) + 5)")
	if val != 7 {
		t.Errorf("((4/2) + 5) evaluate fail: %v\n", val)
	}

	val = ExprEvaluate("((3 / 2) + 5)")
	if val != 6.5 {
		t.Errorf("((3/2) + 5) evaluate fail: %v\n", val)
	}

	val = ExprEvaluate("((1 / 3) + 1)")
	if val != 1.333333 {
		t.Errorf("((1/3) + 1) evaluate fail: %v\n", val)
	}

	assert.Panics(t, func() {
		val = ExprEvaluate("((a / 2) + 5)")
		if val != 6.5 {
			t.Errorf("((3/2) + 5) evaluate fail: %v\n", val)
		}
	})

	assert.Panics(t, func() {
		val = ExprEvaluate("((100 / 0) + 5)")
		if val != 6.5 {
			t.Errorf("((100/0) + 5) evaluate fail: %v\n", val)
		}
	})

}
