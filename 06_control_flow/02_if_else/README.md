# 02_if_else

## What is if else?

The `if else` statement is used when a program needs to choose between two possible actions.

- If the condition is `true`, the `if` block runs.
- If the condition is `false`, the `else` block runs.

Only one block executes.

---

## Basic Syntax

```go
if condition {
    // runs if condition is true
} else {
    // runs if condition is false
}
```

Example:

```go
age := 15

if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

Output:

```text
Minor
```

---

## How it Works

```go
score := 90

if score >= 50 {
    fmt.Println("Pass")
} else {
    fmt.Println("Fail")
}
```

Since `90 >= 50` is true, Go executes:

```go
fmt.Println("Pass")
```

and skips the `else` block.

---

## Multiple Statements

You can place multiple statements inside an `if` or `else` block.

```go
if age >= 18 {
    fmt.Println("Can Vote")
    fmt.Println("Can Drive")
} else {
    fmt.Println("Minor")
    fmt.Println("Cannot Vote")
}
```

---

## else if

Use `else if` when there are multiple conditions.

Syntax:

```go
if condition1 {
} else if condition2 {
} else {
}
```

Example:

```go
marks := 85

if marks >= 90 {
    fmt.Println("Grade A+")
} else if marks >= 80 {
    fmt.Println("Grade A")
} else if marks >= 70 {
    fmt.Println("Grade B")
} else {
    fmt.Println("Fail")
}
```

Output:

```text
Grade A
```

---

## Condition Order Matters

Go checks conditions from top to bottom.

Correct:

```go
if marks >= 90 {
    fmt.Println("A+")
} else if marks >= 80 {
    fmt.Println("A")
} else if marks >= 70 {
    fmt.Println("B")
}
```

The first true condition wins.

---

## Using Logical Operators

### AND (&&)

Both conditions must be true.

```go
if age >= 18 && hasID {
    fmt.Println("Entry Allowed")
}
```

### OR (||)

At least one condition must be true.

```go
if isAdmin || isManager {
    fmt.Println("Access Granted")
}
```

### NOT (!)

Reverses a boolean value.

```go
if !isBlocked {
    fmt.Println("User Active")
}
```

---

## Variable Declaration Inside if

Go allows creating variables directly inside an `if`.

```go
if age := 20; age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

The variable exists only within the `if` and `else` blocks.

---

## Real Backend Examples

### Authentication

```go
if isLoggedIn {
    fmt.Println("Show Dashboard")
} else {
    fmt.Println("Redirect To Login")
}
```

### Request Validation

```go
if name != "" {
    fmt.Println("Valid Request")
} else {
    fmt.Println("Name Required")
}
```

### Package Version Validation

```go
if version != "" {
    fmt.Println("Process Package")
} else {
    fmt.Println("Version Missing")
}
```

### Role-Based Access

```go
if role == "admin" {
    fmt.Println("Full Access")
} else if role == "editor" {
    fmt.Println("Limited Access")
} else {
    fmt.Println("Read Only Access")
}
```

---

## Common Mistakes

### Using Multiple if Statements

❌ Wrong

```go
if marks >= 90 {
    fmt.Println("A+")
}

if marks >= 80 {
    fmt.Println("A")
}
```

Output:

```text
A+
A
```

Both conditions run.

✅ Correct

```go
if marks >= 90 {
    fmt.Println("A+")
} else if marks >= 80 {
    fmt.Println("A")
}
```

Only one block executes.

---

### Wrong Condition Order

❌ Wrong

```go
if marks >= 70 {
    fmt.Println("B")
} else if marks >= 90 {
    fmt.Println("A+")
}
```

Output:

```text
B
```

The first condition already matched.

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

## Key Points

- `if else` provides two execution paths.
- Only one block executes.
- `else if` is used for multiple conditions.
- Conditions are checked from top to bottom.
- The first true condition wins.
- Order of conditions is important.
- Variables can be declared inside an `if`.
- Variables declared inside an `if` only exist within the `if` and `else` blocks.
- Curly braces `{}` are mandatory.

---
