package main

import "fmt"

func main() {

	// ========================================
	// Increment Operator (++)
	// ========================================
	fmt.Println("=== Increment Operator (++) ===")

	count := 1

	fmt.Println("Before increment:", count)

	count++

	fmt.Println("After increment:", count)

	// ========================================
	// Multiple Increments
	// ========================================
	fmt.Println("\n=== Multiple Increments ===")

	counter := 1

	fmt.Println("Start:", counter)

	counter++
	counter++
	counter++

	fmt.Println("After 3 increments:", counter)

	// ========================================
	// Decrement Operator (--)
	// ========================================
	fmt.Println("\n=== Decrement Operator (--) ===")

	lives := 5

	fmt.Println("Before decrement:", lives)

	lives--

	fmt.Println("After decrement:", lives)

	// ========================================
	// Multiple Decrements
	// ========================================
	fmt.Println("\n=== Multiple Decrements ===")

	number := 5

	fmt.Println("Start:", number)

	number--
	number--
	number--

	fmt.Println("After 3 decrements:", number)

	// ========================================
	// Counter Example
	// ========================================
	fmt.Println("\n=== Counter Example ===")

	visitors := 0

	fmt.Println("Visitors:", visitors)

	visitors++

	fmt.Println("After one new visitor:", visitors)

	// ========================================
	// Product Stock Example
	// ========================================
	fmt.Println("\n=== Product Stock Example ===")

	stock := 10

	fmt.Println("Initial stock:", stock)

	stock--

	fmt.Println("Stock after one sale:", stock)

	// ========================================
	// Game Lives Example
	// ========================================
	fmt.Println("\n=== Game Lives Example ===")

	gameLives := 3

	fmt.Println("Initial lives:", gameLives)

	gameLives--

	fmt.Println("Lives after damage:", gameLives)

	// ========================================
	// Alternative Using Assignment Operators
	// ========================================
	fmt.Println("\n=== Alternative Using Assignment Operators ===")

	score := 10

	fmt.Println("Initial score:", score)

	score += 1

	fmt.Println("After score += 1:", score)

	score -= 1

	fmt.Println("After score -= 1:", score)

	// ========================================
	// Important Go Rule
	// ========================================
	fmt.Println("\n=== Important Go Rule ===")

	x := 5

	x++

	fmt.Println("x after x++:", x)

	// In Go, x++ and x-- are statements, not expressions.
	// So these are invalid:
	//
	// y := x++          // invalid
	// fmt.Println(x++) // invalid
	// ++x              // invalid
	// --x              // invalid
	//
	// Correct way:
	x++
	fmt.Println("x after another x++:", x)

	x--
	fmt.Println("x after x--:", x)
}
