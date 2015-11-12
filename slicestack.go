// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gotoolkit

import "errors"

// SliceStack implements a stack backed by a growing slice. Useful for memory
// constrained environments. It also has better cache locality.
type SliceStack struct {
	size    uint64
	backing []interface{}
}

// Push adds an element to the top of the stack.
func (s *SliceStack) Push(item interface{}) {
	s.size++
	s.backing = append(s.backing, item)
}

// Peek returns the latest added element without removing it from the stack.
func (s *SliceStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("unable to peek element, stack is empty")
	}
	return s.backing[s.size-1], nil
}

// Pop removes and return the latest added element from the stack.
func (s *SliceStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("unable to pop element, stack is empty")
	}

	s.size--
	item := s.backing[s.size]
	s.backing = s.backing[0:s.size]
	return item, nil
}

// IsEmpty returns whether or not the stack is empty.
func (s *SliceStack) IsEmpty() bool {
	return len(s.backing) == 0
}

// Size returns the number of elements in the stack.
func (s *SliceStack) Size() uint64 {
	return s.size
}
