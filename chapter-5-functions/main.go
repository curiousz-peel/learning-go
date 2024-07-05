package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func add(i, j int) (int, error) { return i + j, nil }
func sub(i, j int) (int, error) { return i - j, nil }
func mul(i, j int) (int, error) { return i * j, nil }
func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("error: division by 0")
	}
	return i / j, nil
}

var calcMap = map[string]func(i, j int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func fileLen(name string) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	buffer := make([]byte, 512)
	var bytesNo int
	for {
		count, err := file.Read(buffer)
		bytesNo += count
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}
	return bytesNo, nil
}

func prefixer(prefix string) func(string) string {
	return func(base string) string {
		return fmt.Sprintf("%s %s", prefix, base)
	}
}

func main() {
	//1.
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"2", "/", "0"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := calcMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}

	//2.
	bytes, err := fileLen(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File: ", os.Args[1], "\nBytes: ", bytes)

	//3.
	helloPrefixer := prefixer("hello")
	fmt.Println(helloPrefixer("world"))
	fmt.Println(helloPrefixer("swirl"))
}
