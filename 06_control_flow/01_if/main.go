package main

import "fmt"

func main() {

	// =====================================
	// 1. Basic if statement
	// =====================================
	age := 20

	if age >= 18 {
		fmt.Println("Adult")
	}

	// =====================================
	// 2. Comparison operators
	// =====================================
	score := 90

	if score >= 80 {
		fmt.Println("Passed")
	}

	// =====================================
	// 3. Boolean variable
	// =====================================
	isAdmin := true

	if isAdmin {
		fmt.Println("Access Granted")
	}

	// =====================================
	// 4. Logical AND (&&)
	// =====================================
	hasID := true

	if age >= 18 && hasID {
		fmt.Println("Entry Allowed")
	}

	// =====================================
	// 5. Logical OR (||)
	// =====================================
	isManager := true

	if isAdmin || isManager {
		fmt.Println("Admin or Manager Access")
	}

	// =====================================
	// 6. Logical NOT (!)
	// =====================================
	isBlocked := false

	if !isBlocked {
		fmt.Println("User Active")
	}

	// =====================================
	// 7. Multiple statements inside if
	// =====================================
	if age >= 18 {
		fmt.Println("Can Vote")
		fmt.Println("Can Drive")
		fmt.Println("Can Apply for Passport")
	}

	// =====================================
	// 8. Variable declaration inside if
	// =====================================
	if marks := 85; marks >= 80 {
		fmt.Println("Excellent Score")
		fmt.Println("Marks:", marks)
	}

	// =====================================
	// 9. String validation example
	// =====================================
	name := "Ehtisham"

	if name != "" {
		fmt.Println("Valid Name:", name)
	}

	// =====================================
	// 10. Backend-style example
	// =====================================
	version := "1.2.0"

	if version != "" {
		fmt.Println("Processing Package Version:", version)
	}
}
