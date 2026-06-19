# 01_if

## What is `if`?

The `if` statement is used to make decisions in Go.

It allows a block of code to run only when a condition is `true`.

---

## Basic Syntax

```go
if condition {
    // code to execute
}
```

Example:

```go
age := 20

if age >= 18 {
    fmt.Println("Adult")
}
```

Output:

```text
Adult
```

---

## Comparison Operators

Comparison operators are commonly used in `if` conditions.

| Operator | Meaning                  |
| -------- | ------------------------ |
| `==`     | Equal to                 |
| `!=`     | Not equal to             |
| `>`      | Greater than             |
| `<`      | Less than                |
| `>=`     | Greater than or equal to |
| `<=`     | Less than or equal to    |

Example:

```go
score := 90

if score >= 80 {
    fmt.Println("Passed")
}
```

---

## Using Boolean Variables

A boolean variable can be used directly inside an `if`.

```go
isAdmin := true

if isAdmin {
    fmt.Println("Access Granted")
}
```

Output:

```text
Access Granted
```

---

## Logical Operators

### AND (`&&`)

Both conditions must be true.

```go
age := 25
hasID := true

if age >= 18 && hasID {
    fmt.Println("Entry Allowed")
}
```

### OR (`||`)

At least one condition must be true.

```go
isAdmin := false
isManager := true

if isAdmin || isManager {
    fmt.Println("Access Granted")
}
```

### NOT (`!`)

Reverses a boolean value.

```go
isBlocked := false

if !isBlocked {
    fmt.Println("User Active")
}
```

---

## Multiple Statements Inside an if Block

All statements inside the curly braces run when the condition is true.

```go
if age >= 18 {
    fmt.Println("Can Vote")
    fmt.Println("Can Drive")
    fmt.Println("Can Apply for Passport")
}
```

---

## Variable Declaration Inside if

Go allows variable creation inside an `if`.

```go
if marks := 85; marks >= 80 {
    fmt.Println("Excellent Score")
}
```

The variable exists only inside that `if` block.

---

## Scope Example

```go
if age := 20; age >= 18 {
    fmt.Println(age)
}

fmt.Println(age)
```

This produces an error because `age` only exists inside the `if` statement.

---

## Real Backend Examples

### Validate Request Data

```go
name := "Ehtisham"

if name != "" {
    fmt.Println("Valid Request")
}
```

### Process Package Version

```go
version := "1.2.0"

if version != "" {
    fmt.Println("Processing Package")
}
```

### Check Authentication

```go
isLoggedIn := true

if isLoggedIn {
    fmt.Println("Show Dashboard")
}
```

---

## Common Mistakes

### Using `=` Instead of `==`

❌ Wrong

```go
if age = 18 {
}
```

✅ Correct

```go
if age == 18 {
}
```

---

### Forgetting Curly Braces

❌ Wrong

```go
if age >= 18
    fmt.Println("Adult")
```

✅ Correct

```go
if age >= 18 {
    fmt.Println("Adult")
}
```

---

### Using Non-Boolean Values

❌ Wrong

```go
if age {
}
```

✅ Correct

```go
if age > 18 {
}
```

---

## Key Points

- `if` is used for decision making.
- Conditions must return `true` or `false`.
- Curly braces `{}` are mandatory.
- Comparison operators are commonly used in conditions.
- Logical operators (`&&`, `||`, `!`) combine conditions.
- Variables can be declared inside an `if`.
- Variables declared inside an `if` only exist within that block.
- Go does not allow non-boolean conditions like Python.

---
