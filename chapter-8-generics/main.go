package main

import (
	"fmt"
)

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

	// //binary tree w interfaces
	// var btInt *Tree
	// btInt = btInt.Insert(OrderableInt(8))
	// btInt = btInt.Insert(OrderableInt(6))
	// btInt = btInt.Insert(OrderableInt(4))
	// btInt = btInt.Insert(OrderableInt(7))
	// btInt = btInt.Insert(OrderableInt(10))
	// btInt = btInt.Insert(OrderableInt(11))
	// btInt = btInt.Insert(OrderableInt(13))
	// fmt.Println(btInt)

	// var btString *Tree
	// btString = btString.Insert(OrderableString("baoleu"))
	// btString = btString.Insert(OrderableString("aoleu"))
	// btString = btString.Insert(OrderableString("caoleu"))
	// btString = btString.Insert(OrderableString("daoleu"))
	// btString = btString.Insert(OrderableString("eaoleu"))
	// btString = btString.Insert(OrderableString("aaoleu"))
	// fmt.Println(btString)

	// //binary tree w generics
	// t1 := NewGenericTree(cmp.Compare[int])
	// t1.Add(10)
	// t1.Add(20)
	// t1.Add(15)
	// fmt.Println(t1.Contains(15))
	// fmt.Println(t1.Contains(40))

	// t2 := NewGenericTree(Person.Order)
	// t2.Add(Person{"Ted", 45})
	// t2.Add(Person{"Angel", 26})
	// t2.Add(Person{"Bob", 72})
	// t2.Add(Person{"Tony", 9})
	// fmt.Println(t2.Contains(Person{"Bob", 30}))
	// fmt.Println(t2.Contains(Person{"Bob", 72}))

	//exercises
	//1.
	var myInt int = 32
	var myFloat32 float32 = 32.2
	var myFloat64 float64 = 32.22

	fmt.Println(Double(myInt))
	fmt.Println(Double(myFloat32))
	fmt.Println(Double(myFloat64))

	//2.
	var printableInt PrintableInt = 3
	printableFloat64 := PrintableFloat64(32.222)

	DoPrint(printableInt)
	DoPrint(printableFloat64)

	//3.
	var list LinkedList[int]
	list.Add(10)
	list.Add(12)
	list.Add(13)
	list.Add(14)
	list.Add(15)
	list.Insert(1, 0)
	list.Insert(1, 5)
	fmt.Println(list.Index(10))
	fmt.Println(list.Index(16))

	for el := list.head; el != nil; el = el.next {
		fmt.Println(el.val)
	}
}
