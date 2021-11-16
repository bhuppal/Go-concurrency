// using channel

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		ch <- c
	}(5, 10)

	result := <- ch
	fmt.Printf("Computation of result using channel is %v\n", result)
}