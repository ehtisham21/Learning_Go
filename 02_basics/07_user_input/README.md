# 07 - User Input

## Introduction

User input means taking values from the user while the program is running.

Until now, most values were written directly in code:

```go
var name string = "Ali"
var age int = 20
```

But real programs usually ask the user for data.

Example:

```text
Enter your name: Ali
Hello Ali
```

In Go, we can take basic user input using the `fmt` package.

```go
import "fmt"
```

---

# Why Do We Need User Input?

User input makes a program interactive.

Without input, the program always uses the same values.

With input, the program can behave differently depending on what the user enters.

Example:

```text
Enter first number: 10
Enter second number: 20
Sum: 30
```

---

# Basic Input Flow

The basic flow is:

```text
Declare variable
Ask user for input
Read input using Scan
Store input in variable
Use the variable
```

Example:

```go
var name string

fmt.Print("Enter your name: ")
fmt.Scan(&name)

fmt.Println("Hello", name)
```

---

# fmt.Scan()

`fmt.Scan()` is the simplest way to take input in Go.

Example:

```go
var name string

fmt.Print("Enter your name: ")
fmt.Scan(&name)

fmt.Println("Hello", name)
```

Input:

```text
Ali
```

Output:

```text
Hello Ali
```

---

# Why Do We Use `&`?

In input functions, we usually write:

```go
fmt.Scan(&name)
```

The `&` means we are giving Go the location of the variable.

So Go can store the input inside that variable.

For now, remember this simple rule:

```text
When using Scan, use & before the variable name.
```

Correct:

```go
fmt.Scan(&name)
```

Incorrect:

```go
fmt.Scan(name)
```

---

# String Input

A string stores text.

Example:

```go
var name string

fmt.Print("Enter your name: ")
fmt.Scan(&name)

fmt.Println("Name:", name)
```

Input:

```text
Ali
```

Output:

```text
Name: Ali
```

Important: `fmt.Scan()` reads only one word.

If the user enters:

```text
Muhammad Ali
```

Only this is stored:

```text
Muhammad
```

---

# Integer Input

An integer stores whole numbers.

Example:

```go
var age int

fmt.Print("Enter your age: ")
fmt.Scan(&age)

fmt.Println("Age:", age)
```

Input:

```text
20
```

Output:

```text
Age: 20
```

---

# Float Input

A float stores decimal numbers.

Example:

```go
var marks float64

fmt.Print("Enter your marks: ")
fmt.Scan(&marks)

fmt.Println("Marks:", marks)
```

Input:

```text
85.5
```

Output:

```text
Marks: 85.5
```

---

# Boolean Input

A boolean stores `true` or `false`.

Example:

```go
var passed bool

fmt.Print("Did you pass? Enter true or false: ")
fmt.Scan(&passed)

fmt.Println("Passed:", passed)
```

Input:

```text
true
```

Output:

```text
Passed: true
```

---

# Multiple Inputs

We can take multiple inputs at once.

Example:

```go
var city string
var country string

fmt.Print("Enter city and country: ")
fmt.Scan(&city, &country)

fmt.Println("City:", city)
fmt.Println("Country:", country)
```

Input:

```text
Lahore Pakistan
```

Output:

```text
City: Lahore
Country: Pakistan
```

`Scan()` reads values separated by spaces.

---

# Taking Two Numbers and Adding Them

Example:

```go
var num1 int
var num2 int

fmt.Print("Enter two numbers: ")
fmt.Scan(&num1, &num2)

sum := num1 + num2

fmt.Println("Sum:", sum)
```

Input:

```text
10 20
```

Output:

```text
Sum: 30
```

---

# fmt.Scanln()

`fmt.Scanln()` is similar to `fmt.Scan()`.

It reads input until the user presses Enter.

Example:

```go
var language string

fmt.Print("Enter your favorite language: ")
fmt.Scanln(&language)

fmt.Println("Favorite Language:", language)
```

Input:

```text
Go
```

Output:

```text
Favorite Language: Go
```

For simple examples, `Scan()` and `Scanln()` often look similar.

---

# Difference Between Scan and Scanln

## Scan

```go
fmt.Scan(&name, &age)
```

Can read values separated by spaces or new lines.

Example input:

```text
Ali 20
```

or:

```text
Ali
20
```

---

## Scanln

```go
fmt.Scanln(&name, &age)
```

Reads values from the current line only.

Example input:

```text
Ali 20
```

This works.

But if the values are not on the same line, it may not behave as expected.

---

# fmt.Scanf()

`fmt.Scanf()` reads input according to a specific format.

Example:

```go
var name string
var age int

fmt.Print("Enter name and age: ")
fmt.Scanf("%s %d", &name, &age)

fmt.Println("Name:", name)
fmt.Println("Age:", age)
```

Input:

```text
Ali 20
```

Output:

```text
Name: Ali
Age: 20
```

Here:

