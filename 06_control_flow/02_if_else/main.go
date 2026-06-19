package main

import "fmt"

func main() {

	// =====================================
	// 1. Basic if else
	// =====================================
	age := 15

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// =====================================
	// 2. Pass or Fail
	// =====================================
	score := 90

	if score >= 50 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// =====================================
	// 3. Multiple statements
	// =====================================
	age = 20

	if age >= 18 {
		fmt.Println("Can Vote")
		fmt.Println("Can Drive")
		fmt.Println("Can Apply for Passport")
	} else {
		fmt.Println("Cannot Vote")
		fmt.Println("Cannot Drive")
	}

	// =====================================
	// 4. else if example (Grades)
	// =====================================
	marks := 85

	if marks >= 90 {
		fmt.Println("Grade A+")
	} else if marks >= 80 {
		fmt.Println("Grade A")
	} else if marks >= 70 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Fail")
	}

	// =====================================
	// 5. AND (&&)
	// =====================================
	hasID := true

	if age >= 18 && hasID {
		fmt.Println("Entry Allowed")
	} else {
		fmt.Println("Entry Denied")
	}

	// =====================================
	// 6. OR (||)
	// =====================================
	isAdmin := false
	isManager := true

	if isAdmin || isManager {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}

	// =====================================
	// 7. NOT (!)
	// =====================================
	isBlocked := false

	if !isBlocked {
		fmt.Println("User Active")
	} else {
		fmt.Println("User Blocked")
	}

	// =====================================
	// 8. Variable declaration inside if
	// =====================================
	if age := 15; age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
		fmt.Println("Age:", age)
	}

	// =====================================
	// 9. Backend Example - Login
	// =====================================
	isLoggedIn := true

	if isLoggedIn {
		fmt.Println("Show Dashboard")
	} else {
		fmt.Println("Redirect To Login")
	}

	// =====================================
	// 10. Backend Example - Request Validation
	// =====================================
	name := ""

	if name != "" {
		fmt.Println("Valid Request")
	} else {
		fmt.Println("Name Required")
	}

	// =====================================
	// 11. Backend Example - Package Version
	// =====================================
	version := "1.2.0"

	if version != "" {
		fmt.Println("Processing Package")
	} else {
		fmt.Println("Version Missing")
	}

	// =====================================
	// 12. Backend Example - User Role
	// =====================================
	role := "editor"

	if role == "admin" {
		fmt.Println("Full Access")
	} else if role == "editor" {
		fmt.Println("Limited Access")
	} else {
		fmt.Println("Read Only Access")
	}
}
