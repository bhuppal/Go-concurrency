/*
select statement like a switch statement

Each cases specifies communication

All channel operation are considered simultaneously

select waits until some case is ready to proceed

When one the channels is ready, that operation will performed

Non-blocking communication using Select if we use default
Select will wait for channel to communicate, if available, it will execute
otherwise executes the default (select won;t wait for other cases to execute if we mention default)

 */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "two"
	}()
	for i:=1;i<=2;i++ {
		// Multiplex receive on channel - ch1, ch2
		select {
			case m1 := <-ch1:
				fmt.Println(m1)
			case m2 := <-ch2:
				fmt.Println(m2)
		}
	}
}