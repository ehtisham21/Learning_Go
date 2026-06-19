package main

import "fmt"

func main() {

	// =====================================
	// 1. Basic Nested Loop
	// =====================================
	fmt.Println("Basic Nested Loop")

	for i := 1; i <= 3; i++ {

		for j := 1; j <= 2; j++ {
			fmt.Println(i, j)
		}
	}

	// =====================================
	// 2. Row and Column Example
	// =====================================
	fmt.Println("\nRows and Columns")

	for row := 1; row <= 3; row++ {

		for col := 1; col <= 4; col++ {
			fmt.Printf("(%d,%d) ", row, col)
		}

		fmt.Println()
	}

	// =====================================
	// 3. Multiplication Table
	// =====================================
	fmt.Println("\nMultiplication Table")

	for i := 1; i <= 3; i++ {

		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}

		fmt.Println()
	}

	// =====================================
	// 4. Square Pattern
	// =====================================
	fmt.Println("Square Pattern")

	for row := 1; row <= 4; row++ {

		for col := 1; col <= 4; col++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}

	// =====================================
	// 5. Triangle Pattern
	// =====================================
	fmt.Println("\nTriangle Pattern")

	for row := 1; row <= 5; row++ {

		for col := 1; col <= row; col++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}

	// =====================================
	// 6. break in Nested Loop
	// =====================================
	fmt.Println("\nBreak Example")

	for i := 1; i <= 3; i++ {

		for j := 1; j <= 5; j++ {

			if j == 3 {
				break
			}

			fmt.Println(i, j)
		}
	}

	// =====================================
	// 7. continue in Nested Loop
	// =====================================
	fmt.Println("\nContinue Example")

	for i := 1; i <= 2; i++ {

		for j := 1; j <= 4; j++ {

			if j == 2 {
				continue
			}

			fmt.Println(i, j)
		}
	}

	// =====================================
	// 8. Backend Example
	// Compare Packages and Vulnerabilities
	// =====================================
	fmt.Println("\nPackages and Vulnerabilities")

	packages := []string{
		"django",
		"flask",
	}

	vulnerabilities := []string{
		"CVE-2024-1001",
		"CVE-2024-1002",
	}

	for _, pkg := range packages {

		for _, vuln := range vulnerabilities {
			fmt.Printf("%s -> %s\n", pkg, vuln)
		}
	}

	// =====================================
	// 9. Matrix Example
	// =====================================
	fmt.Println("\nMatrix Values")

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for row := 0; row < len(matrix); row++ {

		for col := 0; col < len(matrix[row]); col++ {
			fmt.Println(matrix[row][col])
		}
	}
}
