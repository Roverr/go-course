package main

type order struct {
	identification
	items []item
}

type item struct {
	identification
}

type identification struct {
	id string
}

// worker recieves an only readable channel "in" and an only writable channel "out"
// tries to send every item to the item processor
// which is represented by the out channel
func worker(in <-chan order, out chan<- item) {
	for {
		value := <-in
		for _, i := range value.items {
			out <- i
		}
	}
}
