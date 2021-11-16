package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedRsc) == 0 {
//			time.Sleep(1 * time.Millisecond)
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	c.L.Lock()
	sharedRsc["rsc1"] = "bhuppal"
	c.Signal()
	c.L.Unlock()

	wg.Wait()

}
