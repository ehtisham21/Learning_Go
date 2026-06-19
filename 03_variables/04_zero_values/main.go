package main

import "fmt"

func main() {

	// =====================================================
	// Zero Value for string
	// =====================================================

	var firstName string

	fmt.Println("String Zero Value:")
	fmt.Printf("firstName = %q\n", firstName)

	// =====================================================
	// Zero Value for int
	// =====================================================

	var age int

	fmt.Println()
	fmt.Println("Integer Zero Value:")
	fmt.Println("age =", age)

	// =====================================================
	// Zero Value for float64
	// =====================================================

	var marks float64

	fmt.Println()
	fmt.Println("Float Zero Value:")
	fmt.Println("marks =", marks)

	// =====================================================
	// Zero Value for bool
	// =====================================================

	var passed bool

	fmt.Println()
	fmt.Println("Boolean Zero Value:")
	fmt.Println("passed =", passed)

	// =====================================================
	// Display Variable Types
	// =====================================================

	fmt.Println()
	fmt.Println("=== Variable Types ===")

	fmt.Printf("firstName: %T\n", firstName)
	fmt.Printf("age: %T\n", age)
	fmt.Printf("marks: %T\n", marks)
	fmt.Printf("passed: %T\n", passed)

	// =====================================================
	// Checking Zero Values
	// =====================================================

	fmt.Println()
	fmt.Println("=== Checking Zero Values ===")

	if firstName == "" {
		fmt.Println("firstName is empty")
	}

	if age == 0 {
		fmt.Println("age is zero")
	}

	if marks == 0 {
		fmt.Println("marks is zero")
	}

	if passed == false {
		fmt.Println("passed is false")
	}

	// =====================================================
	// Updating Zero Values
	// =====================================================

	fmt.Println()
	fmt.Println("=== Updating Values ===")

	firstName = "Ali"
	age = 20
	marks = 85.5
	passed = true

	fmt.Println("Name:", firstName)
	fmt.Println("Age:", age)
	fmt.Println("Marks:", marks)
	fmt.Println("Passed:", passed)

	// =====================================================
	// Counter Example
	// =====================================================

	fmt.Println()
	fmt.Println("=== Counter Example ===")

	var counter int

	fmt.Println("Initial Counter:", counter)

	counter++

	fmt.Println("After Increment:", counter)

	// =====================================================
	// Final Summary
	// =====================================================

	fmt.Println()
	fmt.Println("=== Final Summary ===")

	var city string
	var population int
	var isCapital bool

	fmt.Printf("City: %q\n", city)
	fmt.Println("Population:", population)
	fmt.Println("Is Capital:", isCapital)
}
