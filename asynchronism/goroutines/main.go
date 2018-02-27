package main

import (
	"fmt"
	"sync"
)

// printNumber prints out a number
func printNumber(i int) {
	fmt.Println(i)
}

func main() {
	// Loop variable captured by function literal is not healthy
	arr := make([]int, 6)
	var wg sync.WaitGroup
	fmt.Println("Starting first loop")
	for i := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			printNumber(i)
		}()
	}
	wg.Wait()

	fmt.Println("Starting second loop")
	for i := range arr {
		wg.Add(1)
		// Naming the parameter the same as the iterator is common
		go func(i int) {
			defer wg.Done()
			printNumber(i)
		}(i)
	}
	wg.Wait()
}
