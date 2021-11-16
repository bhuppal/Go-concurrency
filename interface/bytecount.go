package main

import "fmt"

// ByteCounter type
type ByteCounter int

//To count no of bytes
func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var b ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
