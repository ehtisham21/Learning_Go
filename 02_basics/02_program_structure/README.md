# 02 - Program Structure

## Objective

Learn how a Go source file is organized and understand the role of each section in a Go program.

A Go program follows a predictable structure that makes code easy to read, maintain, and scale.

---

# Basic Program Structure

```go
package main

import "fmt"

const AppName = "Learning Go"

var version = "1.0"

func main() {
	fmt.Println(AppName)
	printVersion()
}

func printVersion() {
	fmt.Println(version)
}
```

---

# Structure Overview

A typical Go file may contain:

```text
1. Package Declaration
2. Import Statements
3. Constants
4. Variables
5. Functions
6. Main Function
```

Example:

```go
package main

import "fmt"

const AppName = "Learning Go"

var version = "1.0"

func main() {
	fmt.Println(AppName)
}

func printVersion() {
	fmt.Println(version)
}
```

---

# 1. Package Declaration

```go
package main
```

## What is a Package?

A package is a collection of related Go files.

Every Go source file must belong to a package.

Examples:

```go
package main
```

```go
package calculator
```

```go
package user
```

## Why is `main` Special?

The `main` package tells Go:

> This package should be compiled as an executable application.

To run a Go application, two things are required:

```go
package main
```

and

```go
func main()
```

Without them, the code is treated as a library.

## Rules

- Must be the first non-comment line in the file.
- Every `.go` file must have a package declaration.
- All files in the same folder should generally belong to the same package.

---

# 2. Import Statements

```go
import "fmt"
```

## Purpose

Imports allow your program to use code from other packages.

Without imports, your code can only use built-in language features.

## Single Import

```go
import "fmt"
```

## Multiple Imports

```go
import (
	"fmt"
	"strings"
	"time"
)
```

## Common Standard Library Packages

| Package  | Purpose                       |
| -------- | ----------------------------- |
| fmt      | Input and output              |
| strings  | String manipulation           |
| time     | Date and time                 |
| math     | Mathematical operations       |
| os       | Operating system interactions |
| net/http | HTTP servers and clients      |

---

# 3. Constants

```go
const AppName = "Learning Go"
```

## What is a Constant?

A constant stores a value that cannot change during program execution.

Example:

```go
const Pi = 3.14159
```

```go
const CompanyName = "ComplyVigil"
```

## Invalid Example

```go
const AppName = "Learning Go"

AppName = "New Name"
```

This will cause a compilation error because constants are immutable.

## When to Use Constants

Use constants for values that never change:

- Application name
- API version
- Mathematical values
- Fixed configuration values

---

# 4. Variables

```go
var version = "1.0"
```

## What is a Variable?

A variable stores data that can change while the program runs.

Example:

```go
var name = "Ehtisham"
```

Later:

```go
name = "Ali"
```

The value can be updated.

## Global Variables

Variables declared outside functions are called global variables.

Example:

```go
var version = "1.0"

func main() {
	fmt.Println(version)
}
```

Global variables can be accessed by all functions in the same package.

---

# 5. Functions

```go
func printVersion() {
	fmt.Println(version)
}
```

## What is a Function?

A function is a reusable block of code.

Instead of repeating code, we place it inside a function and call it when needed.

Example:

```go
func greet() {
	fmt.Println("Hello")
}
```

Calling the function:

```go
greet()
```

Output:

```text
Hello
```

## Benefits

- Reusability
- Cleaner code
- Easier maintenance
- Better organization

---

# 6. Main Function

```go
func main() {
	fmt.Println(AppName)
	printVersion()
}
```

## Purpose

The `main()` function is the starting point of every Go application.

When you run:

```bash
go run .
```

Go looks for:

```go
func main()
```

and starts execution there.

## Important Rule

There can only be one `main()` function in a package.

---

# Program Execution Flow

Consider:

```go
package main

import "fmt"

func main() {
	fmt.Println("Program Started")
	greet()
	fmt.Println("Program Ended")
}

func greet() {
	fmt.Println("Hello")
}
```

Execution order:

```text
1. main() starts
2. Print "Program Started"
3. Call greet()
4. Print "Hello"
5. Return to main()
6. Print "Program Ended"
```

Output:

```text
Program Started
Hello
Program Ended
```

---

# Python vs Go Execution

## Python

```python
print("Start")
print("End")
```

Execution flows from top to bottom.

## Go

```go
func main() {
	fmt.Println("Start")
}
```

Execution always begins inside `main()`.

Functions are only executed when called.

---

# Code Blocks and Braces

Go uses curly braces to define blocks of code.

Example:

```go
func main() {
	fmt.Println("Hello")
}
```

Everything inside the braces belongs to the function.

## Python Comparison

Python:

```python
def greet():
    print("Hello")
```

Go:

```go
func greet() {
	fmt.Println("Hello")
}
```

---

# Go Formatting Standards

Go has a standard formatting style.

Preferred:

```go
func main() {
	fmt.Println("Hello")
}
```

Avoid:

```go
func main(){
fmt.Println("Hello")
}
```

Go provides a formatting tool:

```bash
go fmt
```

which automatically formats code according to Go standards.

---

# Complete Example

```go
package main

import "fmt"

const AppName = "Learning Go"

var version = "1.0"

func main() {
	fmt.Println("Application:", AppName)
	printVersion()
}

func printVersion() {
	fmt.Println("Version:", version)
}
```

Output:

```text
Application: Learning Go
Version: 1.0
```

---

# Running the Program

Run the current package:

```bash
go run .
```

Or run a specific file:

```bash
go run main.go
```

Build an executable:

```bash
go build
```

---

# Summary

A Go program follows a clear and predictable structure:

1. Package declaration
2. Imports
3. Constants
4. Variables
5. Functions
6. Main function
