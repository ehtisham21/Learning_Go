package main

import "fmt"

func main() {

	// Boolean declarations
	var isActive bool = true

	// Type inference
	var isAdmin = false

	// Short variable declaration
	isLoggedIn := true

	// Zero value
	var hasPermission bool

	fmt.Println("=== Boolean Examples ===")

	fmt.Println("isActive:", isActive)
	fmt.Println("isAdmin:", isAdmin)
	fmt.Println("isLoggedIn:", isLoggedIn)
	fmt.Println("hasPermission (zero value):", hasPermission)

	fmt.Println("\n=== Comparison Operators ===")

	fmt.Println("10 == 10 :", 10 == 10)
	fmt.Println("10 != 5  :", 10 != 5)
	fmt.Println("20 > 10  :", 20 > 10)
	fmt.Println("5 < 10   :", 5 < 10)
	fmt.Println("10 >= 10 :", 10 >= 10)
	fmt.Println("5 <= 10  :", 5 <= 10)

	fmt.Println("\n=== Logical Operators ===")

	age := 20
	hasID := true

	fmt.Println("age >= 18 && hasID :", age >= 18 && hasID)
	fmt.Println("true && false      :", true && false)
	fmt.Println("true || false      :", true || false)
	fmt.Println("!true              :", !true)

	fmt.Println("\n=== If Statements ===")

	if isLoggedIn {
		fmt.Println("Welcome! User is logged in.")
	}

	if age >= 18 {
		fmt.Println("User is an adult.")
	}

	fmt.Println("\n=== Type Checking ===")

	fmt.Printf("Type of isActive: %T\n", isActive)
	fmt.Printf("Type of isAdmin: %T\n", isAdmin)
}