```text
%s means string
%d means integer
```

So this format:

```go
"%s %d"
```

means:

```text
First read a string, then read an integer.
```

---

# Difference Between Scan, Scanln, and Scanf

| Function       | Meaning                         | Best For               |
| -------------- | ------------------------------- | ---------------------- |
| `fmt.Scan()`   | Reads input separated by spaces | Simple beginner input  |
| `fmt.Scanln()` | Reads input until Enter         | Values on one line     |
| `fmt.Scanf()`  | Reads input using a format      | Specific input pattern |

For beginners, `fmt.Scan()` is usually enough.

---

# Important Limitation

`fmt.Scan()`, `fmt.Scanln()`, and `fmt.Scanf()` are not good for reading full names or full sentences with spaces.

Example:

```text
Muhammad Ali
```

With:

```go
fmt.Scan(&name)
```

Only this is stored:

```text
Muhammad
```

Later, for full-line input, we can use:

```go
bufio.NewReader(os.Stdin)
```

But for now, `fmt.Scan()` is enough for learning basics.

---

# Complete Example

```go
package main

import "fmt"

// This program shows different ways to take user input in Go.
func main() {

	var name string

	fmt.Print("Enter your first name: ")
	fmt.Scan(&name)

	fmt.Println("Hello", name)

	var age int

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Your age is:", age)

	var marks float64

	fmt.Print("Enter your marks: ")
	fmt.Scan(&marks)

	fmt.Println("Your marks are:", marks)

	var passed bool

	fmt.Print("Did you pass? Enter true or false: ")
	fmt.Scan(&passed)

	fmt.Println("Passed:", passed)

	var city string
	var country string

	fmt.Print("Enter city and country: ")
	fmt.Scan(&city, &country)

	fmt.Println("City:", city)
	fmt.Println("Country:", country)

	var num1 int
	var num2 int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&num1, &num2)

	sum := num1 + num2

	fmt.Println("First Number:", num1)
	fmt.Println("Second Number:", num2)
	fmt.Println("Sum:", sum)

	var favoriteLanguage string

	fmt.Print("Enter your favorite programming language: ")
	fmt.Scanln(&favoriteLanguage)

	fmt.Println("Favorite Language:", favoriteLanguage)

	var studentName string
	var studentAge int

	fmt.Print("Enter student name and age: ")
	fmt.Scanf("%s %d", &studentName, &studentAge)

	fmt.Println("Student Name:", studentName)
	fmt.Println("Student Age:", studentAge)

	fmt.Println()
	fmt.Println("===== Final Summary =====")

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Marks: %.2f\n", marks)
	fmt.Printf("Passed: %t\n", passed)
	fmt.Printf("City: %s\n", city)
	fmt.Printf("Country: %s\n", country)
	fmt.Printf("Sum of numbers: %d\n", sum)
}
```

---

# How to Run

From this folder:

```bash
cd 02_basics/07_user_input
go run main.go
```

---

# Example Program Run

```text
Enter your first name: Ali
Hello Ali

Enter your age: 20
Your age is: 20

Enter your marks: 85.5
Your marks are: 85.5

Did you pass? Enter true or false: true
Passed: true

Enter city and country: Lahore Pakistan
City: Lahore
Country: Pakistan

Enter two numbers: 10 20
First Number: 10
Second Number: 20
Sum: 30

Enter your favorite programming language: Go
Favorite Language: Go

Enter student name and age: Ahmed 22
Student Name: Ahmed
Student Age: 22
```

---

# Common Mistakes

## Forgetting `&`

Incorrect:

```go
fmt.Scan(name)
```

Correct:

```go
fmt.Scan(&name)
```

---

## Entering the Wrong Type

If variable is `int`:

```go
var age int
fmt.Scan(&age)
```

The user should enter:

```text
20
```

Not:

```text
twenty
```

---

## Expecting Full Sentence Input

This will not read a full sentence:

```go
var sentence string
fmt.Scan(&sentence)
```

Input:

```text
I love Go
```

Only this is stored:

```text
I
```

---

# Key Points

* User input makes a program interactive.
* `fmt.Scan()` is the simplest input function.
* Use `&` before variable names when taking input.
* `Scan()` reads input separated by spaces.
* `Scanln()` reads input until Enter.
* `Scanf()` reads input based on a format.
* Use `%s` for string input.
* Use `%d` for integer input.
* Use `float64` for decimal values.
* Use `bool` for true or false values.
* For now, use `fmt.Scan()` for beginner practice.

---

# Summary

In this section, you learned:

* What user input is.
* Why user input is useful.
* How to take string input.
* How to take integer input.
* How to take float input.
* How to take boolean input.
* How to take multiple inputs.
* The difference between `Scan()`, `Scanln()`, and `Scanf()`.
* Why `&` is required in input functions.
* Common mistakes beginners make with input.

This topic is important because it allows your Go programs to receive data from users instead of using only hardcoded values.
