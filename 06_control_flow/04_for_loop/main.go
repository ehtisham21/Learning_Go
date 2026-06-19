package main

import "fmt"

func main() {

	// =====================================
	// 1. Basic for loop
	// =====================================
	fmt.Println("Basic For Loop")

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// =====================================
	// 2. Counting backwards
	// =====================================
	fmt.Println("\nCounting Backwards")

	for i := 5; i >= 1; i-- {
		fmt.Println(i)
	}

	// =====================================
	// 3. Increment by 2
	// =====================================
	fmt.Println("\nIncrement By 2")

	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// =====================================
	// 4. While-style loop
	// =====================================
	fmt.Println("\nWhile Style Loop")

	count := 1

	for count <= 5 {
		fmt.Println(count)
		count++
	}

	// =====================================
	// 5. Break example
	// =====================================
	fmt.Println("\nBreak Example")

	for i := 1; i <= 10; i++ {

		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	// =====================================
	// 6. Continue example
	// =====================================
	fmt.Println("\nContinue Example")

	for i := 1; i <= 5; i++ {

		if i == 3 {
			continue
		}

		fmt.Println(i)
	}

	// =====================================
	// 7. Traditional loop through slice
	// =====================================
	fmt.Println("\nLoop Through Slice")

	packages := []string{
		"django",
		"requests",
		"flask",
	}

	for i := 0; i < len(packages); i++ {
		fmt.Println(packages[i])
	}

	// =====================================
	// 8. Retry Logic Example
	// =====================================
	fmt.Println("\nRetry Logic")

	for attempt := 1; attempt <= 3; attempt++ {
		fmt.Println("Retry Attempt:", attempt)
	}

	// =====================================
	// 9. Backend Example
	// Processing package versions
	// =====================================
	fmt.Println("\nProcessing Packages")

	packageVersions := []string{
		"1.0.0",
		"1.1.0",
		"2.0.0",
	}

	for i := 0; i < len(packageVersions); i++ {
		fmt.Println("Processing Version:", packageVersions[i])
	}

	// =====================================
	// 10. Infinite Loop Example
	// Uncomment to test
	// =====================================

	/*
		for {
			fmt.Println("Running Forever...")
		}
	*/
}
