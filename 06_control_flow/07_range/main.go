package main

import "fmt"

func main() {

	// A simple slice
	names := []string{
		"Ali",
		"Ahmed",
		"Sara",
	}

	// range gives:
	// index = position
	// value = actual item
	for index, value := range names {
		fmt.Println(index, value)
	}
}
