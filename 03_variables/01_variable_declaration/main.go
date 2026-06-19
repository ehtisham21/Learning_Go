package main

import "fmt"

// main() is the starting point of the program.
func main() {

	// =====================================================
	// String Variable
	// =====================================================

	// A string variable stores text.
	var firstName string = "Ali"

	fmt.Println("First Name:", firstName)

	// =====================================================
	// Integer Variable
	// =====================================================

	// An int variable stores whole numbers.
	var age int = 20

	fmt.Println("Age:", age)

	// =====================================================
	// Float Variable
	// =====================================================

	// float64 stores decimal numbers.
	var marks float64 = 85.5

	fmt.Println("Marks:", marks)

	// =====================================================
	// Boolean Variable
	// =====================================================

	// bool stores true or false.
	var passed bool = true

	fmt.Println("Passed:", passed)

	// =====================================================
	// Type Inference
	// =====================================================

	// Go can automatically determine the type.
	var city = "Lahore"

	fmt.Println("City:", city)

	// =====================================================
	// Printing Multiple Variables
	// =====================================================

	fmt.Println()
	fmt.Println("=== Student Information ===")

	fmt.Println("Name:", firstName)
	fmt.Println("Age:", age)
	fmt.Println("Marks:", marks)
	fmt.Println("Passed:", passed)
	fmt.Println("City:", city)

	// =====================================================
	// Updating Variable Values
	// =====================================================

	fmt.Println()
	fmt.Println("=== Updating Variables ===")

	fmt.Println("Old Name:", firstName)

	// Change the value stored in the variable.
	firstName = "Ahmed"

	fmt.Println("New Name:", firstName)

	// Update age.
	age = 21

	fmt.Println("Updated Age:", age)

	// =====================================================
	// Display Variable Types
	// =====================================================

	fmt.Println()
	fmt.Println("=== Variable Types ===")

	fmt.Printf("Type of firstName: %T\n", firstName)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of marks: %T\n", marks)
	fmt.Printf("Type of passed: %T\n", passed)
	fmt.Printf("Type of city: %T\n", city)

	// =====================================================
	// Final Summary
	// =====================================================

	fmt.Println()
	fmt.Println("=== Final Summary ===")

	fmt.Printf("Name: %s\n", firstName)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Marks: %.1f\n", marks)
	fmt.Printf("Passed: %t\n", passed)
	fmt.Printf("City: %s\n", city)
}