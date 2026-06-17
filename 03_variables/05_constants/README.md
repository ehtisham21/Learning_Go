# 05 - Constants

## Introduction

A constant is a value that cannot be changed after it is created.

In Go, constants are created using the `const` keyword.

Example:

```go
const appName = "Learning Go"
```

The value of `appName` will remain the same throughout the program.

---

# What is a Constant?

A constant stores a fixed value.

Example:

```go
const appName = "Learning Go"
```

Output:

```text
Learning Go
```

Unlike variables, constants cannot be modified.

---

# Why Do We Need Constants?

Some values should never change.

Examples:

```text
Application Name
Version Number
Tax Rate
Pi Value
Days in a Week
Maximum Users
```

These are good examples of constants.

---

# Variable vs Constant

## Variable

```go
var age = 20

age = 21
```

Output:

```text
21
```

Variables can change.

---

## Constant

```go
const maxUsers = 100

maxUsers = 200
```

Error:

```text
cannot assign to maxUsers
```

Constants cannot change.

---

# Constant Syntax

Basic syntax:

```go
const constantName = value
```

Example:

```go
const appName = "Learning Go"
```

---

# String Constant

Example:

```go
const appName = "Learning Go"

fmt.Println(appName)
```

Output:

```text
Learning Go
```

---

# Integer Constant

Example:

```go
const maxUsers = 100

fmt.Println(maxUsers)
```

Output:

```text
100
```

---

# Float Constant

Example:

```go
const taxRate = 0.15

fmt.Println(taxRate)
```

Output:

```text
0.15
```

---

# Boolean Constant

Example:

```go
const isProduction = false

fmt.Println(isProduction)
```

Output:

```text
false
```

---

# Type Inference

Just like variables, Go can automatically determine the type of a constant.

Example:

```go
const appName = "Learning Go"
```

Go automatically understands:

```text
appName → string
```

---

Example:

```go
const maxUsers = 100
```

Go automatically understands:

```text
maxUsers → integer
```

---

# Explicit Type

We can also specify the type.

Example:

```go
const appName string = "Learning Go"
```

Example:

```go
const maxUsers int = 100
```

Both approaches are valid.

---

# Multiple Constants

We can create multiple constants on one line.

Example:

```go
const firstName, country = "Ali", "Pakistan"
```

Output:

```text
Ali
Pakistan
```

---

# Constant Block

When many constants are needed, a constant block is cleaner.

Example:

```go
const (
	appName = "Learning Go"
	version = "1.0"
	author  = "Ali"
)
```

This is a common style in Go projects.

---

# Real Example

```go
const (
	appName  = "Learning Go"
	version  = "1.0"
	maxUsers = 100
)
```

Output:

```text
Learning Go
1.0
100
```

---

# Using Constants in Calculations

Constants can be used in mathematical calculations.

Example:

```go
const taxRate = 0.15

price := 100.0

taxAmount := price * taxRate

fmt.Println(taxAmount)
```

Output:

```text
15
```

---

# Common Real-World Constants

## Application Name

```go
const appName = "Learning Go"
```

---

## Application Version

```go
const version = "1.0"
```

---

## Tax Rate

```go
const taxRate = 0.15
```

---

## Days in a Week

```go
const daysInWeek = 7
```

---

## Pi Value

```go
const pi = 3.14159
```

---

# Naming Convention

Constants follow the same naming rules as variables.

Recommended:

```go
const appName = "Learning Go"
const maxUsers = 100
const taxRate = 0.15
```

Use camelCase.

---

Avoid:

```go
const MAX_USERS = 100
```

This style is common in other languages but not standard Go style.

---

# Exported Constants

Lowercase:

```go
const appName = "Learning Go"
```

Only available inside the package.

---

Uppercase:

```go
const AppName = "Learning Go"
```

Available to other packages.

---

# Displaying Constant Types

Use `%T` to display the type.

Example:

```go
const appName = "Learning Go"

fmt.Printf("%T\n", appName)
```

Output:

```text
string
```

---

Example:

```go
const maxUsers = 100

fmt.Printf("%T\n", maxUsers)
```

Output:

```text
int
```

---

# Constants Must Have a Value

Variables can be declared without a value:

```go
var age int
```

Go gives them a zero value.

---

Constants must always have a value.

Incorrect:

```go
const maxUsers
```

Error.

Correct:

```go
const maxUsers = 100
```

---

# Common Mistake

Trying to change a constant.

Incorrect:

```go
const appName = "Learning Go"

appName = "My App"
```

Error:

```text
cannot assign to appName
```

Constants are read-only.

---

# Key Points

- Constants are created using the `const` keyword.
- Constants cannot be changed after creation.
- Constants store fixed values.
- Common constant types:
  - string
  - int
  - float64
  - bool

- Go can automatically determine constant types.
- Constant blocks help organize multiple constants.
- Constants are commonly used for application settings and fixed values.

---

# Summary

In this section, you learned:

- What constants are.
- How constants differ from variables.
- How to create constants.
- Different constant types.
- Constant blocks.
- Using constants in calculations.
- Naming conventions for constants.
- Common real-world uses of constants.

Constants are useful whenever a value should remain fixed and never change during the execution of a program.
