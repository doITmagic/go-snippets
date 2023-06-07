/*
The Fan out/fan in pattern in Go allows for the processing of activities in parallel,
where multiple functions are spun to execute the activities, and at the end of processing,
the orchestrator waits for all parallel activities to finish
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func cube(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n * n
		}
	}()
	return out
}

func merge(in ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(in))
	for _, c := range in {
		go output(c)
	}

	go func() {
		wg.Wait()
	}()

	return out

}

func main() {
	fmt.Println("Start Fan Out ")
	// Producer : Fan Out
	c1 := generator(1, 2, 3)
	c2 := generator(4, 5, 6)

	// Transformer
	cube1 := cube(c1)
	cube2 := cube(c2)

	// Fan In
	out := merge(cube1, cube2)
	go func() {
		for n := range out {
			fmt.Println(n)
		}
	}()

	time.Sleep(time.Second * 2)

}
