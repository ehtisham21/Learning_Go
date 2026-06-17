# 10 - Go Format (go fmt)

## Introduction

`go fmt` is a built-in Go command used to automatically format Go code.

It fixes:

- Indentation
- Spacing
- Blank lines
- Code layout

This helps keep Go code clean, readable, and consistent.

One of the best things about Go is that all Go developers use the same code style.

---

# What is `go fmt`?

The `go fmt` command automatically formats your Go code according to Go's official formatting rules.

Example:

Before formatting:

```go
package main
import "fmt"
func main(){
fmt.Println("Hello Go")
}
```

Run:

```bash
go fmt
```

After formatting:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go")
}
```

Go automatically fixed the formatting.

---

# Why Use `go fmt`?

Without a formatter, every developer may write code differently.

Example:

Developer 1:

```go
if age > 18 {
	fmt.Println("Adult")
}
```

Developer 2:

```go
if(age>18){
fmt.Println("Adult")
}
```

Developer 3:

```go
if age > 18
{
	fmt.Println("Adult")
}
```

All work, but they look different.

With `go fmt`, everyone gets the same style.

This makes code easier to read and maintain.

---

# Basic Usage

Format the current package:

```bash
go fmt
```

This is the command you will use most often.

---

# Format a Specific File

Example:

```bash
go fmt main.go
```

Go formats only that file.

---

# Format the Entire Project

Example:

```bash
go fmt ./...
```

The `./...` means:

```text
Current folder and all subfolders
```

This is useful for large projects.

---

# What Does `go fmt` Fix?

## Indentation

Before:

```go
func main() {
fmt.Println("Hello")
}
```

After:

```go
func main() {
	fmt.Println("Hello")
}
```

---

## Spaces

Before:

```go
age:=25
```

After:

```go
age := 25
```

---

## Blank Lines

Before:

```go
package main
import "fmt"
```

After:

```go
package main

import "fmt"
```

---

# Example

Before:

```go
package main
import "fmt"
func main(){
name:="Ali"
fmt.Println(name)
}
```

Run:

```bash
go fmt
```

After:

```go
package main

import "fmt"

func main() {
	name := "Ali"
	fmt.Println(name)
}
```

---

# What `go fmt` Cannot Fix

`go fmt` only fixes formatting.

It does not fix coding mistakes.

Example:

```go
fmt.Println("Hello"
```

Missing:

```go
)
```

Running:

```bash
go fmt
```

will still show an error because the code itself is invalid.

---

# When Should You Run `go fmt`?

## Before Running Code

```bash
go fmt
go run .
```

---

## Before Committing Code

```bash
go fmt ./...
git add .
git commit -m "Update code"
```

---

## Before Creating a Pull Request

Always format your code before opening a PR.

This is standard practice in Go projects.

---

# VS Code and go fmt

If you have the Go extension installed in VS Code, your code can be formatted automatically when you save the file.

Example:

```text
Ctrl + S
```

↓

VS Code runs:

```bash
go fmt
```

↓

Code becomes properly formatted.

---

# go fmt vs gofmt

You may also see:

```bash
gofmt
```

Example:

```bash
gofmt -w main.go
```

Both are related.

For beginners:

```bash
go fmt
```

is enough.

Most Go developers use `go fmt`.

---

# Typical Go Workflow

Write code:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go")
}
```

Format code:

```bash
go fmt
```

Run code:

```bash
go run .
```

Build application:

```bash
go build
```

---

# Common Commands

Format current package:

```bash
go fmt
```

Format a specific file:

```bash
go fmt main.go
```

Format the entire project:

```bash
go fmt ./...
```

---

# Why Go Developers Like go fmt

Benefits:

- Clean code
- Same style everywhere
- Easier code reviews
- Easier teamwork
- No arguments about formatting

Everyone writes code differently, but after running `go fmt`, the code looks the same.

---

# Key Points

- `go fmt` automatically formats Go code.
- It fixes spacing, indentation, and layout.
- It follows Go's official style guide.
- It does not fix programming errors.
- Most developers run it before running or committing code.
- VS Code can run it automatically on save.
- `go fmt` is one of the most commonly used Go commands.

---

# Summary

In this section, you learned:

- What `go fmt` is.
- Why Go uses automatic formatting.
- How to format a file.
- How to format a package.
- How to format an entire project.
- What `go fmt` can and cannot do.
- How Go developers use formatting in daily work.

`go fmt` helps keep Go code clean, consistent, and easy to read. It is a tool you will use almost every day while writing Go programs.
