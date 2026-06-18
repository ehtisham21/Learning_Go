package main

import "fmt"

func main() {

	// Float declarations
	var price float64 = 10.99
	var temperature float32 = 36.5

	// Type inference
	var salary = 50000.75

	// Short variable declaration
	rating := 4.8

	// Zero value
	var amount float64

	fmt.Println("=== Float Examples ===")

	fmt.Println("Price:", price)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Salary:", salary)
	fmt.Println("Rating:", rating)
	fmt.Println("Amount (zero value):", amount)

	fmt.Println("\n=== Arithmetic Operations ===")

	a := 10.5
	b := 2.5

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)

	fmt.Println("\n=== Integer vs Float Division ===")

	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10.0 / 3.0 =", 10.0/3.0)

	fmt.Println("\n=== Type Checking ===")

	fmt.Printf("Type of price: %T\n", price)
	fmt.Printf("Type of temperature: %T\n", temperature)
	fmt.Printf("Type of salary: %T\n", salary)

	fmt.Println("\n=== Float Formatting ===")

	pi := 3.14159265359

	fmt.Printf("Default: %f\n", pi)
	fmt.Printf("2 decimal places: %.2f\n", pi)
	fmt.Printf("4 decimal places: %.4f\n", pi)

	fmt.Println("\n=== Float Precision Example ===")

	fmt.Println("0.1 + 0.2 =", 0.1+0.2)

	fmt.Println("\n=== Type Conversion ===")

	age := 25
	price2 := 10.99

	fmt.Println("float64(age) + price2 =", float64(age)+price2)
}
