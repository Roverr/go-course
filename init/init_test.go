package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	fmt.Println("Run")
	code := t.Run()
	os.Exit(code)
}

func TestFunction(t *testing.T) {
	t.Log("TestFunction")
}
func TestInit(t *testing.T) {
	t.Log("TestInit")
}
