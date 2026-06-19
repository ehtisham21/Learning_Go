# 05_nested_loops

## What are Nested Loops?

A nested loop means a loop inside another loop.

```go
for i := 1; i <= 3; i++ {
	for j := 1; j <= 2; j++ {
		fmt.Println(i, j)
	}
}
```

The outer loop runs first, and for every outer loop iteration, the inner loop runs completely.

---

## Basic Example

```go
for i := 1; i <= 3; i++ {
	for j := 1; j <= 2; j++ {
		fmt.Println(i, j)
	}
}
```

Output:

```text
1 1
1 2
2 1
2 2
3 1
3 2
```

---

## How it Works

For `i = 1`, inner loop runs:

```text
1 1
1 2
```

For `i = 2`, inner loop runs again:

```text
2 1
2 2
```

For `i = 3`, inner loop runs again:

```text
3 1
3 2
```

So the inner loop restarts for every outer loop value.

---

## Row and Column Example

Nested loops are useful for rows and columns.

```go
for row := 1; row <= 3; row++ {
	for col := 1; col <= 4; col++ {
		fmt.Printf("(%d,%d) ", row, col)
	}

	fmt.Println()
}
```

Output:

```text
(1,1) (1,2) (1,3) (1,4)
(2,1) (2,2) (2,3) (2,4)
(3,1) (3,2) (3,3) (3,4)
```

---

## Multiplication Table

```go
for i := 1; i <= 3; i++ {
	for j := 1; j <= 3; j++ {
		fmt.Printf("%d x %d = %d\n", i, j, i*j)
	}

	fmt.Println()
}
```

Output:

```text
1 x 1 = 1
1 x 2 = 2
1 x 3 = 3

2 x 1 = 2
2 x 2 = 4
2 x 3 = 6

3 x 1 = 3
3 x 2 = 6
3 x 3 = 9
```

---

## Square Pattern

```go
for row := 1; row <= 4; row++ {
	for col := 1; col <= 4; col++ {
		fmt.Print("* ")
	}

	fmt.Println()
}
```

Output:

```text
* * * *
* * * *
* * * *
* * * *
```

---

## Triangle Pattern

```go
for row := 1; row <= 5; row++ {
	for col := 1; col <= row; col++ {
		fmt.Print("* ")
	}

	fmt.Println()
}
```

Output:

```text
* 
* * 
* * * 
* * * * 
* * * * *
```

---

## break in Nested Loops

`break` stops only the loop where it is written.

```go
for i := 1; i <= 3; i++ {
	for j := 1; j <= 5; j++ {
		if j == 3 {
			break
		}

		fmt.Println(i, j)
	}
}
```

Output:

```text
1 1
1 2
2 1
2 2
3 1
3 2
```

Here, `break` stops only the inner loop. The outer loop continues.

---

## continue in Nested Loops

`continue` skips the current iteration of the loop where it is written.

```go
for i := 1; i <= 2; i++ {
	for j := 1; j <= 4; j++ {
		if j == 2 {
			continue
		}

		fmt.Println(i, j)
	}
}
```

Output:

```text
1 1
1 3
1 4
2 1
2 3
2 4
```

Here, `j == 2` is skipped.

---

## Backend Example: Compare Packages and Vulnerabilities

```go
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
```

Output:

```text
django -> CVE-2024-1001
django -> CVE-2024-1002
flask -> CVE-2024-1001
flask -> CVE-2024-1002
```

---

## Matrix Example

A matrix is like rows and columns.

```go
matrix := [][]int{
	{1, 2, 3},
	{4, 5, 6},
}

for row := 0; row < len(matrix); row++ {
	for col := 0; col < len(matrix[row]); col++ {
		fmt.Println(matrix[row][col])
	}
}
```

Output:

```text
1
2
3
4
5
6
```

---

## Performance Note

Nested loops can become expensive.

```go
for i := 0; i < n; i++ {
	for j := 0; j < n; j++ {
	}
}
```

This runs approximately:

```text
n x n
```

If `n = 1000`, this becomes:

```text
1,000,000 iterations
```

So avoid unnecessary nested loops on large data.

---

## Common Mistakes

### Using the Same Variable Name

❌ Wrong:

```go
for i := 1; i <= 3; i++ {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
}
```

This creates confusion because the inner `i` hides the outer `i`.

✅ Better:

```go
for row := 1; row <= 3; row++ {
	for col := 1; col <= 3; col++ {
		fmt.Println(row, col)
	}
}
```

---

### Thinking break Stops All Loops

`break` only stops the current loop.

```go
for i := 1; i <= 3; i++ {
	for j := 1; j <= 5; j++ {
		if j == 3 {
			break
		}
	}
}
```

This stops only the inner loop, not the outer loop.

---

## Key Points

* Nested loop means loop inside another loop.
* The inner loop runs completely for every outer loop iteration.
* Useful for rows, columns, tables, matrices, and comparing datasets.
* `break` affects only the current loop.
* `continue` skips only the current iteration of the current loop.
* Nested loops can be expensive on large datasets.
* Use clear variable names like `row`, `col`, `pkg`, and `vuln`.

---