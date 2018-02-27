package main

import (
	"fmt"
	"sync"
)

var emails = []string{
	"john.doe@email.com",
	"jane.doe@email.com",
	"jeremy.doe@email.com",
	"jason.doe@email.com",
	"jay.doe@email.com",
	"jackson.doe@email.com",
	"jeremia.doe@email.com",
	"jack.doe@email.com",
	"jill.doe@email.com",
	"jones.doe@email.com",
}

// Creates a new copied array using mutexes
func mutexCopy() {
	var copies []string
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, email := range emails {
		wg.Add(1)
		go func(email string) {
			defer wg.Done()
			mutex.Lock()
			//fmt.Printf("%d locked mutex", i)
			copies = append(copies, email)
			mutex.Unlock()
			//fmt.Printf("%d released mutex", i)
		}(email)
	}
	wg.Wait()
	fmt.Println("\nWith mutex length:", len(copies))
	fmt.Println(copies)
}

// Creates a simple copy without using mutex
func simpleCopy() {
	var copies []string
	var wg sync.WaitGroup
	for _, email := range emails {
		wg.Add(1)
		go func(email string) {
			defer wg.Done()
			copies = append(copies, email)
		}(email)
	}
	wg.Wait()
	fmt.Println("\nWithout mutex length:", len(copies))
	fmt.Println(copies)
	fmt.Println()
}

func main() {
	mutexCopy()
	simpleCopy()
}
