package main

import "fmt"

func main() {

	// ========================================
	// Addition
	// ========================================
	fmt.Println("=== Addition ===")

	a := 10
	b := 5

	fmt.Println("10 + 5 =", a+b)

	// String Concatenation
	firstName := "Ehtisham"
	lastName := "Butt"

	fullName := firstName + " " + lastName

	fmt.Println("Full Name:", fullName)

	// ========================================
	// Subtraction
	// ========================================
	fmt.Println("\n=== Subtraction ===")

	fmt.Println("10 - 5 =", a-b)

	// ========================================
	// Multiplication
	// ========================================
	fmt.Println("\n=== Multiplication ===")

	fmt.Println("10 * 5 =", a*b)

	// ========================================
	// Division
	// ========================================
	fmt.Println("\n=== Division ===")

	fmt.Println("10 / 5 =", a/b)

	// Integer Division
	fmt.Println("10 / 3 =", 10/3)

	// Float Division
	fmt.Println("10.0 / 3.0 =", float64(10)/float64(3))

	// ========================================
	// Modulus (Remainder)
	// ========================================
	fmt.Println("\n=== Modulus ===")

	fmt.Println("10 % 3 =", 10%3)
	fmt.Println("20 % 4 =", 20%4)

	// ========================================
	// Even / Odd Check
	// ========================================
	fmt.Println("\n=== Even/Odd Check ===")

	num1 := 8
	num2 := 7

	fmt.Println(num1, "is even:", num1%2 == 0)
	fmt.Println(num2, "is even:", num2%2 == 0)

	// ========================================
	// Arithmetic with Variables
	// ========================================
	fmt.Println("\n=== Arithmetic with Variables ===")

	price := 100
	tax := 20

	total := price + tax

	fmt.Println("Price:", price)
	fmt.Println("Tax:", tax)
	fmt.Println("Total:", total)

	// ========================================
	// Operator Precedence
	// ========================================
	fmt.Println("\n=== Operator Precedence ===")

	result1 := 10 + 5*2
	fmt.Println("10 + 5 * 2 =", result1)

	result2 := (10 + 5) * 2
	fmt.Println("(10 + 5) * 2 =", result2)

	// ========================================
	// Real-World Example
	// ========================================
	fmt.Println("\n=== Shopping Cart Example ===")

	item1 := 100
	item2 := 50

	cartTotal := item1 + item2

	fmt.Println("Item 1:", item1)
	fmt.Println("Item 2:", item2)
	fmt.Println("Cart Total:", cartTotal)
}
