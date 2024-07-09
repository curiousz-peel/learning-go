package main

import "fmt"

func main() {
	// //generic stack
	// var stringStack Stack[string]
	// stringStack.Push("hello")
	// stringStack.Push("world")

	// fmt.Println(stringStack)

	// stringStack.Pop()

	// fmt.Println(stringStack)

	// var intStack Stack[int]
	// intStack.Push(6)
	// intStack.Push(5)
	// intStack.Push(4)
	// intStack.Push(3)

	// fmt.Println(intStack)
	// fmt.Println(intStack.Contains(3))
	// fmt.Println(intStack.Contains(14))

	//binary tree w interfaces
	var bt *Tree
	bt = bt.Insert(OrderableInt(8))
	bt = bt.Insert(OrderableInt(6))
	bt = bt.Insert(OrderableInt(4))
	bt = bt.Insert(OrderableInt(7))
	bt = bt.Insert(OrderableInt(10))
	bt = bt.Insert(OrderableInt(11))
	bt = bt.Insert(OrderableInt(13))
	fmt.Println(bt)
}
