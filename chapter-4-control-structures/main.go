package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//1.
	randNumbers := make([]int, 0, 100)

	for range 100 {
		randNumbers = append(randNumbers, rand.Intn(101))
	}

	//2.
	for _, v := range randNumbers {
		switch {
		case (v%2 == 0 && v%3 == 0):
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}

	//3.
	var total int
	for i := 0; i < 10; i++ {
		// total := total + i
		total = total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}
