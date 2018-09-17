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

func (l *LinkedList) PushBack(item interface{}) *Element {
	if l.first == nil && l.last == nil {
		l.first = &Element{
			prev:  nil,
			next:  nil,
			Value: item,
		}
		l.last = l.first
	} else {
		lLast := l.last
		l.last = &Element{
			prev:  lLast,
			next:  nil,
			Value: item,
		}
		lLast.next = l.last
		if l.first.next == nil {
			l.first.next = l.last
		}
	}
	l.size++
	return l.last
}

func (l *LinkedList) PushFront(item interface{}) *Element {
	if l.first == nil && l.last == nil {
		l.first = &Element{
			prev:  nil,
			next:  nil,
			Value: item,
		}
		l.last = l.first
	} else {
		lFirst := l.first
		l.first = &Element{
			prev:  nil,
			next:  lFirst,
			Value: item,
		}
		lFirst.prev = l.first
		if l.last.prev == nil {
			l.last.prev = l.first
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

func (l *LinkedList) Remove(e *Element) interface{} {
	if l.first == nil {
		return nil
	}

	element := l.first.Next()
	for {
		if element == nil {
			return nil
		}
		if element.Value == e.Value {
			nextElement := e.Next()
			element.Prev().next = nextElement
			nextElement.prev = element.Prev()
			l.size--
			e.prev = nil
			e.next = nil
			return e.Value
		}
		element = element.Next()
	}
}

func main() {

}
