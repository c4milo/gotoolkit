// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package gotoolkit

import "errors"

// A stack backed by a growing slice. Useful for memory constrained environments.
// Has better cache locality.
type SliceStack struct {
	Size    uint64
	backing []interface{}
}

// Pushes an element to the stack
func (s *SliceStack) Push(item interface{}) {
	s.Size++
	s.backing = append(s.backing, item)
}

// Removes and return the latest added element from the stack.
func (s *SliceStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty.")
	}

	s.Size--
	item := s.backing[s.Size]
	s.backing = s.backing[0:s.Size]
	return item, nil
}

// Whether or not the stack is empty
func (s *SliceStack) IsEmpty() bool {
	return len(s.backing) == 0
}
