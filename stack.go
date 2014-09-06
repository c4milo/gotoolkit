package gotoolkit

import "errors"

type node struct {
	item interface{}
	next *node
}

// A stack backed by a linked list, it may be faster than SliceStack
// but consumes more memory and has worse cache locality.
type Stack struct {
	first *node
	Size  uint64
}

// Pushes an element to the stack
func (s *Stack) Push(item interface{}) {
	oldFirst := s.first
	s.first = new(node)
	s.first.item = item
	s.first.next = oldFirst
	s.Size++
}

// Removes and return the latest added element from the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Unable to pop element, stack is empty")
	}
	item := s.first.item
	s.first = s.first.next
	s.Size--
	return item, nil
}

// Whether or not the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.first == nil
}
