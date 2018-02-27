package main

import (
	"math/rand"
	"net/http"
	"time"
)

// doOtherThings is to mock out the main thread
func doOtherThings() {
	for {
		<-time.After(10000000)
	}
}

// isPingSuccessful is just a mock, to receive an error sometimes therefore prove that the concept works
func isPingSuccessful(resp *http.Response) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Int()%3 == 0 {
		return false
	}
	return true
}
