package main

import "fmt"

func main() {

	// =====================================
	// 1. Basic break example
	// =====================================
	fmt.Println("Basic break example")

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	// =====================================
	// 2. Basic continue example
	// =====================================
	fmt.Println("\nBasic continue example")

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Println(i)
	}

	// =====================================
	// 3. Search example using break
	// =====================================
	fmt.Println("\nSearch example using break")

	users := []string{"Ali", "Ahmed", "Sara"}
	target := "Ahmed"

	for _, user := range users {
		if user == target {
			fmt.Println("User found:", user)
			break
		}

		fmt.Println("Checking:", user)
	}

	// =====================================
	// 4. Skip invalid data using continue
	// =====================================
	fmt.Println("\nSkip invalid data using continue")

	versions := []string{"1.0.0", "", "2.0.0", "", "3.0.0"}

	for _, version := range versions {
		if version == "" {
			continue
		}

		fmt.Println("Processing version:", version)
	}

	// =====================================
	// 5. break in nested loop
	// break stops only the inner loop here
	// =====================================
	fmt.Println("\nbreak in nested loop")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				break
			}

			fmt.Println(i, j)
		}
	}

	// =====================================
	// 6. continue in nested loop
	// continue skips only current inner-loop iteration
	// =====================================
	fmt.Println("\ncontinue in nested loop")

	for i := 1; i <= 2; i++ {
		for j := 1; j <= 4; j++ {
			if j == 2 {
				continue
			}

			fmt.Println(i, j)
		}
	}

	// =====================================
	// 7. Labeled break
	// stops both outer and inner loops
	// =====================================
	fmt.Println("\nLabeled break example")

OuterBreak:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				break OuterBreak
			}

			fmt.Println(i, j)
		}
	}

	// =====================================
	// 8. Labeled continue
	// skips to next outer-loop iteration
	// =====================================
	fmt.Println("\nLabeled continue example")

OuterContinue:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue OuterContinue
			}

			fmt.Println(i, j)
		}
	}

	// =====================================
	// 9. Backend example
	// Stop retrying after success
	// =====================================
	fmt.Println("\nRetry until success")

	successAt := 3

	for attempt := 1; attempt <= 5; attempt++ {
		fmt.Println("Attempt:", attempt)

		if attempt == successAt {
			fmt.Println("Success, stopping retries")
			break
		}
	}

	// =====================================
	// 10. Backend example
	// Skip invalid packages
	// =====================================
	fmt.Println("\nSkip invalid packages")

	packages := []string{"django", "", "flask", "", "requests"}

	for _, pkg := range packages {
		if pkg == "" {
			continue
		}

		fmt.Println("Valid package:", pkg)
	}
}
