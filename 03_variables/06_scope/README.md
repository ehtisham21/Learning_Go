# 06 - Scope

## Introduction

Scope determines where a variable can be accessed and used in a program.

Think of scope as:

```text
Scope = Visibility of a variable
```

Not every variable can be used everywhere.

Go uses scope to keep code organized and prevent conflicts between variables.

---

# Why is Scope Important?

Imagine a program with hundreds of variables.

Without scope:

```text
Every variable would be accessible everywhere.
```

This would make programs difficult to understand and maintain.

Scope helps by controlling where variables can be used.

---

# Types of Scope in Go

In Go, beginners should learn these three scopes:

```text
1. Package Scope (Global Scope)
2. Function Scope (Local Scope)
3. Block Scope
```

---

# 1. Package Scope (Global Scope)

Variables declared outside all functions belong to the package.

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

Because `appName` is declared outside all functions, it can be used by every function in the package.

---

## Example

```go
var appName = "Learning Go"

func showAppName() {
	fmt.Println(appName)
}

func main() {
	fmt.Println(appName)
	showAppName()
}
```

Output:

```text
Learning Go
Learning Go
```

---

# 2. Function Scope (Local Scope)

Variables declared inside a function are local to that function.

Example:

```go
func main() {

	userName := "Ali"

	fmt.Println(userName)
}
```

Output:

```text
Ali
```

The variable `userName` exists only inside `main()`.

---

## Invalid Example

```go
func main() {
	userName := "Ali"
}

func showUser() {
	fmt.Println(userName)
}
```

Error:

```text
undefined: userName
```

Because `userName` belongs only to `main()`.

---

# 3. Block Scope

A block is code inside curly braces:

```go
{
}
```

Example:

```go
if true {

	message := "Hello"

	fmt.Println(message)
}
```

Output:

```text
Hello
```

The variable `message` exists only inside the `if` block.

---

## Invalid Example

```go
if true {

	message := "Hello"
}

fmt.Println(message)
```

Error:

```text
undefined: message
```

Because the variable is outside its scope.

---

# Inner Blocks Can Access Outer Variables

Example:

```go
city := "Lahore"

if true {
	fmt.Println(city)
}
```

Output:

```text
Lahore
```

Inner blocks can use variables created outside them.

---

# Outer Blocks Cannot Access Inner Variables

Example:

```go
if true {

	message := "Hello"
}

fmt.Println(message)
```

Output:

```text
Error
```

Variables created inside a block stay inside that block.

---

# Variable Shadowing

Shadowing happens when a new variable uses the same name as an existing variable.

Example:

```go
name := "Ali"

if true {

	name := "Ahmed"

	fmt.Println(name)
}

fmt.Println(name)
```

Output:

```text
Ahmed
Ali
```

---

## Why?

Inside the block:

```go
name := "Ahmed"
```

creates a completely new variable.

It temporarily hides the outer variable.

---

# Visual Example

```text
Outer Variable

name = Ali

    if block
        name = Ahmed
```

Inside block:

```text
Ahmed
```

Outside block:

```text
Ali
```

---

# Updating an Existing Variable

To update a variable, use:

```go
=
```

Example:

```go
name := "Ali"

if true {
	name = "Ahmed"
}

fmt.Println(name)
```

Output:

```text
Ahmed
```

Notice:

```go
name = "Ahmed"
```

updates the original variable.

---

# Shadowing vs Updating

## Shadowing

```go
name := "Ali"

if true {
	name := "Ahmed"
}
```

Creates a new variable.

---

## Updating

```go
name := "Ali"

if true {
	name = "Ahmed"
}
```

Updates the existing variable.

---

# Local Variable Hides Global Variable

Example:

```go
var appName = "Learning Go"

func main() {

	appName := "My App"

	fmt.Println(appName)
}
```

Output:

```text
My App
```

The local variable hides the global variable inside `main()`.

---

# Real Example

```go
package main

import "fmt"

var companyName = "ABC Company"

func showCompany() {
	fmt.Println(companyName)
}

func main() {

	userName := "Ali"

	fmt.Println(userName)

	showCompany()
}
```

Output:

```text
Ali
ABC Company
```

---

# Common Beginner Mistakes

## Using Local Variable Outside Function

```go
func main() {
	name := "Ali"
}

fmt.Println(name)
```

Error:

```text
undefined: name
```

---

## Using Block Variable Outside Block

```go
if true {
	message := "Hello"
}

fmt.Println(message)
```

Error:

```text
undefined: message
```

---

## Accidentally Shadowing Variables

```go
name := "Ali"

if true {
	name := "Ahmed"
}
```

This creates a new variable.

It does not update the original one.

---

# Scope Summary

| Scope Type | Where It Is Declared | Where It Can Be Used |
|------------|---------------------|----------------------|
| Package Scope | Outside functions | Entire package |
| Function Scope | Inside function | Same function only |
| Block Scope | Inside `{}` block | Same block only |

---

# Key Points

- Scope controls where variables can be used.
- Package scope variables are accessible throughout the package.
- Function scope variables are accessible only inside the function.
- Block scope variables are accessible only inside the block.
- Inner blocks can access outer variables.
- Outer blocks cannot access inner variables.
- Shadowing creates a new variable.
- Updating uses `=` instead of `:=`.

---

# Summary

In this section, you learned:

- What scope is.
- Package scope (global variables).
- Function scope (local variables).
- Block scope.
- Variable shadowing.
- Updating variables correctly.
- Common scope-related mistakes.

Scope is one of the most important concepts in Go because it affects how variables behave in functions, loops, conditions, structs, packages, and real-world applications.