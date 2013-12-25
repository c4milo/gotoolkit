package gotoolkit

import "errors"

type SliceStack struct {
    Size uint64
    backing []interface{}
}

func (s *SliceStack) Push(item interface{}) {
    s.Size++
    s.backing = append(s.backing, item)
}

func (s *SliceStack) Pop() (interface{}, error) {
    if s.IsEmpty() {
        return nil, errors.New("Stack is empty.")
    }

    s.Size--
    item := s.backing[s.Size]
    s.backing = s.backing[0:s.Size]
    return item, nil
}

func (s *SliceStack) IsEmpty() bool {
    return len(s.backing) == 0
}

