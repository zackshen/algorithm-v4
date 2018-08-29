package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	if queue.Size() != 4 {
		t.Errorf("Queue size was incorrect")
	}

	for _, value := range []int{1, 2, 3, 4} {
		item, _ := queue.Dequeue()
		if item != value {
			t.Errorf("Queue pop item is not equal %d", value)
		}
	}

	_, err := queue.Dequeue()
	if err == nil {
		t.Errorf("Queue is empty, should raise error")
	}

	if !queue.IsEmpty() {
		t.Errorf("stack is not empty")
	}

}
