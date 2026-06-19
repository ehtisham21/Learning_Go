# 02 - Short Variable Declaration

## Introduction

In Go, variables can be created using the `var` keyword:

```go
var firstName string = "Ali"
```

But Go also provides a shorter and more commonly used way:

```go
firstName := "Ali"
```

This is called a **Short Variable Declaration**.

It is one of the most frequently used features in Go.

---

# What is a Short Variable Declaration?

A short variable declaration uses the `:=` operator to:

1. Create a new variable.
2. Assign a value.
3. Let Go automatically determine the data type.

Example:

```go
firstName := "Ali"
```

Go automatically understands:

```text
Variable Name: firstName
Data Type: string
Value: Ali
```

---

# Syntax

Basic syntax:

```go
variableName := value
```

Example:

```go
age := 20
```

Go automatically determines that:

```text
age is an int
```

---

# Why Use `:=`?

The normal declaration:

```go
var firstName string = "Ali"
```

works perfectly.

But this:

```go
firstName := "Ali"
```

is shorter and easier to read.

Both create the same variable.

---

# Examples

## String Variable

```go
firstName := "Ali"

fmt.Println(firstName)
```

Output:

```text
Ali
```

---

## Integer Variable

```go
age := 20

fmt.Println(age)
```

Output:

```text
20
```

---

## Float Variable

```go
marks := 85.5

fmt.Println(marks)
```

Output:

```text
85.5
```

Go automatically uses:

```text
float64
```

---

## Boolean Variable

```go
passed := true

fmt.Println(passed)
```

Output:

```text
true
```

Go automatically uses:

```text
bool
```

---

# Type Inference

One of the biggest advantages of `:=` is type inference.

Type inference means Go automatically determines the data type.

Example:

```go
city := "Lahore"
```

Go understands:

```text
city is string
```

---

Example:

```go
age := 20
```

Go understands:

```text
age is int
```

---

Example:

```go
price := 99.99
```

Go understands:

```text
price is float64
```

You do not need to write the type manually.

---

# Multiple Variable Declaration

You can create multiple variables in one line.

Example:

```go
name, age := "Ali", 20
```

Go creates:

```text
name = Ali
age = 20
```

Example:

```go
city, country := "Lahore", "Pakistan"
```

Output:

```text
Lahore
Pakistan
```

---

# Updating Variables

Variables created with `:=` can be updated later.

Example:

```go
firstName := "Ali"

firstName = "Ahmed"
```

Notice:

```text
Create → :=
Update → =
```

---

Correct:

```go
firstName := "Ali"

firstName = "Ahmed"
```

---

Incorrect:

```go
firstName := "Ali"

firstName := "Ahmed"
```

This causes an error because the variable already exists.

Error:

```text
no new variables on left side of :=
```

---

# Difference Between := and =

## :=

Used to create a new variable.

Example:

```go
age := 20
```

Meaning:

```text
Create variable age
Store value 20
```

---

## =

Used to update an existing variable.

Example:

```go
age = 21
```

Meaning:

```text
Update age
Change value from 20 to 21
```

---

# Where Can We Use := ?

Short variable declarations can only be used inside functions.

Example:

```go
func main() {

	firstName := "Ali"

	fmt.Println(firstName)
}
```

This works.

---

Example:

```go
firstName := "Ali"

func main() {
}
```

This does not work.

Error:

```text
non-declaration statement outside function body
```

---

# var vs :=

## Using var

```go
var firstName string = "Ali"
```

---

## Using Type Inference

```go
var firstName = "Ali"
```

---

## Using Short Declaration

```go
firstName := "Ali"
```

All three create the same variable.

---

# Comparison

| Feature         | var | :=  |
| --------------- | --- | --- |
| Create Variable | Yes | Yes |
| Type Inference  | Yes | Yes |
| Explicit Type   | Yes | No  |
| Package Level   | Yes | No  |
| Inside Function | Yes | Yes |
| Shorter Syntax  | No  | Yes |

---

# When Should We Use var?

Use `var` when:

### No Value Yet

```go
var firstName string
```

---

### Package Level Variable

```go
var appName = "Learning Go"
```

---

### Specific Type Required

```go
var age int64 = 20
```

---

# When Should We Use := ?

Use `:=` when:

```go
firstName := "Ali"
age := 20
city := "Lahore"
```

inside functions.

This is the most common use case in Go.

---

# Display Variable Types

We can use `%T` to see a variable's type.

Example:

```go
firstName := "Ali"

fmt.Printf("%T\n", firstName)
```

Output:

```text
string
```

---

Example:

```go
age := 20

fmt.Printf("%T\n", age)
```

Output:

```text
int
```

---

# Complete Example

```go
package main

import "fmt"

func main() {

	firstName := "Ali"
	age := 20
	city := "Lahore"
	passed := true
	marks := 85.5

	fmt.Println("Name:", firstName)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Passed:", passed)
	fmt.Println("Marks:", marks)

	firstName = "Ahmed"
	age = 21

	fmt.Println()
	fmt.Println("Updated Name:", firstName)
	fmt.Println("Updated Age:", age)
}
```

Output:

```text
Name: Ali
Age: 20
City: Lahore
Passed: true
Marks: 85.5

Updated Name: Ahmed
Updated Age: 21
```

---

# Common Mistakes

## Using := Twice

Incorrect:

```go
name := "Ali"

name := "Ahmed"
```

Error:

```text
no new variables on left side of :=
```

---

Correct:

```go
name := "Ali"

name = "Ahmed"
```

---

## Using := Outside Functions

Incorrect:

```go
name := "Ali"

func main() {
}
```

Error:

```text
non-declaration statement outside function body
```

---

# Key Points

* `:=` is called a short variable declaration.
* It creates a variable and assigns a value.
* Go automatically determines the type.
* `:=` can only be used inside functions.
* Use `=` to update an existing variable.
* `:=` is shorter than `var`.
* Most Go developers use `:=` for local variables.

---

# Summary

In this section, you learned:

* What a short variable declaration is.
* How to use `:=`.
* How Go automatically determines data types.
* The difference between `:=` and `=`.
* The difference between `var` and `:=`.
* Where short declarations can be used.
* Common mistakes beginners make.

Short variable declarations are one of the most common ways to create variables in Go and are heavily used in real-world Go applications.
