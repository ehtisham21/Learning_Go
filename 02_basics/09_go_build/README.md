# 09 - Go Build

## Introduction

`go build` is used to compile Go source code and create an executable application.

Unlike `go run`, which compiles and runs the program immediately, `go build` creates a binary file that can be run later.

This binary can be copied to another machine and executed without needing the Go source code.

---

# What is `go build`?

The `go build` command compiles Go code into a machine-executable program.

Example:

```bash
go build
```

Go will:

1. Read all Go files in the current package.
2. Compile the source code.
3. Create an executable binary.
4. Store the executable in the current directory.

---

# Why Do We Use `go build`?

During development, we usually use:

```bash
go run .
```

because it is fast and convenient.

Before deploying or distributing an application, we use:

```bash
go build
```

to create the final executable.

---

# Example Program

main.go

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go")
}
```

---

# Building the Program

Run:

```bash
go build
```

Before build:

```text
project/
│
└── main.go
```

After build:

```text
project/
│
├── main.go
└── project
```

The new file:

```text
project
```

is the executable application.

---

# Running the Executable

On Linux and macOS:

```bash
./project
```

Output:

```text
Hello Go
```

On Windows:

```cmd
project.exe
```

Output:

```text
Hello Go
```

---

# What Does Go Create?

Go creates a binary executable.

The binary contains:

- Your compiled code.
- Required dependencies.
- Machine instructions for the operating system.

The operating system can execute this file directly.

---

# What Happens Internally?

When we run:

```bash
go build
```

Go performs:

```text
Source Code
     ↓
Compilation
     ↓
Machine Code
     ↓
Executable Binary
```

The result is a runnable application.

---

# Difference Between `go run` and `go build`

## go run

Command:

```bash
go run .
```

Process:

```text
Compile
↓
Run
↓
Delete executable
```

Best for:

- Learning
- Development
- Testing

---

## go build

Command:

```bash
go build
```

Process:

```text
Compile
↓
Create executable
↓
Keep executable
```

Best for:

- Deployment
- Distribution
- Production applications

---

# Creating a Custom Executable Name

By default, Go uses the folder name.

We can choose our own name using:

```bash
go build -o myapp
```

Result:

```text
project/
│
├── main.go
└── myapp
```

Run:

```bash
./myapp
```

---

# Build and Run Example

Build:

```bash
go build -o calculator
```

Run:

```bash
./calculator
```

Output:

```text
Hello Go
```

---

# Real Backend Workflow

Development:

```bash
go run .
```

Testing:

```bash
go test
```

Build:

```bash
go build
```

Deployment:

```bash
./myapp
```

This is a common workflow in Go projects.

---

# Common Commands

Build current package:

```bash
go build
```

Build with custom name:

```bash
go build -o myapp
```

Run executable:

```bash
./myapp
```

---

# Common Errors

## No Main Package

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

Only the `main` package can create executable programs.

---

## Missing main Function

Example:

```text
function main is undeclared in the main package
```

Reason:

```go
func main() {
}
```

is missing.

Every executable Go application must have a `main()` function.

---

# Key Points

- `go build` compiles Go code into an executable application.
- The executable remains in the project directory.
- `go run` creates a temporary executable and removes it after execution.
- `go build -o` allows custom executable names.
- Executables can be run directly from the terminal.
- Go binaries are easy to deploy and distribute.

---

# Summary

In this section, you learned:

- What `go build` is.
- Why developers use `go build`.
- How to create executable binaries.
- How to run compiled applications.
- The difference between `go run` and `go build`.
- How to create custom executable names.
- Why Go binaries are useful for deployment.

`go build` is the command used to turn Go source code into a real application that can be executed and deployed.
