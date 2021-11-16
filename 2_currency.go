// This sample program demonstrates how the goroutine scheduler
// will time slice gorountines on a single thread

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg is used to wait for the program to finish
var wg sync.WaitGroup

//main is the entry point for all Go programs
func main() {

	// Allocate 1 logical processors for the scheduler to use
	//runtime.GOMAXPROCS(1)

	// Allocate a logical processor for every available core
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Total no. of CPUs: ", runtime.NumCPU())
	// Add a count of tow, one for each gorouutine
	//wg.Add(2)

	// Create two goroutines
	fmt.Println("Create Goroutines")

	//go printPrime("A")
	//go printPrime("B")

	// Wait for the goroutines to finish
	fmt.Println("Waiting To Finish")
	//wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first 5000 numbers
func printPrime(prefix string) {
	// Schedule the call to done to tell main we are done
	defer wg.Done()

	next:
		for outer := 2; outer < 5000; outer++ {
			for inner := 2; inner < outer; inner++ {
				if outer%inner == 0 {
					continue next
	//				fmt.Println(inner)
				}
			}
			fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}