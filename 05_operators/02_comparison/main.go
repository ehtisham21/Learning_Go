package main

import "fmt"

func main() {

	// ========================================
	// Equal To (==)
	// ========================================
	fmt.Println("=== Equal To (==) ===")

	age := 25

	fmt.Println("age == 25:", age == 25)
	fmt.Println("10 == 20:", 10 == 20)

	// ========================================
	// Not Equal To (!=)
	// ========================================
	fmt.Println("\n=== Not Equal To (!=) ===")

	fmt.Println("age != 18:", age != 18)
	fmt.Println("10 != 10:", 10 != 10)

	// ========================================
	// Greater Than (>)
	// ========================================
	fmt.Println("\n=== Greater Than (>) ===")

	fmt.Println("20 > 10:", 20 > 10)
	fmt.Println("5 > 10:", 5 > 10)

	salary := 80000
	fmt.Println("salary > 50000:", salary > 50000)

	// ========================================
	// Less Than (<)
	// ========================================
	fmt.Println("\n=== Less Than (<) ===")

	fmt.Println("5 < 10:", 5 < 10)
	fmt.Println("20 < 10:", 20 < 10)

	// ========================================
	// Greater Than Or Equal To (>=)
	// ========================================
	fmt.Println("\n=== Greater Than Or Equal To (>=) ===")

	fmt.Println("18 >= 18:", 18 >= 18)
	fmt.Println("25 >= 18:", 25 >= 18)
	fmt.Println("15 >= 18:", 15 >= 18)

	// ========================================
	// Less Than Or Equal To (<=)
	// ========================================
	fmt.Println("\n=== Less Than Or Equal To (<=) ===")

	fmt.Println("18 <= 18:", 18 <= 18)
	fmt.Println("10 <= 20:", 10 <= 20)
	fmt.Println("25 <= 20:", 25 <= 20)

	// ========================================
	// String Comparisons
	// ========================================
	fmt.Println("\n=== String Comparisons ===")

	name := "Ehtisham"

	fmt.Println(`name == "Ehtisham":`, name == "Ehtisham")
	fmt.Println(`name != "Ali":`, name != "Ali")
	fmt.Println(`"Go" == "go":`, "Go" == "go")

	// ========================================
	// Variable Comparisons
	// ========================================
	fmt.Println("\n=== Variable Comparisons ===")

	a := 10
	b := 20

	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)

	// ========================================
	// Store Comparison Result
	// ========================================
	fmt.Println("\n=== Store Result In Variable ===")

	isAdult := age >= 18

	fmt.Println("isAdult:", isAdult)

	// ========================================
	// Real-World Examples
	// ========================================
	fmt.Println("\n=== Real-World Examples ===")

	// Age Check
	userAge := 20
	fmt.Println("Can vote:", userAge >= 18)

	// Password Check
	password := "admin123"
	fmt.Println("Password correct:", password == "admin123")

	// Stock Check
	stock := 15
	fmt.Println("Product available:", stock > 0)
}
