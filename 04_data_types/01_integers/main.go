package main

import "fmt"

func main() {

	// Integer declarations
	var age int = 25
	var temperature int = -5

	// Type inference
	var score = 100

	// Short variable declaration
	count := 10

	// Unsigned integer
	var users uint = 50

	// Zero value
	var marks int

	fmt.Println("=== Integer Examples ===")

	fmt.Println("Age:", age)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Score:", score)
	fmt.Println("Count:", count)
	fmt.Println("Users:", users)
	fmt.Println("Marks (zero value):", marks)

	fmt.Println("\n=== Arithmetic Operations ===")

	a := 10
	b := 3

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b) // Integer division
	fmt.Println("Remainder:", a%b)

	fmt.Println("\n=== Increment and Decrement ===")

	counter := 5

	counter++
	fmt.Println("After ++ :", counter)

	counter--
	fmt.Println("After -- :", counter)

	fmt.Println("\n=== Checking Types ===")

	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of users: %T\n", users)
	fmt.Printf("Type of score: %T\n", score)
}
