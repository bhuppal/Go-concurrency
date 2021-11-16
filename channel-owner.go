package main

import "fmt"

func main() {
	/*
	Create channel owner, which creates channel
	return receive only channel to caller
	spins a goroutine, which writes data into channel and
	closes the channel when done.
	 */

	owner := func() <-chan int { // Owner channel
		ch := make(chan int) // creating a channel

		go func() { // creating a goroutine
			defer close(ch) // closing the channel
			for i:=0; i<5;i++{
				ch <- i // sending a value to the channel
			}
		}()
		return ch
	}

	consumer := func(ch <- chan int) {
		// read values from channel
		for v := range ch {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("Done receiving!!!")
	}

	ch := owner()
	consumer(ch)
}