package main

import "fmt"

// processor takes an item from the channel and prints it out
func processor(source <-chan item) {
	for {
		item := <-source
		fmt.Println("Lookup and processs item number ", item.id)
	}
}
