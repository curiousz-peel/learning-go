package chapter10modules

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add [adds] two numbers of type [constraints.Integer] or [constraints.Float]
//
// [adds] : https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](x, y T) T {
	return x + y
}
