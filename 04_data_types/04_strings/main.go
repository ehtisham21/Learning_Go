package main

import "fmt"

func main() {

	// String declarations
	var firstName string = "Ehtisham"

	// Type inference
	var lastName = "Butt"

	// Short variable declaration
	city := "Islamabad"

	// Zero value
	var country string

	fmt.Println("=== String Examples ===")

	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("City:", city)
	fmt.Println("Country (zero value):", country)

	fmt.Println("\n=== Type Checking ===")

	fmt.Printf("Type of firstName: %T\n", firstName)

	fmt.Println("\n=== String Length ===")

	name := "Ehtisham"

	fmt.Println("Name:", name)
	fmt.Println("Length:", len(name))

	fmt.Println("\n=== String Concatenation ===")

	fullName := firstName + " " + lastName

	fmt.Println("Full Name:", fullName)

	fmt.Println("\n=== Multi-line String ===")

	message := `Hello
				Welcome to Go
				Learning Strings`

	fmt.Println(message)

	fmt.Println("\n=== Escape Characters ===")

	fmt.Println("Hello\nWorld")
	fmt.Println("Name:\tEhtisham")
	fmt.Println("He said \"Hello\"")
	fmt.Println("C:\\Users\\Admin")

	fmt.Println("\n=== String Comparison ===")

	fmt.Println("\"go\" == \"go\" :", "go" == "go")
	fmt.Println("\"go\" == \"python\" :", "go" == "python")
	fmt.Println("\"go\" != \"python\" :", "go" != "python")

	fmt.Println("\n=== Accessing Characters ===")

	language := "Go"

	fmt.Println("language[0]:", language[0])
	fmt.Println("language[1]:", language[1])

	fmt.Println("\n=== String Formatting ===")

	age := 25

	fmt.Printf("Name: %s, Age: %d\n", fullName, age)
}
