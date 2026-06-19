# 04 - Zero Values

## Introduction

In Go, when a variable is created but no value is assigned, Go automatically gives it a default value.

This default value is called a **Zero Value**.

Example:

```go
var age int
```

Even though we did not assign a value, Go automatically sets:

```text
age = 0
```

This helps prevent errors and makes variables ready to use immediately.

---

# What is a Zero Value?

A zero value is the default value assigned by Go when a variable is declared without a value.

Example:

```go
var age int

fmt.Println(age)
```

Output:

```text
0
```

Go automatically provides the value.

---

# Why Are Zero Values Useful?

Without zero values:

```text
Create Variable
↓
Must Assign Value
↓
Use Variable
```

With zero values:

```text
Create Variable
↓
Use Variable
```

The variable always contains a valid value.

---

# Zero Values for Common Data Types

| Data Type | Zero Value |
| --------- | ---------- |
| int       | 0          |
| float64   | 0          |
| string    | ""         |
| bool      | false      |

These are the most common zero values you will use.

---

# Zero Value for int

Example:

```go
var age int

fmt.Println(age)
```

Output:

```text
0
```

Go automatically sets:

```text
age = 0
```

---

# Zero Value for float64

Example:

```go
var marks float64

fmt.Println(marks)
```

Output:

```text
0
```

Go automatically sets:

```text
marks = 0.0
```

---

# Zero Value for string

Example:

```go
var firstName string

fmt.Println(firstName)
```

Output:

```text

```

Nothing appears because the value is:

```go
""
```

which is an empty string.

To see it clearly:

```go
fmt.Printf("%q\n", firstName)
```

Output:

```text
""
```

---

# Zero Value for bool

Example:

```go
var passed bool

fmt.Println(passed)
```

Output:

```text
false
```

Go automatically sets:

```text
passed = false
```

---

# Multiple Variables with Zero Values

Example:

```go
var firstName string
var age int
var marks float64
var passed bool

fmt.Println(firstName)
fmt.Println(age)
fmt.Println(marks)
fmt.Println(passed)
```

Output:

```text

0
0
false
```

---

# Checking Zero Values

Sometimes we want to check whether a variable still has its zero value.

## String

```go
var firstName string

if firstName == "" {
	fmt.Println("Name is empty")
}
```

Output:

```text
Name is empty
```

---

## Integer

```go
var age int

if age == 0 {
	fmt.Println("Age is zero")
}
```

Output:

```text
Age is zero
```

---

## Boolean

```go
var passed bool

if !passed {
	fmt.Println("Student did not pass")
}
```

Output:

```text
Student did not pass
```

---

# Updating Zero Values

Variables can be assigned new values later.

Example:

```go
var firstName string

firstName = "Ali"

fmt.Println(firstName)
```

Output:

```text
Ali
```

The variable no longer contains the zero value.

---

# Counter Example

A common use of zero values is counters.

Example:

```go
var counter int

counter++

fmt.Println(counter)
```

Output:

```text
1
```

Why?

Because:

```text
counter starts at 0
```

Then:

```go
counter++
```

adds 1.

---

# Displaying Variable Types

We can use `%T` to display the type of a variable.

Example:

```go
var age int

fmt.Printf("%T\n", age)
```

Output:

```text
int
```

Example:

```go
var firstName string

fmt.Printf("%T\n", firstName)
```

Output:

```text
string
```

---

# Variable with Value vs Zero Value

## Variable with Assigned Value

```go
var age int = 20
```

Output:

```text
20
```

---

## Variable Without Value

```go
var age int
```

Output:

```text
0
```

Go automatically assigns the zero value.

---

# Real Examples

## User Name

```go
var firstName string
```

Default value:

```text
""
```

---

## User Age

```go
var age int
```

Default value:

```text
0
```

---

## Marks

```go
var marks float64
```

Default value:

```text
0.0
```

---

## Login Status

```go
var loggedIn bool
```

Default value:

```text
false
```

---

# Common Beginner Question

## Is Zero Value the Same as nil?

No.

For now remember:

```text
int      -> 0
float64  -> 0
string   -> ""
bool     -> false
```

Later you will learn about:

```text
nil
```

for pointers, slices, maps, and other reference types.

---

# Key Points

- Every variable in Go has a default value.
- This default value is called a Zero Value.
- Zero values are assigned automatically.
- Common zero values are:
  - int → 0
  - float64 → 0
  - string → ""
  - bool → false

- Variables can be used immediately after declaration.
- Zero values help reduce bugs and make code simpler.

---

# Summary

In this section, you learned:

- What zero values are.
- Why Go uses zero values.
- Zero values for common data types.
- How to check zero values.
- How to update variables later.
- Real-world examples using zero values.

Zero values are one of Go's most useful features because every variable always starts with a safe and predictable value.
