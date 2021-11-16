// range over channels
/*
Iterate over values received from a channel
Loop automatically breaks, when a channel is closed
range does not return the second boolean value
 */
package main

import "fmt"

// Unbuffered channel
/*
There is no buffer between sender and receiver channels
since there is no buffer, the sender channel will block until there is receiver channel
The receiver channel will block until it have sender channel to send the data
 */

// Buffered channels
/*
There is buffered is specified - capacity.
We can send data no of elements without receiver channel, the sender keep sending the value
without any blockings until the data is buffered (reached the capacity), the receiver will block
in-memory FIFO queue
Asynchronous
ch := make(chan type, capacity
 */

/*
// unbuffered channel example
func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i:=0; i<6; i++{
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n",v)
	}
}
*/


func main() {
	ch := make(chan int, 6)

	go func() {
		defer close(ch)
		for i:=0; i<6;i++{
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}

}
