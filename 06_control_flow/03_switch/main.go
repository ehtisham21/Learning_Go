package main

import "fmt"

func main() {

	// =====================================
	// 1. Basic switch
	// =====================================
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of Week")
	case "Friday":
		fmt.Println("Weekend Coming")
	default:
		fmt.Println("Normal Day")
	}

	// =====================================
	// 2. Switch with user role
	// =====================================
	role := "editor"

	switch role {
	case "admin":
		fmt.Println("Full Access")
	case "editor":
		fmt.Println("Limited Access")
	case "viewer":
		fmt.Println("Read Only Access")
	default:
		fmt.Println("Unknown Role")
	}

	// =====================================
	// 3. Switch with HTTP method
	// =====================================
	method := "POST"

	switch method {
	case "GET":
		fmt.Println("Fetch Data")
	case "POST":
		fmt.Println("Create Record")
	case "PUT":
		fmt.Println("Update Record")
	case "DELETE":
		fmt.Println("Delete Record")
	default:
		fmt.Println("Unsupported Method")
	}

	// =====================================
	// 4. Multiple values in one case
	// =====================================
	weekDay := "Sunday"

	switch weekDay {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Working Day")
	}

	// =====================================
	// 5. Switch with numbers
	// =====================================
	month := 1

	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	default:
		fmt.Println("Unknown Month")
	}

	// =====================================
	// 6. Switch without expression
	// Similar to if else if
	// =====================================
	score := 85

	switch {
	case score >= 90:
		fmt.Println("Grade A+")
	case score >= 80:
		fmt.Println("Grade A")
	case score >= 70:
		fmt.Println("Grade B")
	default:
		fmt.Println("Fail")
	}

	// =====================================
	// 7. Package registry example
	// =====================================
	registry := "pypi"

	switch registry {
	case "npm":
		fmt.Println("Node.js Package Registry")
	case "pypi":
		fmt.Println("Python Package Registry")
	case "maven":
		fmt.Println("Java Package Registry")
	case "nuget":
		fmt.Println("Dotnet Package Registry")
	default:
		fmt.Println("Unknown Package Registry")
	}

	// =====================================
	// 8. Severity example
	// =====================================
	severity := "high"

	switch severity {
	case "critical":
		fmt.Println("Immediate Action Required")
	case "high":
		fmt.Println("Fix Soon")
	case "medium":
		fmt.Println("Monitor")
	case "low":
		fmt.Println("Low Priority")
	default:
		fmt.Println("Unknown Severity")
	}

	// =====================================
	// 9. fallthrough example
	// =====================================
	level := "critical"

	switch level {
	case "critical":
		fmt.Println("Send SMS Alert")
		fallthrough
	case "high":
		fmt.Println("Create Incident Ticket")
	default:
		fmt.Println("Normal Monitoring")
	}
}
