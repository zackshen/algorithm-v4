package main

type Set struct {
	container map[interface{}]bool
}

func NewSet() *Set {
	return &Set{container: make(map[interface{}]bool)}
}

func (s *Set) Add(item interface{}) bool {
	if hasElement := s.container[item]; hasElement {
		return false
	}
	s.container[item] = true
	return true
}

func (s *Set) Remove(item interface{}) {
	delete(s.container, item)
}

func (s *Set) Size() int {
	return len(s.Items())
}

func (s *Set) Items() []interface{} {
	items := make([]interface{}, 0)
	for k := range s.container {
		items = append(items, k)
	}
	return items
}

func main() {

}
