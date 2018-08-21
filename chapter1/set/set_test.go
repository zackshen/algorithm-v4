package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	set := NewSet()
	set.Add("1")
	set.Add("hello")
	set.Add("world")

	if hasAdded := set.Add("1"); hasAdded {
		t.Error("add same element '1'")
	}

	t.Log(set.Size())
	if set.Size() != 3 {
		t.Error("Set has wrong size")
	}

	set.Remove("1")
	set.Remove("hello")
	set.Remove("gg")
	t.Log(set.Size())
	if set.Size() != 1 {
		t.Error("Set has wrong size")
	}

	expectedItems := []string{"world"}

	assert.Equal(t, set.Items(), expectedItems, "should be equal")
}
