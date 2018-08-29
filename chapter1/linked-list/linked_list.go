package main

type Element struct {
	prev  *Element
	next  *Element
	Value interface{}
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

type LinkedList struct {
	first *Element
	last  *Element
	size  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		first: nil,
		last:  nil,
	}
}

// TODO PushBack和PushFront还没有将list连接起来
func (l *LinkedList) PushBack(item interface{}) *Element {
	var firstElement *Element
	if l.first == nil {
		l.first = &Element{
			prev:  nil,
			next:  nil,
			Value: item,
		}
		firstElement = l.first
	}

	if l.last == nil {
		l.last = firstElement
	} else {
		l.last = &Element{
			prev:  l.last,
			next:  nil,
			Value: item,
		}
	}
	l.size++
	return l.last
}

// TODO PushBack和PushFront还没有将list连接起来
func (l *LinkedList) PushFront(item interface{}) *Element {
	var lastElement *Element
	if l.last == nil {
		l.last = &Element{
			prev:  nil,
			next:  nil,
			Value: item,
		}
		lastElement = l.last
	}

	if l.first == nil {
		l.first = lastElement
	} else {
		l.first = &Element{
			prev:  nil,
			next:  l.first,
			Value: item,
		}
	}
	l.size++
	return l.first
}

func (l *LinkedList) Front() *Element {
	return l.first
}

func (l *LinkedList) Back() *Element {
	return l.last
}

func (l *LinkedList) Len() int {
	return l.size
}

func main() {

}
