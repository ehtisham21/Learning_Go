package main

import "fmt"

// Function that returns multiple values.
func getUser() (string, int) {
	return "Ali", 20
}

func main() {

	// =====================================================
	// Multiple Variables of Same Type
	// =====================================================

	var firstName, lastName string

	firstName = "Ali"
	lastName = "Khan"

	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)

	// =====================================================
	// Multiple Variables with Values
	// =====================================================

	var city, country string = "Lahore", "Pakistan"

	fmt.Println()
	fmt.Println("City:", city)
	fmt.Println("Country:", country)

	// =====================================================
	// Different Types with Type Inference
	// =====================================================

	var age, passed = 20, true

	fmt.Println()
	fmt.Println("Age:", age)
	fmt.Println("Passed:", passed)

	// =====================================================
	// Variable Block
	// =====================================================

	var (
		host  = "localhost"
		port  = 8080
		debug = true
	)

	fmt.Println()
	fmt.Println("=== Server Configuration ===")

	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("Debug:", debug)

	// =====================================================
	// Multiple Variables Using :=
	// =====================================================

	name, marks := "Ahmed", 85.5

	fmt.Println()
	fmt.Println("Name:", name)
	fmt.Println("Marks:", marks)

	// =====================================================
	// Updating Multiple Variables
	// =====================================================

	name, marks = "Usman", 90.0

	fmt.Println()
	fmt.Println("Updated Name:", name)
	fmt.Println("Updated Marks:", marks)

	// =====================================================
	// Receiving Multiple Values from Function
	// =====================================================

	userName, userAge := getUser()

	fmt.Println()
	fmt.Println("=== User Information ===")

	fmt.Println("User Name:", userName)
	fmt.Println("User Age:", userAge)

	// ========================================================
	// Using Blank Identifier (_) for ignoring the second value
	// ========================================================

	onlyName, _ := getUser()

	fmt.Println()
	fmt.Println("Only Name:", onlyName)

	// =====================================================
	// Display Types
	// =====================================================

	fmt.Println()
	fmt.Println("=== Variable Types ===")

	fmt.Printf("firstName: %T\n", firstName)
	fmt.Printf("age: %T\n", age)
	fmt.Printf("passed: %T\n", passed)
	fmt.Printf("marks: %T\n", marks)
	fmt.Printf("host: %T\n", host)

	// =====================================================
	// Final Summary
	// =====================================================

	fmt.Println()
	fmt.Println("=== Final Summary ===")

	fmt.Println("Name:", onlyName)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
}
