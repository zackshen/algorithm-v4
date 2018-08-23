package main

import (
	"testing"
)

var stackWithElements *Stack
var StackElementCount = 100000

func init() {
	stackWithElements = NewStack()
	for i := 0; i < StackElementCount; i++ {
		stackWithElements.Push("A")
	}
}

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

func BenchmarkStackPush(b *testing.B) {
	stack := NewStack()
	for i := 0; i < b.N; i++ {
		stack.Push("A")
	}
}

func BenchmarkStackPop(b *testing.B) {
	for i := 0; i < StackElementCount; i++ {
		stackWithElements.Pop()
	}
}

func BenchmarkStackSize(b *testing.B) {
	stackWithElements.Size()
}
