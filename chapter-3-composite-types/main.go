package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	//1.
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	first := greetings[:2]
	second := greetings[1:4]
	third := greetings[3:]

	fmt.Println(first, second, third)

	//2.
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Println(string(runes[3]))

	//3.
	employeeOne := Employee{"first", "last", 1}
	employeeTwo := Employee{
		firstName: "firsty",
		lastName:  "lasty",
		id:        2,
	}
	var employeeThree Employee
	employeeThree.firstName = "firster"
	employeeThree.lastName = "laster"
	employeeThree.id = 3

	fmt.Println(employeeOne, employeeTwo, employeeThree)
}
