package main

import (
	"fmt"
	"sync"
)

// generator - converts a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square - receive on inbound channel
// square the number
// output on outbount channel
func square(in <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <- chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
    in := generator(2, 3)

	// fan out square stage to run two instances
	ch1 := square(in)
	ch2 := square(in)

	// fan in - results of sqaure stages
	//for n := range merge(ch1, ch2) {
//		fmt.Println(n)
//	}
    out := merge(ch1, ch2)
	fmt.Println(<-out)

}
