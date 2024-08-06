package main

import (
	"fmt"
	"math"
	"sync"
)

func exerciseOne() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	valsOne := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	valsTwo := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	ch := make(chan int)

	go func() {
		for _, v := range valsOne {
			ch <- v
		}
		wg.Done()
	}()

	go func() {
		for _, v := range valsTwo {
			ch <- v
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	wg2 := sync.WaitGroup{}
	wg2.Add(1)

	go func() {
		for v := range ch {
			println(v)
		}
		wg2.Done()
	}()

	wg2.Wait()
}

func exerciseTwo() {
	chOne, chTwo := make(chan int), make(chan int)

	valsOne := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	valsTwo := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	go func() {
		for _, v := range valsOne {
			chOne <- v
		}
		close(chOne)
	}()

	go func() {
		for _, v := range valsTwo {
			chTwo <- v
		}
		close(chTwo)
	}()

	activeChannels := 2
	for activeChannels != 0 {
		select {
		case val, ok := <-chOne:
			if !ok {
				activeChannels--
				chOne = nil
				continue
			}
			fmt.Println("from chOne:", val)
		case val, ok := <-chTwo:
			if !ok {
				activeChannels--
				chTwo = nil
				continue
			}
			fmt.Println("from chTwo:", val)
		}
	}
}

func exerciseThree() {
	var sqMap = sync.OnceValue(sqRootMap)

	for i := 0; i < 100_000; i += 1_000 {
		fmt.Println(i, sqMap()[i])
	}
}

func sqRootMap() map[int]float64 {
	fmt.Println("Calculating square roots...")
	var sqMap = make(map[int]float64, 100_000)
	for i := 0; i < 100_000; i++ {
		sqMap[i] = math.Sqrt(float64(i))
	}
	return sqMap
}

func main() {
	// exerciseOne()
	// exerciseTwo()
	exerciseThree()
}
