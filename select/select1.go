package main

import (
	"fmt"
	"time"
)

func main() {
	ch5 := make(chan string, 1)

	go func() {
		//defer close(ch5)
		time.Sleep(2 * time.Second)
		ch5 <- "one"
	}()
/*
	m := <- ch5
	fmt.Println("The value is ",m)
	ch5 <- "Two"
	m = <- ch5
	fmt.Println("The value again is ", m)
*/
	select {
		case x := <- ch5:
			fmt.Println("The value is ", x)
		case <-time.After(4 * time.Second):
			fmt.Println("timeout")
	}
	close(ch5)
}
