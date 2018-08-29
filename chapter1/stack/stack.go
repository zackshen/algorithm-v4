package main

import (
	"container/list"
	"errors"
)

/*
Stack 栈
*/
type Stack struct {
	items *list.List
}

/*
NewStack
*/
func NewStack() *Stack {
	return &Stack{items: list.New()}
}

/*
Push 添加一个元素
*/
func (s *Stack) Push(item interface{}) {
	s.items.PushBack(item)
}

/*
Pop 删除最近添加的元素
*/
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("no elements in stack")
	}
	item := s.items.Back()
	value := item.Value
	s.items.Remove(item)
	return value, nil
}

/*
isEmpty 栈是否为空
*/
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

/*
Size 栈中的元素数量
*/
func (s *Stack) Size() int {
	return s.items.Len()
}

func main() {
}
