package main

import (
	"fmt"
)

func main() {
	//1.
	var (
		i int = 10
		f float64
	)

	f = float64(i)

	fmt.Println(i)
	fmt.Println(f)

	//2.
	const value = 5

	i = value
	f = value

	fmt.Println(i)
	fmt.Println(f)

	//3.
	var (
		b      byte   = 255
		smallI int32  = 2_147_483_647
		bigI   uint64 = 18_446_744_073_709_551_615
	)

	b, smallI, bigI = b+1, smallI+1, bigI+1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
