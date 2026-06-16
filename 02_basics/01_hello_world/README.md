# 01 - Hello World

## Objective

Learn the basic structure of a Go program and understand how Go applications start and run.

---

## Hello World Program

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

---

## Program Breakdown

### 1. Package Declaration

```go
package main
```

* Every Go file belongs to a package.
* A package is a collection of related Go files.
* `main` is a special package that creates an executable application.
* Go programs can only run when they belong to the `main` package and contain a `main()` function.

---

### 2. Import Statement

```go
import "fmt"
```

* Imports functionality from another package.
* `fmt` is part of Go's standard library.
* The `fmt` package provides functions for formatted input and output.
* Without importing `fmt`, functions such as `Println()` cannot be used.

---

### 3. Main Function

```go
func main() {
}
```

* Functions are reusable blocks of code.
* `main()` is the entry point of every Go application.
* Program execution always starts from the `main()` function.

---

### 4. Printing Output

```go
fmt.Println("Hello, World!")
```

* `Println()` prints text to the console.
* Automatically adds a new line after printing.

Example:

```go
fmt.Println("Hello")
fmt.Println("Go")
```

Output:

```text
Hello
Go
```

---

## Print vs Println

### Print

```go
fmt.Print("Hello ")
fmt.Print("World")
```

Output:

```text
Hello World
```

### Println

```go
fmt.Println("Hello")
fmt.Println("World")
```

Output:

```text
Hello
World
```

---

## Strings

Strings are text values enclosed in double quotes.

Examples:

```go
"Hello"
"Go"
"Backend Development"
```

---

## How the Program Runs

1. Go identifies `package main`.
2. Required packages are imported.
3. Go finds the `main()` function.
4. Execution starts inside `main()`.
5. `fmt.Println()` prints output to the terminal.

---

## Running the Program

Run the file directly:

```bash
go run main.go
```

Or run the current package:

```bash
go run .
```

Output:

```text
Hello, World!
```

---

## Practice Exercises

### Exercise 1

Print your name.

```go
fmt.Println("Ehtisham Butt")
```

### Exercise 2

Print multiple lines.

```go
fmt.Println("Learning Go")
fmt.Println("Day 1")
fmt.Println("Hello World")
```

### Exercise 3

Use both `Print()` and `Println()`.

```go
fmt.Print("Hello ")
fmt.Println("World")
```

---

## Key Concepts Learned

* Go program structure
* `package main`
* Importing packages
* `fmt` package
* Functions
* `main()` function
* `fmt.Print()`
* `fmt.Println()`
* Strings
* Running Go programs with `go run`

---

## Summary

The Hello World program introduces the foundation of every Go application:

* Packages organize code.
* Imports provide external functionality.
* The `main()` function is the program entry point.
* The `fmt` package is used for console output.
* Go programs are executed using `go run`.

Understanding this structure is essential before moving to more advanced Go concepts.
