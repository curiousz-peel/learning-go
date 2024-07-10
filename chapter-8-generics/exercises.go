package main

import (
	"fmt"
	"strconv"
)

// 1.
type Doubler interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | uint
}

func Double[T Doubler](val T) T {
	return 2 * val
}

// 2.
type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintableInt int

func (p PrintableInt) String() string {
	return strconv.Itoa(int(p))
}

type PrintableFloat64 float64

func (p PrintableFloat64) String() string {
	// return strconv.FormatFloat(float64(p), 'f', 3, 64)
	return fmt.Sprintf("%f", p)
}

func DoPrint[T Printable](p T) {
	fmt.Println(p)
}

// 3.
type LinkedList[T comparable] struct {
	head *LinkedListNode[T]
	len  int
}

type LinkedListNode[T comparable] struct {
	val  T
	next *LinkedListNode[T]
}

func (l *LinkedList[T]) Add(v T) {
	if l.head == nil {
		l.head = &LinkedListNode[T]{val: v}
		l.len++
		return
	}
	var el = l.head
	for ; el.next != nil; el = el.next {
	}
	el.next = &LinkedListNode[T]{val: v}
	l.len++
}

func (l *LinkedList[T]) Insert(v T, idx int) bool {
	if idx < 0 || idx > l.len-1 {
		return false
	}

	if idx == 0 {
		oldHead := l.head
		newHead := &LinkedListNode[T]{val: v, next: oldHead}
		l.head = newHead
		l.len++
		return true
	}

	var elPrev = l.head
	for i := 0; i < idx; {
		elPrev = elPrev.next
		i++
	}

	elCurr := &LinkedListNode[T]{val: v, next: elPrev.next}
	elPrev.next = elCurr
	l.len++
	return true
}

func (l *LinkedList[T]) Index(v T) (int, bool) {
	var idx int

	for el := l.head; el.next != nil; el = el.next {
		if el.val == v {
			return idx, true
		}
		idx++
	}
	return 0, false
}
