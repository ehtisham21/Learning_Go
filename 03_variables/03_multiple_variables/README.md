# 03 - Multiple Variables

## Introduction

Sometimes we need more than one variable in a program.

Instead of creating variables one by one:

```go
var firstName string = "Ali"
var lastName string = "Khan"
var city string = "Lahore"
```

Go allows us to create multiple variables in a cleaner and shorter way.

---

# Why Use Multiple Variables?

Imagine a program storing user information:

```go
var firstName string = "Ali"
var age int = 20
var city string = "Lahore"
```

This works, but when many variables are needed, the code becomes repetitive.

Go provides several ways to declare multiple variables together.

---

# Multiple Variables of the Same Type

If variables have the same type, they can be declared on one line.

Example:

```go
var firstName, lastName string
```

Go creates:

```text
firstName -> string
lastName  -> string
```

Assign values:

```go
firstName = "Ali"
lastName = "Khan"
```

Print values:

```go
fmt.Println(firstName)
fmt.Println(lastName)
```

Output:

```text
Ali
Khan
```

---

# Multiple Variables with Values

Example:

```go
var city, country string = "Lahore", "Pakistan"
```

Go stores:

```text
city    -> Lahore
country -> Pakistan
```

Output:

```text
Lahore
Pakistan
```

---

# Different Types Together

Go can automatically determine different data types.

Example:

```go
var name, age = "Ali", 20
```

Go understands:

```text
name -> string
age  -> int
```

Output:

```text
Ali
20
```

---

# Variable Block

When many variables are needed, a variable block can make code cleaner.

Example:

```go
var (
	host  = "localhost"
	port  = 8080
	debug = true
)
```

This creates:

```text
host  -> string
port  -> int
debug -> bool
```

Output:

```text
localhost
8080
true
```

---

# Multiple Variables Using :=

Short variable declarations also support multiple variables.

Example:

```go
name, age := "Ali", 20
```

Go creates:

```text
name -> Ali
age  -> 20
```

Output:

```text
Ali
20
```

This is one of the most common styles used in Go.

---

# Updating Multiple Variables

Variables can be updated after creation.

Example:

```go
name, age := "Ali", 20

name, age = "Ahmed", 21
```

Output:

```text
Ahmed
21
```

Notice:

```text
Create -> :=
Update -> =
```

---

# Functions Returning Multiple Values

One of Go's unique features is that a function can return multiple values.

Example:

```go
func getUser() (string, int) {
	return "Ali", 20
}
```

Receive values:

```go
name, age := getUser()
```

Output:

```text
Ali
20
```

This feature is used heavily in Go programs.

---

# Blank Identifier (_)

Sometimes we only need one of the returned values.

Example:

```go
name, _ := getUser()
```

The underscore (`_`) means:

```text
Ignore this value
```

Output:

```text
Ali
```

The age value is ignored.

---

# Real Examples

## User Information

```go
name, age := "Ali", 20
```

---

## Location

```go
city, country := "Lahore", "Pakistan"
```

---

## Server Settings

```go
host, port := "localhost", 8080
```

---

## Coordinates

```go
x, y := 10, 20
```

---

# Different Styles

## Separate Variables

```go
var firstName string = "Ali"
var age int = 20
```

---

## Multiple Variables

```go
var firstName, age = "Ali", 20
```

---

## Short Declaration

```go
firstName, age := "Ali", 20
```

Most commonly used.

---

## Variable Block

```go
var (
	firstName = "Ali"
	age = 20
	city = "Lahore"
)
```

Useful when many variables are needed.

---

# Displaying Types

Use `%T` to display variable types.

Example:

```go
name := "Ali"

fmt.Printf("%T\n", name)
```

Output:

```text
string
```

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

# Common Mistakes

## Wrong Number of Values

Incorrect:

```go
var name, age = "Ali"
```

Error:

```text
assignment mismatch
```

Because two variables need two values.

Correct:

```go
var name, age = "Ali", 20
```

---

## Using := Instead of =

Create:

```go
name, age := "Ali", 20
```

Update:

```go
name, age = "Ahmed", 21
```

Do not use `:=` again unless creating a new variable.

---

# Key Points

* Multiple variables can be declared together.
* Variables of the same type can share one declaration.
* Different data types can be inferred automatically.
* Variable blocks help organize many variables.
* `:=` supports multiple variable declarations.
* Functions can return multiple values.
* `_` is used to ignore unwanted values.
* Multiple assignment is very common in Go.

---

# Summary

In this section, you learned:

* How to declare multiple variables.
* How to assign values to multiple variables.
* How to use variable blocks.
* How to use `:=` with multiple variables.
* How functions can return multiple values.
* How to ignore values using `_`.
* How to update multiple variables.

Multiple variable declarations make Go code cleaner, shorter, and easier to read, especially when working with related values.
