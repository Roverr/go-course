package main

import (
	"time"
)

func main() {
	orders := []order{
		order{
			identification: identification{id: "1"},
			items: []item{
				item{
					identification: identification{id: "1"},
				},
				item{
					identification: identification{id: "2"},
				},
				item{
					identification: identification{id: "3"},
				},
			},
		},
		order{
			identification: identification{id: "2"},
			items: []item{
				item{
					identification: identification{id: "4"},
				},
				item{
					identification: identification{id: "2"},
				},
				item{
					identification: identification{id: "1"},
				},
			},
		},
	}

	in := make(chan order, 1)
	out := make(chan item, 3)
	go worker(in, out)
	go processor(out)
	for _, o := range orders {
		in <- o
	}
	time.Sleep(time.Second * 1)
}
