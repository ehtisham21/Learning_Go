# 04 - Packages

## Introduction

Packages are one of the most important concepts in Go. They help organize code into reusable and manageable units. Every Go source file belongs to a package, and every Go program is built using packages.

You can think of a package as a container that groups related functions, variables, constants, and types together.

Benefits of packages:

* Organize code logically.
* Improve code reusability.
* Make large projects easier to manage.
* Avoid naming conflicts.
* Enable code sharing across files and projects.

---

# What is a Package?

A package is a collection of Go files located in the same directory and sharing the same package name.

Every Go file must begin with a package declaration.

Example:

```go
package main
```

This tells the Go compiler that the file belongs to the `main` package.

---

# The Special `main` Package

The `main` package is a special package in Go.

When a Go program contains:

```go
package main
```

and

```go
func main() {
}
```

Go creates an executable application.

Example:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

Output:

```text
Hello, Go!
```

---

# The `main()` Function

The `main()` function is the entry point of a Go application.

Program execution always starts from the `main()` function.

Example:

```go
func main() {
	fmt.Println("Program Started")
}
```

Without a `main()` function, a program in the `main` package cannot run.

---

# Package Declaration Rules

The package declaration:

* Must be the first non-comment line in a file.
* Must appear only once in a file.
* All files in the same directory should normally use the same package name.

Correct:

```go
package main
```

Incorrect:

```go
fmt.Println("Hello")

package main
```

---

# Standard Library Packages

Go provides many built-in packages known as the Standard Library.

Common examples:

| Package       | Purpose                        |
| ------------- | ------------------------------ |
| fmt           | Input and Output               |
| strings       | String operations              |
| math          | Mathematical functions         |
| time          | Working with dates and times   |
| os            | Operating system functionality |
| strconv       | Type conversions               |
| net/http      | Building web applications      |
| encoding/json | JSON handling                  |

Example:

```go
import "fmt"
```

The `fmt` package provides functions such as:

```go
fmt.Println()
fmt.Printf()
fmt.Sprintf()
```

---

# Package-Level Variables

Variables declared outside functions belong to the package scope.

Example:

```go
package main

var appName = "Learning Go"
```

These variables can be accessed by all functions within the same package.

Example:

```go
package main

import "fmt"

var appName = "Learning Go"

func main() {
	fmt.Println(appName)
}
```

Output:

```text
Learning Go
```

---

# Exported and Unexported Identifiers

Go uses capitalization to control visibility.

## Exported (Public)

Names starting with a capital letter are accessible from other packages.

Example:

```go
func Add() {}
```

```go
var Version = "1.0"
```

```go
type User struct {}
```

These can be used outside the package.

---

## Unexported (Private)

Names starting with a lowercase letter are only accessible within the same package.

Example:

```go
func add() {}
```

```go
var version = "1.0"
```

```go
type user struct {}
```

These cannot be used by other packages.

---

# Example

```go
package main

import "fmt"

var appName = "Go Packages Example"

func Add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println("App:", appName)

	fmt.Println("Addition:", Add(10, 5))

	fmt.Println("Subtraction:", subtract(10, 5))
}
```

Output:

```text
App: Go Packages Example
Addition: 15
Subtraction: 5
```

---

# Package Naming Conventions

Go package names should be:

* Short
* Descriptive
* Lowercase
* Singular when possible

Good Examples:

```go
package utils
package config
package models
package handlers
package services
```

Bad Examples:

```go
package MyUtils
package my_utils
package my-utils
```

---

# Important Points

* Every Go file belongs to a package.
* The package declaration must be the first non-comment line.
* The `main` package creates executable programs.
* Program execution starts from `main()`.
* Go uses packages to organize code.
* Standard library packages provide built-in functionality.
* Capitalized names are exported (public).
* Lowercase names are unexported (private).
* Package-level variables are accessible throughout the package.

---

# Summary

In this section, you learned:

* What packages are.
* Why packages are important.
* The purpose of the `main` package.
* The role of the `main()` function.
* How to use standard library packages.
* Package-level variables.
* Exported vs unexported identifiers.
* Package naming conventions.

Packages are the foundation of code organization in Go. Understanding packages will make it much easier to build larger applications and work with external libraries.
