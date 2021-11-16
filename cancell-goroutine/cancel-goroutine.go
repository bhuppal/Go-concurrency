package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// generator - converts a list of integers to a channel
func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// square - receive on inbound channel
// square the number
// output on outbount channel
func square(done <-chan struct{}, in <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <- chan int) {
		defer wg.Done()
		for n := range c {
			select {
				case out <- n:
				case <-done:
					return
			}
		}
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
	done := make(chan struct{})
	in := generator(done,2, 3)

	// fan out square stage to run two instances
	ch1 := square(done, in)
	ch2 := square(done, in)

	// fan in - results of sqaure stages
	//for n := range merge(ch1, ch2) {
	//		fmt.Println(n)
	//	}
	out := merge(done, ch1, ch2)
	fmt.Println(<-out)
	close(done)
     time.Sleep(10 * time.Millisecond)
	g := runtime.NumGoroutine()
	fmt.Println("No of goroutine are alive ", g)
}
