package main

import "fmt"

// Package means a group of Go code.
// Every Go file must start with package name.

// package main means this file can run as a program.
// func main() is the starting point of the program.

var appName = "Go Packages Example"

// This function starts with capital letter.
// In Go, capital letter means public/exported.
// Other packages can use it.
func Add(a int, b int) int {
	return a + b
}

// This function starts with small letter.
// In Go, small letter means private/unexported.
// Only this package can use it.
func subtract(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println("Learning Go Packages")
	fmt.Println("--------------------")

	// Using package-level variable
	fmt.Println("App Name:", appName)

	// Calling exported function
	sum := Add(10, 5)
	fmt.Println("Add result:", sum)

	// Calling unexported function
	difference := subtract(10, 5)
	fmt.Println("Subtract result:", difference)

	// fmt is also a package from Go standard library.
	fmt.Println("fmt package is used for printing output.")
}
