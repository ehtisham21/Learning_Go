package main

import "fmt"

// =====================================================
// Constant Block
// =====================================================

// Constants are created using the const keyword.
// Constants cannot be changed after creation.

const (
	appName      = "Learning Go"
	version      = "1.0"
	author       = "Ehtisham"
	maxUsers     = 100
	taxRate      = 0.15
	isProduction = false
)

func main() {

	// =====================================================
	// Display Constants
	// =====================================================

	fmt.Println("=== Application Information ===")

	fmt.Println("App Name:", appName)
	fmt.Println("Version:", version)
	fmt.Println("Author:", author)

	// =====================================================
	// Integer Constant
	// =====================================================

	fmt.Println()
	fmt.Println("=== Integer Constant ===")

	fmt.Println("Maximum Users:", maxUsers)

	// =====================================================
	// Float Constant
	// =====================================================

	fmt.Println()
	fmt.Println("=== Float Constant ===")

	fmt.Println("Tax Rate:", taxRate)

	// =====================================================
	// Boolean Constant
	// =====================================================

	fmt.Println()
	fmt.Println("=== Boolean Constant ===")

	fmt.Println("Production Mode:", isProduction)

	// =====================================================
	// Constants in Calculations
	// =====================================================

	fmt.Println()
	fmt.Println("=== Using Constants in Calculations ===")

	price := 100.0

	taxAmount := price * taxRate
	totalPrice := price + taxAmount

	fmt.Println("Price:", price)
	fmt.Println("Tax Amount:", taxAmount)
	fmt.Println("Total Price:", totalPrice)

	// =====================================================
	// Display Constant Types
	// =====================================================

	fmt.Println()
	fmt.Println("=== Constant Types ===")

	fmt.Printf("appName: %T\n", appName)
	fmt.Printf("version: %T\n", version)
	fmt.Printf("maxUsers: %T\n", maxUsers)
	fmt.Printf("taxRate: %T\n", taxRate)
	fmt.Printf("isProduction: %T\n", isProduction)

	// =====================================================
	// Variable vs Constant
	// =====================================================

	fmt.Println()
	fmt.Println("=== Variable vs Constant ===")

	var userName = "Ali"

	fmt.Println("Original User Name:", userName)

	// Variables can change
	userName = "Ahmed"

	fmt.Println("Updated User Name:", userName)

	// Constants cannot change
	// Uncommenting the line below will cause an error

	// appName = "My New App"

	// Error:
	// cannot assign to appName

	// =====================================================
	// Final Summary
	// =====================================================

	fmt.Println()
	fmt.Println("=== Final Summary ===")

	fmt.Println("Application:", appName)
	fmt.Println("Version:", version)
	fmt.Println("Maximum Users:", maxUsers)
	fmt.Println("Tax Rate:", taxRate)
}
