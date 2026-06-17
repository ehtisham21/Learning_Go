package main

import "fmt"

// This program demonstrates short variable declarations in Go.
func main() {

	// =====================================================
	// Short Variable Declaration
	// =====================================================

	// Go automatically determines the type.
	firstName := "Ali"

	fmt.Println("First Name:", firstName)

	// =====================================================
	// Integer Variable
	// =====================================================

	age := 20

	fmt.Println("Age:", age)

	// =====================================================
	// Float Variable
	// =====================================================

	marks := 85.5

	fmt.Println("Marks:", marks)

	// =====================================================
	// Boolean Variable
	// =====================================================

	passed := true

	fmt.Println("Passed:", passed)

	// =====================================================
	// Multiple Variable Declaration
	// =====================================================

	city, country := "Lahore", "Pakistan"

	fmt.Println("City:", city)
	fmt.Println("Country:", country)

	// =====================================================
	// Type Inference
	// =====================================================

	fmt.Println()
	fmt.Println("=== Variable Types ===")

	fmt.Printf("firstName: %T\n", firstName)
	fmt.Printf("age: %T\n", age)
	fmt.Printf("marks: %T\n", marks)
	fmt.Printf("passed: %T\n", passed)
	fmt.Printf("city: %T\n", city)

	// =====================================================
	// Updating Variables
	// =====================================================

	fmt.Println()
	fmt.Println("=== Updating Variables ===")

	fmt.Println("Old Name:", firstName)

	// Update existing variable using =
	firstName = "Ahmed"

	fmt.Println("New Name:", firstName)

	// Update age
	age = 21

	fmt.Println("Updated Age:", age)

	// =====================================================
	// var vs :=
	// =====================================================

	fmt.Println()
	fmt.Println("=== var vs := ===")

	// Using var and go will auto decide the type based on the input
	var appName = "Learning Go"

	// Using short declaration
	version := "1.0"

	fmt.Println("App Name:", appName)
	fmt.Println("Version:", version)

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
	fmt.Printf("Country: %s\n", country)
	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("Version: %s\n", version)
}
