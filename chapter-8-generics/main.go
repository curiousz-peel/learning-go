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
	var intList LinkedList[int]
	intList.Add(10)
	intList.Add(12)
	intList.Add(13)
	intList.Add(14)
	intList.Add(15)
	intList.Insert(1, 0)
	intList.Insert(1, 5)
	fmt.Println(intList.Index(10))
	fmt.Println(intList.Index(16))

	for el := intList.head; el != nil; el = el.next {
		fmt.Println(el.val)
	}

	var stringList LinkedList[string]
	stringList.Add("a")
	stringList.Add("b")
	stringList.Add("c")
	stringList.Add("d")
	stringList.Add("e")
	stringList.Insert("a", 0)

	for el := stringList.head; el != nil; el = el.next {
		fmt.Println(el.val)
	}

	fmt.Println(stringList.Index("c"))
}
