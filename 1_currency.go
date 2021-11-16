package main

import (
	"fmt"
	"runtime"
	"sync"
)

func DisplayPrimeNumbers() func(startNum int, endNum int, waitGroup *sync.WaitGroup) {
	return func(startNum int, endNum int, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		for i := startNum; i <= endNum; i++ {
			if i%2 == 0 {
				fmt.Printf("%v is a prime number\n", i)
			}
		}
	}
}

func DisplayAlphabetUpperCase(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	// Display 3 time
	for i:=1; i<=3;i++{
		fmt.Printf("%v times: ", i)
		for char:='A'; char<='Z'; char++ {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}
}

func DisplayAlphabetLowerCase(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	//Display 3 times
	for i:=1; i<=3; i++ {
		fmt.Printf("%v times: ", i)
		for char:='a'; char<='z'; char++ {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}
	waitGroup.Done()
}


func main() {
	// Allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(2)

	// wg is used to wait for the program to finish
	// Add a count of two, one for each goroutine
	// WaitGroup is a counting semaphore that can be used
	// to maintain a record of running goroutines
	var wg sync.WaitGroup
	wg.Add(4)

	//using direct function call
	go DisplayAlphabetLowerCase(&wg)

	//using anonymous function
	go func() {
		DisplayAlphabetUpperCase(&wg)
	}()

	//using function value
	fn := DisplayAlphabetLowerCase
	go fn(&wg)

	//using function closure
	fnClosure := DisplayPrimeNumbers()
	go fnClosure(2, 20, &wg)

	//DisplayAlfaphatUpperCase()
	//DisplayAlfaphatLowerCase()

	wg.Wait()
	fmt.Println("Ending the program")
}
