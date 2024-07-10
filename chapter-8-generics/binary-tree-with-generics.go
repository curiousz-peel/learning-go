package main

import "cmp"

type OrderableFunc[T any] func(t1, t2 T) int

type GenericTree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func NewGenericTree[T any](f OrderableFunc[T]) *GenericTree[T] {
	return &GenericTree[T]{
		f: f,
	}
}

func (t *GenericTree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *GenericTree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *Node[T]) Add(f OrderableFunc[T], val T) *Node[T] {
	if n == nil {
		return &Node[T]{val: val}
	}
	switch cmp := f(val, n.val); {
	case cmp >= 1:
		return n.right.Add(f, val)
	case cmp <= -1:
		return n.left.Add(f, val)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], val T) bool {
	if n == nil {
		return false
	}
	switch cmp := f(val, n.val); {
	case cmp >= 1:
		return n.right.Contains(f, val)
	case cmp <= -1:
		return n.left.Contains(f, val)
	}
	return true
}

// dummy struct for code example
type Person struct {
	Name string
	Age  int
}

func (p1 Person) Order(p2 Person) int {
	r := cmp.Compare(p1.Name, p2.Name)
	if r == 0 {
		r = cmp.Compare(p1.Age, p2.Age)
	}
	return r
}
