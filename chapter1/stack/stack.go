package main

import (
	"errors"
)

/*
Stack 栈
*/
type Stack struct {
	items  []interface{}
	cursor int
}

/*
NewStack
*/
func NewStack() *Stack {
	return &Stack{items: make([]interface{}, 0), cursor: 0}
}

/*
Push 添加一个元素
*/
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
	s.cursor++
}

/*
Pop 删除最近添加的元素
*/
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("no elements in stack")
	}
	item := s.items[s.cursor-1]
	s.items[s.cursor-1] = nil
	s.cursor--
	return item, nil
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
	return s.cursor
}

func main() {
}
