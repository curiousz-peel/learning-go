package main

import "fmt"

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.vals) == 0 {
		return zero, false
	}
	first := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return first, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func (s Stack[T]) String() string {
	var print string
	for i := len(s.vals) - 1; i >= 0; i-- {
		print += fmt.Sprintf("%v\n", s.vals[i])
	}
	return print
}
