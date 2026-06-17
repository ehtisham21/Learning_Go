# 06 - Print Output

## Introduction

Printing output is one of the most basic and important tasks in programming.

In Go, we use the `fmt` package to display information on the screen.

Printing is useful for:

* Displaying messages to users.
* Showing program results.
* Debugging code.
* Checking variable values.
* Understanding program flow.

To print output, we first import the `fmt` package:

```go
package main

import "fmt"
```

---

# fmt.Print()

`Print()` displays output without adding a new line at the end.

Example:

```go
fmt.Print("Hello ")
fmt.Print("Go")
```

Output:

```text
Hello Go
```

Notice that both outputs appear on the same line.

---

# fmt.Println()

`Println()` displays output and automatically moves to the next line.

Example:

```go
fmt.Println("Hello")
fmt.Println("Go")
```

Output:

```text
Hello
Go
```

---

## Printing Multiple Values

```go
fmt.Println("Name:", "Ali")
fmt.Println("Age:", 20)
```

Output:

```text
Name: Ali
Age: 20
```

Go automatically adds spaces between values.

---

# Variables and Output

Example:

```go
var name string = "Ali"
var age int = 20

fmt.Println(name)
fmt.Println(age)
```

Output:

```text
Ali
20
```

---

# fmt.Printf()

`Printf()` is used when we want more control over how values are displayed.

Syntax:

```go
fmt.Printf("format", values...)
```

Example:

```go
var name string = "Ali"

fmt.Printf("Name: %s", name)
```

Output:

```text
Name: Ali
```

---

# Formatting Verbs

Formatting verbs start with `%`.

They act as placeholders where values will be inserted.

---

## %s - String

Used for text values.

Example:

```go
var name string = "Ali"

fmt.Printf("Name: %s", name)
```

Output:

```text
Name: Ali
```

---

## %d - Integer

Used for integer values.

Example:

```go
var age int = 20

fmt.Printf("Age: %d", age)
```

Output:

```text
Age: 20
```

---

## %f - Float

Used for decimal numbers.

Example:

```go
var marks float64 = 85.5

fmt.Printf("Marks: %f", marks)
```

Output:

```text
Marks: 85.500000
```

---

## Limiting Decimal Places

Example:

```go
fmt.Printf("%.2f", marks)
```

Output:

```text
85.50
```

`.2` means show only 2 digits after the decimal point.

---

## %t - Boolean

Used for true or false values.

Example:

```go
var passed bool = true

fmt.Printf("Passed: %t", passed)
```

Output:

```text
Passed: true
```

---

## %v - Default Value

`%v` prints the value in its default format.

Example:

```go
fmt.Printf("%v", name)
fmt.Printf("%v", age)
fmt.Printf("%v", marks)
```

Output:

```text
Ali
20
85.5
```

This is useful when you do not want to remember specific formatting verbs.

---

## %T - Type

`%T` displays the data type of a variable.

Example:

```go
fmt.Printf("%T", age)
```

Output:

```text
int
```

Example:

```go
fmt.Printf("%T\n", name)
fmt.Printf("%T\n", age)
fmt.Printf("%T\n", passed)
```

Output:

```text
string
int
bool
```

---

# New Line Character (\n)

`\n` moves output to the next line.

Example:

```go
fmt.Printf("Line 1\nLine 2\nLine 3")
```

Output:

```text
Line 1
Line 2
Line 3
```

---

# Tab Character (\t)

`\t` creates spacing similar to pressing the Tab key.

Example:

```go
fmt.Printf("Name\tAge\n")
fmt.Printf("Ali\t20\n")
```

Output:

```text
Name    Age
Ali     20
```

---

# Percent Sign (%%)

To print a percent sign, use `%%`.

Example:

```go
fmt.Printf("Progress: 80%%")
```

Output:

```text
Progress: 80%
```

---

# fmt.Sprintf()

`Sprintf()` formats text and returns it as a string.

It does not print automatically.

Example:

```go
var name string = "Ali"

message := fmt.Sprintf(
	"Welcome %s to Go",
	name,
)

fmt.Println(message)
```

Output:

```text
Welcome Ali to Go
```

---

## Why Use Sprintf?

Useful when you want to build a string first and use it later.

Example:

```go
var age int = 20

message := fmt.Sprintf(
	"Age is %d",
	age,
)
```

Now `message` contains:

```text
Age is 20
```

---

# fmt.Sprintln()

`Sprintln()` is similar to `Println()`, but instead of printing, it returns a string.

Example:

```go
output := fmt.Sprintln(
	"Welcome",
	"Ali",
	"to Go",
)

fmt.Print(output)
```

Output:

```text
Welcome Ali to Go
```

---

## Difference Between Sprintf and Sprintln

### Sprintf

```go
output := fmt.Sprintf(
	"Welcome %s to Go",
	name,
)
```

Uses formatting verbs.

---

### Sprintln

```go
output := fmt.Sprintln(
	"Welcome",
	name,
	"to Go",
)
```

Automatically adds spaces between values.

---

# Complete Example

```go
package main

import "fmt"

func main() {

	var name string = "Ali"
	var age int = 20
	var marks float64 = 85.5
	var passed bool = true

	fmt.Println("Name:", name)

	fmt.Printf("Age: %d\n", age)

	fmt.Printf("Marks: %.2f\n", marks)

	fmt.Printf("Passed: %t\n", passed)

	fmt.Printf("Type of marks: %T\n", marks)

	message := fmt.Sprintf(
		"Student %s is %d years old",
		name,
		age,
	)

	fmt.Println(message)
}
```

---

# Most Common Functions

| Function       | Purpose                   |
| -------------- | ------------------------- |
| fmt.Print()    | Print without newline     |
| fmt.Println()  | Print with newline        |
| fmt.Printf()   | Print formatted output    |
| fmt.Sprintf()  | Return formatted string   |
| fmt.Sprintln() | Return string with spaces |

---

# Most Common Formatting Verbs

| Verb | Meaning       |
| ---- | ------------- |
| %s   | String        |
| %d   | Integer       |
| %f   | Float         |
| %t   | Boolean       |
| %v   | Default value |
| %T   | Type          |
| %%   | Percent sign  |

---

# Summary

In this section, you learned:

* How to print output using the `fmt` package.
* The difference between `Print()`, `Println()`, and `Printf()`.
* How formatting verbs work.
* How to print strings, integers, floats, and booleans.
* How to display variable types using `%T`.
* How to use new lines and tabs.
* How to create formatted strings using `Sprintf()`.
* How to create strings with spaces using `Sprintln()`.

Printing output is one of the most frequently used skills in Go and is essential for debugging, testing, and displaying information.
