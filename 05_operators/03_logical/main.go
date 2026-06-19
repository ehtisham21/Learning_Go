package main

import "fmt"

func main() {

	// ========================================
	// AND Operator (&&)
	// ========================================
	fmt.Println("=== AND Operator (&&) ===")

	age := 25

	fmt.Println("age >= 18 && age <= 60:", age >= 18 && age <= 60)

	youngAge := 15
	fmt.Println("youngAge >= 18 && youngAge <= 60:", youngAge >= 18 && youngAge <= 60)

	// Login example
	isLoggedIn := true
	isAdmin := true

	fmt.Println("isLoggedIn && isAdmin:", isLoggedIn && isAdmin)

	// ========================================
	// OR Operator (||)
	// ========================================
	fmt.Println("\n=== OR Operator (||) ===")

	isOwner := false
	isManager := true

	fmt.Println("isOwner || isManager:", isOwner || isManager)
	fmt.Println("false || false:", false || false)

	// Access permission example
	isAdminUser := false
	isSuperUser := true

	fmt.Println("isAdminUser || isSuperUser:", isAdminUser || isSuperUser)

	// ========================================
	// NOT Operator (!)
	// ========================================
	fmt.Println("\n=== NOT Operator (!) ===")

	userLoggedIn := true
	fmt.Println("!userLoggedIn:", !userLoggedIn)

	guestLoggedIn := false
	fmt.Println("!guestLoggedIn:", !guestLoggedIn)

	// ========================================
	// Combining Multiple Conditions
	// ========================================
	fmt.Println("\n=== Combining Multiple Conditions ===")

	userAge := 25
	salary := 70000

	fmt.Println("userAge >= 18 && salary > 50000:", userAge >= 18 && salary > 50000)

	hasPermission := true
	teenAge := 16

	fmt.Println("teenAge >= 18 || hasPermission:", teenAge >= 18 || hasPermission)

	// ========================================
	// Parentheses For Readability
	// ========================================
	fmt.Println("\n=== Parentheses For Readability ===")

	employeeAge := 30
	employeeSalary := 80000

	isEligible := (employeeAge >= 18 && employeeAge <= 60) && employeeSalary > 50000

	fmt.Println("isEligible:", isEligible)

	// ========================================
	// Real-World Examples
	// ========================================
	fmt.Println("\n=== Real-World Examples ===")

	// Driving license check
	driverAge := 20
	hasLicense := true

	canDrive := driverAge >= 18 && hasLicense

	fmt.Println("Can drive:", canDrive)

	// Discount eligibility
	isMember := true
	hasCoupon := false

	canGetDiscount := isMember || hasCoupon

	fmt.Println("Can get discount:", canGetDiscount)

	// Show login button
	loggedIn := false

	showLoginButton := !loggedIn

	fmt.Println("Show login button:", showLoginButton)
}
