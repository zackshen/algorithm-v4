package main

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	values := []string{"4", "3", "2", "1", "a", "b", "c", "d"}
	l := NewLinkedList()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.PushBack("d")
	l.PushFront("1")
	l.PushFront("2")
	l.PushFront("3")
	l.PushFront("4")

	if l.Len() != 8 {
		t.Error("LinkedList size error")
	}

	firstElement := l.Front()
	v, ok := firstElement.Value.(string)
	if ok && v != "4" {
		t.Error("Get LinkedList first element failed.")
	}

	if firstElement.Prev() != nil {
		t.Error("LinkedList first element has previous element.")
	}

	lastElement := l.Back()
	v, ok = lastElement.Value.(string)
	if ok && v != "d" {
		t.Error("Get LinkedList last element failed.")
	}

	if lastElement.Next() != nil {
		t.Error("LinkedList last element has next element.")
	}

	element := firstElement
	cursor := 0
	for {
		v, ok = element.Value.(string)
		fmt.Printf("%s -> %s\n", v, values[cursor])
		if v != values[cursor] {
			t.Error("LinkedList get item fail")
		}
		cursor++
		if element.Next() == nil {
			break
		}
		element = element.Next()
	}

	toRemoveElement := l.Front().Next().Next().Next()
	t.Log("to remove element is ", toRemoveElement.Value)
	removeValue := l.Remove(toRemoveElement)
	if v := removeValue.(string); v != "1" {
		t.Error("LinkedList remove error")
	}

	if l.Len() != 7 {
		t.Error("LinkedList remove error")
	}

	fmt.Println(l.Front().Value)
	fmt.Println(l.Front().Next().Value)
	fmt.Println(l.Front().Next().Next().Value)
	fmt.Println(l.Front().Next().Next().Next().Value)
	fmt.Println(l.Front().Next().Next().Next().Next().Value)

}
