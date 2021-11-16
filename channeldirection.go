/*
Channel direction
using channels as function parameters, you can specify if a channel
is meant to only send or receive values
Increases type-safety of the program

func pong(in <-chan string, out chan<- string){}
 */
package main


// Implement relaying of message with Channel Direction
func genMsg(ch1 chan<- string) {//send example
	// Send message on ch1
	ch1 <- "message"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {  //ch1 receive only and ch2 is only to send
	//recv message on ch1
	m := <-ch1
	//send it on ch2
	ch2 <- m
}

func main() {
	// Create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)
	// spine goroutine genMsg and relayMsg
	go genMsg(ch1)
	// recv message on ch2
	go relayMsg(ch1, ch2)
}