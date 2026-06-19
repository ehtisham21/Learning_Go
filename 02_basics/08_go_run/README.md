# 08 - Go Run

## Introduction

`go run` is one of the most commonly used Go commands during development.

It allows us to compile and run a Go program in a single command.

Instead of manually compiling the code and then running the executable, Go does everything automatically.

---

# What is `go run`?

The `go run` command:

1. Compiles the Go source code.
2. Creates a temporary executable file.
3. Runs the executable.
4. Deletes the temporary executable after execution.

Example:

```bash
go run main.go
```

---

# Why Use `go run`?

During development, we frequently make changes to our code.

Using `go run` allows us to quickly test our program without creating an executable file every time.

Example:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go")
}
```

Run:

```bash
go run main.go
```

Output:

```text
Hello Go
```

---

# Basic Syntax

## Run a Single File

```bash
go run main.go
```

Example:

```go
package main

import "fmt"

func main() {
	fmt.Println("Learning Go")
}
```

Output:

```text
Learning Go
```

---

# Running the Current Package

Instead of specifying individual files, we can run the entire package.

```bash
go run .
```

The dot (`.`) means:

```text
Current Directory
```

Go will:

* Find all `.go` files in the current package.
* Compile them together.
* Run the application.

This is the most common approach used by Go developers.

---

# Example Project

```text
project/
│
├── main.go
└── helper.go
```

main.go

```go
package main

func main() {
	SayHello()
}
```

helper.go

```go
package main

import "fmt"

func SayHello() {
	fmt.Println("Hello from helper")
}
```

Run:

```bash
go run .
```

Output:

```text
Hello from helper
```

---

# Requirements for `go run`

To run a Go application, the package must contain:

```go
package main
```

and

```go
func main() {
}
```

Without these, Go cannot create an executable program.

---

# Common Error

Example:

```text
package xxx is not a main package
```

Reason:

```go
package utils
```

instead of:

```go
package main
```

To use `go run`, the package must be `main`.

---

# Another Common Error

Example:

```text
function main is undeclared in the main package
```

Reason:

Missing:

```go
func main() {
}
```

Every executable Go application needs a `main()` function.

---

# What Happens Internally?

When we run:

```bash
go run .
```

Go performs the following steps:

```text
Source Code
     ↓
Compilation
     ↓
Temporary Executable
     ↓
Execution
     ↓
Delete Temporary Executable
```

The executable is not kept after the program finishes.

---

# Difference Between `go run` and `go build`

## go run

```bash
go run .
```

Behavior:

```text
Compile
↓
Run
↓
Delete executable
```

Best for:

* Learning
* Testing
* Development

---

## go build

```bash
go build
```

Behavior:

```text
Compile
↓
Create executable
↓
Keep executable
```

Best for:

* Deployment
* Distribution
* Production builds

---

# Recommended Workflow

While learning Go:

```bash
go run .
```

Use this command after every code change to quickly test your program.

When preparing a real application:

```bash
go build
```

Create the final executable.

---

# Key Points

* `go run` compiles and runs a Go program.
* A temporary executable is created and automatically removed.
* The package must be `main`.
* A `main()` function must exist.
* `go run .` runs the current package.
* `go run` is mainly used during development.
* `go build` is used to create a permanent executable.

---

# Summary

In this section, you learned:

* What `go run` is.
* Why developers use `go run`.
* How to run a single file.
* How to run the current package.
* The requirements for executable Go programs.
* Common `go run` errors.
* The difference between `go run` and `go build`.

`go run` is the primary command used while developing Go applications because it allows quick testing without generating permanent executable files.
