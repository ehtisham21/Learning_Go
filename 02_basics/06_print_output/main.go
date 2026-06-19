package main

// fmt package is used for printing output.
import "fmt"

// main() is the starting point of every Go program.
func main() {

	// =====================================================
	// fmt.Print()
	// =====================================================

	// Print() prints text without adding a new line.
	fmt.Print("Hello")
	fmt.Print("Go")

	// Add a manual new line for better output formatting.
	fmt.Print("\n\n")

	// =====================================================
	// fmt.Println()
	// =====================================================

	// Println() prints text and automatically adds a new line.
	fmt.Println("=== Println Example ===")

	// Print multiple values.
	// Go automatically adds spaces between values.
	fmt.Println("Name:", "Ehtisham")
	fmt.Println("Age:", 25)

	fmt.Println()

	// =====================================================
	// Variables
	// =====================================================

	name := "Ehtisham"
	age := 25
	price := 99.99
	isActive := true

	// =====================================================
	// fmt.Printf()
	// =====================================================

	// Printf() is used when we want formatted output.
	fmt.Println("=== Printf Example ===")

	// %s = string
	fmt.Printf("Name: %s\n", name)

	// %d = integer
	fmt.Printf("Age: %d\n", age)

	// %f = float
	fmt.Printf("Price: %f\n", price)

	// %.2f = float with 2 decimal places
	fmt.Printf("Price (2 decimals): %.2f\n", price)

	// %t = boolean
	fmt.Printf("Active: %t\n", isActive)

	fmt.Println()

	// =====================================================
	// %v (Default Value)

	// %v means:
	// 	Print the value in its default format.
	// =====================================================

	fmt.Println("=== %v Example ===")

	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Age: %v\n", age)
	fmt.Printf("Price: %v\n", price)

	fmt.Println()

	// =====================================================
	// %T (Type)
	// =====================================================

	fmt.Println("=== %T Example ===")

	// Shows the datatype of a variable.
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of price: %T\n", price)
	fmt.Printf("Type of active: %T\n", isActive)

	fmt.Println()

	// =====================================================
	// New Line (\n)
	// =====================================================

	fmt.Println("=== New Line Example ===")

	// \n creates a new line.
	fmt.Printf("Line 1\nLine 2\nLine 3\n")

	fmt.Println()

	// =====================================================
	// Tab (\t)
	// =====================================================

	fmt.Println("=== Tab Example ===")

	// \t creates tab spacing.
	fmt.Printf("Name\tAge\n")
	fmt.Printf("Ali\t25\n")
	fmt.Printf("Ahmed\t30\n")

	fmt.Println()

	// =====================================================
	// Percentage Sign %%
	// =====================================================

	fmt.Println("=== Percent Example ===")

	// %% prints a literal percent sign.
	fmt.Printf("Progress: 80%% Complete\n")

	fmt.Println()

	// =====================================================
	// fmt.Sprintf()
	// =====================================================

	fmt.Println("=== Sprintf Example ===")

	// Sprintf() formats a string and returns it.
	// It does NOT print anything.

	message := fmt.Sprintf(
		"User %s is %d years old",
		name,
		age,
	)

	fmt.Println(message)

	fmt.Println()

	// =====================================================
	// fmt.Sprintln()
	// =====================================================

	fmt.Println("=== Sprintln Example ===")

	// Sprintln() returns a formatted string.
	output := fmt.Sprintln(
		"Welcome",
		name,
		"to Go",
	)

	fmt.Print(output)

	fmt.Println()

	// =====================================================
	// Final Summary
	// =====================================================

	fmt.Println("=== Summary ===")

	fmt.Println("Print()    -> Prints without newline")
	fmt.Println("Println()  -> Prints with newline")
	fmt.Println("Printf()   -> Prints formatted output")
	fmt.Println("Sprintf()  -> Returns formatted string")
	fmt.Println("Sprintln() -> Returns string with spaces")
}
