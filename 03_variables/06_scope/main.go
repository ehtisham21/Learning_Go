package main

import "fmt"

// Package scope / global variable.
// This variable can be used by all functions in this file/package.
var appName = "Learning Go"

func showAppName() {
	// This function can access appName because appName is package scoped.
	fmt.Println("App Name from showAppName():", appName)
}

func main() {

	fmt.Println("=== Package Scope ===")

	// appName is accessible here because it is declared outside all functions.
	fmt.Println("App Name from main():", appName)

	showAppName()

	fmt.Println()
	fmt.Println("=== Function Scope ===")

	// Function scope / local variable.
	// userName can only be used inside main().
	userName := "Ali"

	fmt.Println("User Name:", userName)

	fmt.Println()
	fmt.Println("=== Block Scope ===")

	if true {
		// Block scope variable.
		// message can only be used inside this if block.
		message := "Hello from if block"

		fmt.Println(message)
	}

	// This will cause an error because message exists only inside the if block.
	// fmt.Println(message)

	fmt.Println()
	fmt.Println("=== Inner Block Can Access Outer Variable ===")

	city := "Lahore"

	if true {
		// This block can access city because city is declared outside the block.
		fmt.Println("City inside block:", city)
	}

	fmt.Println("City outside block:", city)

	fmt.Println()
	fmt.Println("=== Variable Shadowing ===")

	name := "Ali"

	fmt.Println("Name before block:", name)

	if true {
		// This creates a new variable with the same name.
		// It does not update the outer name.
		name := "Ahmed"

		fmt.Println("Name inside block:", name)
	}

	// Outer name is still Ali.
	fmt.Println("Name after block:", name)

	fmt.Println()
	fmt.Println("=== Updating Outer Variable Correctly ===")

	studentName := "Usman"

	fmt.Println("Before update:", studentName)

	if true {
		// Use = to update existing variable.
		studentName = "Hassan"

		fmt.Println("Inside block:", studentName)
	}

	fmt.Println("After update:", studentName)

	fmt.Println()
	fmt.Println("=== Local Variable Shadowing Global Variable ===")

	// This local variable has the same name as the global variable.
	// Inside main(), this local appName is used.
	appName := "My Local Go App"

	fmt.Println("Local appName:", appName)

	// showAppName() still uses the global appName.
	showAppName()

	fmt.Println()
	fmt.Println("=== Summary ===")

	fmt.Println("Package scope: variable available in whole package")
	fmt.Println("Function scope: variable available only inside function")
	fmt.Println("Block scope: variable available only inside {} block")
	fmt.Println("Shadowing: new variable hides outer variable")
}