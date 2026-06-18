package main

import "fmt"

func main() {

	fmt.Println("=== Type Inference Examples ===")

	// Integer
	var age = 25

	// Float
	var price = 99.99

	// String
	var name = "Ehtisham"

	// Boolean
	var isAdmin = true

	// Rune
	var letter = 'A'

	fmt.Println("Age:", age)
	fmt.Println("Price:", price)
	fmt.Println("Name:", name)
	fmt.Println("Is Admin:", isAdmin)
	fmt.Printf("Letter: %c\n", letter)

	fmt.Println("\n=== Checking Inferred Types ===")

	fmt.Printf("age type: %T\n", age)
	fmt.Printf("price type: %T\n", price)
	fmt.Printf("name type: %T\n", name)
	fmt.Printf("isAdmin type: %T\n", isAdmin)
	fmt.Printf("letter type: %T\n", letter)

	fmt.Println("\n=== Short Variable Declaration (:=) ===")

	city := "Islamabad"
	score := 95
	rating := 4.5
	isActive := true

	fmt.Println("City:", city)
	fmt.Println("Score:", score)
	fmt.Println("Rating:", rating)
	fmt.Println("Is Active:", isActive)

	fmt.Println("\n=== Types of Short Declarations ===")

	fmt.Printf("city type: %T\n", city)
	fmt.Printf("score type: %T\n", score)
	fmt.Printf("rating type: %T\n", rating)
	fmt.Printf("isActive type: %T\n", isActive)

	fmt.Println("\n=== Byte Slice Type Inference ===")

	data := []byte("Hello")

	fmt.Println("Data:", data)
	fmt.Printf("data type: %T\n", data)

	fmt.Println("\n=== Explicit Type vs Type Inference ===")

	var explicitAge int64 = 25
	inferredAge := 25

	fmt.Printf("explicitAge type: %T\n", explicitAge)
	fmt.Printf("inferredAge type: %T\n", inferredAge)
}
