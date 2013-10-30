package gotoolkit

import "errors"

type node struct {
	item interface{}
	next *node
}

type Stack struct {
	first *node
	Size  uint64
}

func (s *Stack) Push(item interface{}) {
	oldFirst := s.first
	s.first = new(node)
	s.first.item = item
	s.first.next = oldFirst
	s.Size++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Unable to pop element, stack is empty")
	}
	item := s.first.item
	s.first = s.first.next
	s.Size--
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return s.first == nil
}
