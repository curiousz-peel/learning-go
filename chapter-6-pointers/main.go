package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func Updateslice(words []string, lastWord string) {
	words[len(words)-1] = lastWord
	fmt.Println("from update slice: ", words)
}

func GrowSlice(words []string, lastWord string) {
	words = append(words, lastWord)
	fmt.Println("from grow slice: ", words)
}

func main() {
	//1.
	_ = MakePerson("Dave", "Grohl", 48)

	_ = MakePersonPointer("Bob", "Irwin", 52)

	//2.
	words1 := []string{"lorem", "ipsum", "dolor"}
	words2 := []string{"lorem", "ipsum", "dolor"}

	Updateslice(words1, "color")
	fmt.Println("words1: ", words1)

	GrowSlice(words2, "color")
	fmt.Println("words2: ", words2)

	//3.
	// main point of exercise 3 is to also give make the desired capacity if it's known beforehand
	// this reduces the amount of reallocations a growing slice needs to perform
	// thus reducing the garbage collector's work, and leaving more resources for the actual program to run
}
