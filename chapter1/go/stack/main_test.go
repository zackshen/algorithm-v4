package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	if stack.Size() != 4 {
		t.Errorf("Stack size was incorrect")
	}

	for _, value := range []int{4, 3, 2, 1} {
		item, _ := stack.Pop()
		if item != value {
			t.Errorf("Stack pop item is not equal %d", value)
		}
	}

	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Stack is empty, should raise error")
	}

	if !stack.IsEmpty() {
		t.Errorf("stack is not empty")
	}

}
