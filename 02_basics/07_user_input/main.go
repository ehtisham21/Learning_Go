package main

import "fmt"

// This program shows different ways to take user input in Go.
func main() {

	// --------------------------------------------------
	// 1. String input
	// --------------------------------------------------

	var name string

	fmt.Print("Enter your first name: ")
	fmt.Scan(&name)

	fmt.Println("Hello", name)

	// --------------------------------------------------
	// 2. Integer input
	// --------------------------------------------------

	var age int

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Your age is:", age)

	// --------------------------------------------------
	// 3. Float input
	// --------------------------------------------------

	var marks float64

	fmt.Print("Enter your marks: ")
	fmt.Scan(&marks)

	fmt.Println("Your marks are:", marks)

	// --------------------------------------------------
	// 4. Boolean input
	// --------------------------------------------------

	var passed bool

	fmt.Print("Did you pass? Enter true or false: ")
	fmt.Scan(&passed)

	fmt.Println("Passed:", passed)

	// --------------------------------------------------
	// 5. Multiple inputs
	// --------------------------------------------------

	var city string
	var country string

	fmt.Print("Enter city and country: ")
	fmt.Scan(&city, &country)

	fmt.Println("City:", city)
	fmt.Println("Country:", country)

	// --------------------------------------------------
	// 6. Two numbers and sum
	// --------------------------------------------------

	var num1 int
	var num2 int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&num1, &num2)

	sum := num1 + num2

	fmt.Println("First Number:", num1)
	fmt.Println("Second Number:", num2)
	fmt.Println("Sum:", sum)

	// --------------------------------------------------
	// 7. Using Scanln()
	// --------------------------------------------------

	var favoriteLanguage string

	fmt.Print("Enter your favorite programming language: ")
	fmt.Scanln(&favoriteLanguage)

	fmt.Println("Favorite Language:", favoriteLanguage)

	// --------------------------------------------------
	// 8. Using Scanf()
	// --------------------------------------------------

	var studentName string
	var studentAge int

	fmt.Print("Enter student name and age: ")
	fmt.Scanf("%s %d", &studentName, &studentAge)

	fmt.Println("Student Name:", studentName)
	fmt.Println("Student Age:", studentAge)

	// --------------------------------------------------
	// 9. Final formatted output
	// --------------------------------------------------

	fmt.Println()
	fmt.Println("===== Final Summary =====")

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Marks: %.2f\n", marks)
	fmt.Printf("Passed: %t\n", passed)
	fmt.Printf("City: %s\n", city)
	fmt.Printf("Country: %s\n", country)
	fmt.Printf("Sum of numbers: %d\n", sum)
}