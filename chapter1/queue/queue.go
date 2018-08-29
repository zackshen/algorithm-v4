package main

import (
	"container/list"
	"errors"
)

/*
Queue (LIFO)
*/
type Queue struct {
	items *list.List
}

/*
NewQueue
*/
func NewQueue() *Queue {
	return &Queue{items: list.New()}
}

/*
enqueue 添加一个元素
*/
func (s *Queue) Enqueue(item interface{}) {
	s.items.PushBack(item)
}

/*
dequeue 删除最早添加的元素
*/
func (s *Queue) Dequeue() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("no elements in queue")
	}
	item := s.items.Front()
	value := item.Value
	s.items.Remove(item)
	return value, nil
}

/*
isEmpty 队列是否为空
*/
func (s *Queue) IsEmpty() bool {
	return s.Size() == 0
}

/*
Size 队列中的元素数量
*/
func (s *Queue) Size() int {
	return s.items.Len()
}

func main() {
}
