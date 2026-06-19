package main

import "fmt"

func main() {

	// ========================================
	// Basic Assignment (=)
	// ========================================
	fmt.Println("=== Basic Assignment (=) ===")

	score := 100

	fmt.Println("Initial score:", score)

	score = 200

	fmt.Println("Updated score:", score)

	// ========================================
	// Add and Assign (+=)
	// ========================================
	fmt.Println("\n=== Add and Assign (+=) ===")

	points := 100

	fmt.Println("Before:", points)

	points += 10

	fmt.Println("After:", points)

	// ========================================
	// Subtract and Assign (-=)
	// ========================================
	fmt.Println("\n=== Subtract and Assign (-=) ===")

	balance := 1000

	fmt.Println("Before:", balance)

	balance -= 200

	fmt.Println("After:", balance)

	// ========================================
	// Multiply and Assign (*=)
	// ========================================
	fmt.Println("\n=== Multiply and Assign (*=) ===")

	price := 50

	fmt.Println("Before:", price)

	price *= 2

	fmt.Println("After:", price)

	// ========================================
	// Divide and Assign (/=)
	// ========================================
	fmt.Println("\n=== Divide and Assign (/=) ===")

	amount := 100

	fmt.Println("Before:", amount)

	amount /= 4

	fmt.Println("After:", amount)

	// ========================================
	// Modulus and Assign (%=)
	// ========================================
	fmt.Println("\n=== Modulus and Assign (%=) ===")

	num := 10

	fmt.Println("Before:", num)

	num %= 3

	fmt.Println("After:", num)

	// ========================================
	// Equivalent Operations
	// ========================================
	fmt.Println("\n=== Equivalent Operations ===")

	x := 10

	fmt.Println("Initial value:", x)

	x = x + 5
	fmt.Println("x = x + 5  ->", x)

	x += 5
	fmt.Println("x += 5     ->", x)

	// ========================================
	// Multiple Assignment Operations
	// ========================================
	fmt.Println("\n=== Multiple Assignment Operations ===")

	score2 := 100

	fmt.Println("Initial score:", score2)

	score2 += 10
	fmt.Println("After += 10:", score2)

	score2 -= 5
	fmt.Println("After -= 5 :", score2)

	score2 *= 2
	fmt.Println("After *= 2 :", score2)

	score2 /= 3
	fmt.Println("After /= 3 :", score2)

	// ========================================
	// Real-World Example: Bank Account
	// ========================================
	fmt.Println("\n=== Bank Account Example ===")

	accountBalance := 5000

	accountBalance += 1000

	fmt.Println("Balance after deposit:", accountBalance)

	// ========================================
	// Real-World Example: Shopping Discount
	// ========================================
	fmt.Println("\n=== Shopping Discount Example ===")

	productPrice := 1000

	productPrice -= 200

	fmt.Println("Final price:", productPrice)

	// ========================================
	// Real-World Example: Salary Increase
	// ========================================
	fmt.Println("\n=== Salary Increase Example ===")

	salary := 50000

	salary *= 2

	fmt.Println("Updated salary:", salary)

	// ========================================
	// Real-World Example: Bill Split
	// ========================================
	fmt.Println("\n=== Bill Split Example ===")

	bill := 1000

	bill /= 4

	fmt.Println("Each person pays:", bill)
}
