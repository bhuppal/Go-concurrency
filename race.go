package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		data++
	}()
  wg.Wait()
	if data == 0 {
		fmt.Printf("The value is %v\n", data)
	} else {
		fmt.Printf("value...%d", data)
	}

	//wg.Wait()
}
