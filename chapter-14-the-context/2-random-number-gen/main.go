package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func genNumber(ch chan<- int) {
	ch <- rand.Intn(100_000_000)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	defer cancel()

	ch := make(chan (int))
	var sum int
	var iter int

	defer func() {
		fmt.Println(iter, sum)
	}()

	for {
		iter++
		go genNumber(ch)
		select {
		case val := <-ch:
			sum += val
			if sum == 1234 {
				fmt.Println("finished:", "number 1234 reached")
				return
			}
		case <-ctx.Done():
			fmt.Println("timeout:", ctx.Err())
			return
		}
	}
}
