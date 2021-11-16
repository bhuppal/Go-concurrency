package main

import (
	"fmt"
	"sync"
)

var sharedRsc1 = make(map[string]interface{})

func main() {

	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)


	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc1) < 1 {
			//time.Sleep(1 * time.Millisecond)
			c.Wait()
		}
		c.L.Unlock()
		fmt.Println(sharedRsc1["rsc1"])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc1) < 2 {
			c.Wait()
			//time.Sleep(1 * time.Millisecond)
		}
		c.L.Unlock()
		fmt.Println(sharedRsc1["rsc2"])
	}()
	c.L.Lock()
	sharedRsc1["rsc1"] = "Bhuppal"
	sharedRsc1["rsc2"] = "Kumar"
	c.Broadcast()
	c.L.Unlock()


	wg.Wait()
}
