# 01 - Variable Declaration

## Introduction

Variables are used to store data in a program.

Think of a variable as a box that holds a value.

Example:

```text
Name Box  -> Ali
Age Box   -> 20
Marks Box -> 85.5
```

In Go, variables help us store and work with data.

---

# Why Do We Need Variables?

Without variables:

```go
fmt.Println("Ali")
fmt.Println(20)
```

The values are fixed.

With variables:

```go
var name string = "Ali"

fmt.Println(name)
```

We can change the value later.

```go
name = "Ahmed"
```

Now the variable stores a new value.

---

# Variable Declaration Syntax

Basic syntax:

```go
var variableName dataType = value
```

Example:

```go
var firstName string = "Ali"
```

Breakdown:

```text
var       -> keyword used to create a variable
firstName -> variable name
string    -> data type
Ali       -> value
```

---

# String Variable

A string stores text.

Example:

```go
var firstName string = "Ali"

fmt.Println(firstName)
```

Output:

```text
Ali
```

---

# Integer Variable

An integer stores whole numbers.

Example:

```go
var age int = 20

fmt.Println(age)
```

Output:

```text
20
```

---

# Float Variable

A float stores decimal numbers.

Example:

```go
var marks float64 = 85.5

fmt.Println(marks)
```

Output:

```text
85.5
```

---

# Boolean Variable

A boolean stores true or false.

Example:

```go
var passed bool = true

fmt.Println(passed)
```

Output:

```text
true
```

---

# Common Data Types

| Data Type | Example |
| --------- | ------- |
| string    | `"Ali"` |
| int       | `20`    |
| float64   | `85.5`  |
| bool      | `true`  |

Example:

```go
var firstName string = "Ali"
var age int = 20
var marks float64 = 85.5
var passed bool = true
```

---

# Printing Variables

Example:

```go
var firstName string = "Ali"

fmt.Println(firstName)
```

Output:

```text
Ali
```

We can also add labels:

```go
fmt.Println("Name:", firstName)
```

Output:

```text
Name: Ali
```

---

# Changing Variable Values

Variables can change after they are created.

Example:

```go
var firstName string = "Ali"

fmt.Println(firstName)

firstName = "Ahmed"

fmt.Println(firstName)
```

Output:

```text
Ali
Ahmed
```

Notice:

```go
firstName = "Ahmed"
```

No `var` is needed because the variable already exists.

---

# Type Inference

Go can often figure out the data type automatically.

Instead of:

```go
var city string = "Lahore"
```

We can write:

```go
var city = "Lahore"
```

Go automatically understands:

```text
Lahore → string
```

Example:

```go
var age = 20
```

Go automatically understands:

```text
20 → int
```

---

# Displaying Variable Types

We can use `%T` to see the type of a variable.

Example:

```go
fmt.Printf("%T\n", firstName)
```

Output:

```text
string
```

Example:

```go
fmt.Printf("%T\n", age)
fmt.Printf("%T\n", marks)
fmt.Printf("%T\n", passed)
```

Output:

```text
int
float64
bool
```

---

# Variable Naming Rules

## Valid Names

```go
var firstName string
var userAge int
var totalMarks int
var city string
```

---

## Invalid Names

Starts with a number:

```go
var 1name string
```

❌ Invalid

---

Contains spaces:

```go
var user name string
```

❌ Invalid

---

Contains hyphens:

```go
var user-name string
```

❌ Invalid

---

# Naming Convention in Go

Go developers usually use **camelCase**.

Good:

```go
var firstName string
var userAge int
var totalMarks int
```

Avoid:

```go
var first_name string
var TOTAL_MARKS int
```

CamelCase is the standard style used in most Go projects.

---

# Complete Example

```go
package main

import "fmt"

func main() {

	var firstName string = "Ali"
	var age int = 20
	var marks float64 = 85.5
	var passed bool = true

	fmt.Println("Name:", firstName)
	fmt.Println("Age:", age)
	fmt.Println("Marks:", marks)
	fmt.Println("Passed:", passed)

	firstName = "Ahmed"

	fmt.Println()
	fmt.Println("Updated Name:", firstName)
}
```

Output:

```text
Name: Ali
Age: 20
Marks: 85.5
Passed: true

Updated Name: Ahmed
```

---

# Variable Declaration Styles

## Full Declaration

```go
var firstName string = "Ali"
```

---

## Type Inference

```go
var firstName = "Ali"
```

---

## Declaration Only

```go
var firstName string
```

Go gives it a default value.

We will learn this in the Zero Values section.

---

# Key Points

* Variables store data.
* Use `var` to create variables.
* Variables have a name, type, and value.
* Common data types are:

  * string
  * int
  * float64
  * bool
* Variables can be updated after creation.
* Go can automatically determine data types.
* Use camelCase for variable names.
* `%T` shows the type of a variable.

---

# Summary

In this section, you learned:

* What variables are.
* Why variables are useful.
* How to declare variables using `var`.
* Common Go data types.
* How to print variables.
* How to update variable values.
* How Go automatically detects data types.
* Variable naming rules and conventions.

Variables are one of the most important concepts in Go because almost every program uses them to store and manage data.
