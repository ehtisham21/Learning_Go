package main

// Import packages that we want to use in this file.
import (
	"fmt"     // Used for printing output
	"math"    // Used for mathematical operations
	"strings" // Used for string operations
	"time"    // Used for date and time operations
)

// main() is the entry point of the program.
func main() {

	// -----------------------------------------
	// fmt package
	// -----------------------------------------
	fmt.Println("=== fmt Package ===")

	// Print a simple message
	fmt.Println("Hello from Go Imports")

	// -----------------------------------------
	// strings package
	// -----------------------------------------
	fmt.Println("\n=== strings Package ===")

	// Convert text to uppercase
	fmt.Println(strings.ToUpper("golang"))

	// Convert text to lowercase
	fmt.Println(strings.ToLower("GOLANG"))

	// Check if a string contains another string
	fmt.Println(strings.Contains("golang", "go"))

	// -----------------------------------------
	// math package
	// -----------------------------------------
	fmt.Println("\n=== math Package ===")

	// Square root of a number
	fmt.Println("Square Root of 64:", math.Sqrt(64))

	// Power calculation
	fmt.Println("2 Power 3:", math.Pow(2, 3))

	// -----------------------------------------
	// time package
	// -----------------------------------------
	fmt.Println("\n=== time Package ===")

	// Current date and time
	fmt.Println("Current Time:", time.Now())

	// -----------------------------------------
	// Variables from imported packages
	// -----------------------------------------
	fmt.Println("\n=== Package Functions ===")

	name := "ehtisham"

	// Using strings package function
	upperName := strings.ToUpper(name)

	fmt.Println("Original Name:", name)
	fmt.Println("Uppercase Name:", upperName)

	// Using math package function
	number := 25.0
	fmt.Println("Square Root:", math.Sqrt(number))

	// Using time package function
	fmt.Println("Current Year:", time.Now().Year())

	fmt.Println("\nProgram Finished Successfully")
}
