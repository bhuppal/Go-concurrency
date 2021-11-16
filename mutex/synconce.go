package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
    var once sync.Once

	load := func() {
		fmt.Println("Run only once initialization function")
	}

	for i:=0; i<10;i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			//should call only once
			once.Do(load)
			//load()
		}()
	}

	wg.Wait()
}
