# 03 - Comments in Go

## Introduction

Comments are lines of text in your source code that are ignored by the Go compiler.

They are used to:

- Explain code
- Improve readability
- Document packages, functions, variables, and constants
- Leave notes for other developers
- Temporarily disable code during testing and debugging

Comments do not affect the execution of a program.

---

## Why Comments Are Important

Good comments make code easier to understand and maintain.

---

## Types of Comments in Go

Go supports two types of comments:

### 1. Single-Line Comments

Single-line comments begin with `//`.

Example:

```go
// Print a welcome message
fmt.Println("Hello Go")
```

Everything after `//` on the same line is ignored by the compiler.

---

### 2. Multi-Line Comments

Multi-line comments begin with `/*` and end with `*/`.

Example:

```go
/*
This is a multi-line comment.
It can span multiple lines.
*/
```

These comments are useful for longer explanations.

---

## Example Program

```go
package main

import "fmt"

// This is a single-line comment

func main() {
    fmt.Println("Hello Go")

    /*
       This is a multi-line comment.
       Go ignores everything inside it.
    */

    fmt.Println("Learning comments in Go")
}
```

Output:

```text
Hello Go
Learning comments in Go
```

---

## Commenting Out Code

Comments can be used to temporarily disable code.

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Program started")

    // fmt.Println("This line is disabled")

    fmt.Println("Program ended")
}
```

Output:

```text
Program started
Program ended
```

---

## Documentation Comments

Go uses comments as part of its documentation system.

### Function Documentation

```go
// Add returns the sum of two integers.
func Add(a int, b int) int {
    return a + b
}
```

### Variable Documentation

```go
// MaxUsers defines the maximum number of users.
var MaxUsers = 100
```

### Package Documentation

```go
// Package calculator provides mathematical operations.
package calculator
```

---

## Best Practices

### Write Comments That Explain Why

Good:

```go
// Retry three times because the API may timeout.
```

Bad:

```go
// Loop three times.
```

The code already shows the loop; the comment should explain the reason.

---

### Keep Comments Updated

Bad:

```go
// Maximum users: 50
var MaxUsers = 100
```

Comments should always match the code.

---

### Avoid Obvious Comments

Bad:

```go
// Create variable
count := 10

// Print count
fmt.Println(count)
```

The code is already clear.

---

### Use Meaningful Variable Names

Instead of:

```go
// Store user age
a := 25
```

Use:

```go
userAge := 25
```

Good names often reduce the need for comments.

---

## Common Use Cases

- Explaining complex logic
- Documenting public functions
- Adding TODO notes
- Temporarily disabling code
- Describing package purpose
- Improving code readability

Example:

```go
// TODO: Add database connection handling.
```

---

## Key Takeaways

After completing this topic, you should understand:

- What comments are
- Why comments are important
- Single-line comments (`//`)
- Multi-line comments (`/* */`)
- Commenting out code
- Documentation comments
- Best practices for writing comments
- Common mistakes to avoid

Comments help make your code easier to read, understand, and maintain. Writing clear code with meaningful comments is an important skill for every Go developer.
