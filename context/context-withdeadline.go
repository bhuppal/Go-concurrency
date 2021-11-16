package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {
	// set deadline for goroutine to return computational result
	deadline := time.Now().Add(10 * time.Millisecond)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	compute := func() <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			deadline, ok := ctx.Deadline()
			if ok {
				if deadline.Sub(time.Now().Add(50 * time.Millisecond)) < 0 {
					fmt.Println("Not sufficient time given...terminating")
				}
			}
			time.Sleep(50 * time.Millisecond)
			select {
			case ch <- data{"123"}:
				case <-ctx.Done():
					return
			}
//			ch <- data{"123"}
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on
	ch := compute()
	d, ok   := <-ch
	if ok {
		fmt.Printf("Work is complete: %s\n", d)
	}

}
