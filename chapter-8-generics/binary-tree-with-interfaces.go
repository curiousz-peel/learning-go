package main

import (
	"fmt"
	"strings"
)

type Orderable interface {
	Order(v any) int
}

type Tree struct {
	val         Orderable
	left, right *Tree
}

func (t *Tree) Insert(v Orderable) *Tree {
	if t == nil {
		return &Tree{val: v}
	}

	switch cmp := v.Order(t.val); {
	case cmp < 0:
		t.left = t.left.Insert(v)
	case cmp > 0:
		t.right = t.right.Insert(v)
	}
	return t
}

type OrderableInt int

func (o OrderableInt) Order(v any) int {
	return int(o - v.(OrderableInt))
}

type OrderableString string

func (o OrderableString) Order(v any) int {
	return strings.Compare(string(o), string(v.(OrderableString)))
}

func (t *Tree) String() string {
	var helper func(tree *Tree, indent int) string
	helper = func(tree *Tree, indent int) string {
		if tree == nil {
			return ""
		}

		var res = ""
		if tree.right != nil {
			res += helper(tree.right, indent+1)
		}

		res += fmt.Sprintf("%s%v\n", strings.Repeat("  ", indent), tree.val)

		if tree.left != nil {
			res += helper(tree.left, indent+1)
		}

		return res
	}

	return helper(t, 0)
}
