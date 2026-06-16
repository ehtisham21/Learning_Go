# 01 - Go Project Setup

## Overview

This section covers the initial setup required to start learning Go, including Go installation, project initialization, module management, dependency handling, and Git integration.

---

## Installing Go

Verify that Go is installed on your system:

```bash
go version
```

Example output:

```text
go version go1.25.0 linux/amd64
```

---

## Creating a Go Project

Create a new project directory:

```bash
mkdir go-learning
cd go-learning
```

## What is a Go Module?

A Go module is the standard way to manage project dependencies and package imports.

Every Go project should have a module.

Initialize a module:

```bash
go mod init github.com/ehtisham21/go-learning
```

This creates:

```text
go.mod
```

---

## What is go.mod?

The `go.mod` file defines:

- Module name
- Go version
- Project dependencies
- Dependency versions

Example:

```go
module github.com/buttehtisham123/go-learning

go 1.25.0
```

As dependencies are added, they will automatically appear in this file.

Example:

```go
require github.com/gin-gonic/gin v1.10.1
```

### Purpose of go.mod

- Identifies the project
- Tracks dependencies
- Tracks dependency versions
- Enables reproducible builds
- Allows Go to download required packages automatically

---

## What is go.sum?

The `go.sum` file stores cryptographic checksums (hashes) for downloaded dependencies.

Example:

```text
github.com/gin-gonic/gin v1.10.1 h1:xxxxxxxxxxxx
github.com/gin-gonic/gin v1.10.1/go.mod h1:xxxxxxxxxxxx
```

### Purpose of go.sum

- Verifies package integrity
- Prevents corrupted downloads
- Ensures everyone gets the exact same dependency
- Improves security and reproducibility

---

## When is go.sum Created?

Initially:

```bash
go mod init github.com/buttehtisham123/go-learning
```

creates only:

```text
go.mod
```

The `go.sum` file appears automatically when external dependencies are downloaded.

Example:

```bash
go get github.com/gin-gonic/gin
```

Result:

```text
go.mod
go.sum
```

---

## Difference Between go.mod and go.sum

| File   | Purpose                                        |
| ------ | ---------------------------------------------- |
| go.mod | Defines project configuration and dependencies |
| go.sum | Stores checksums for dependency verification   |

Think of them as:

| Ecosystem | Configuration                     | Lock / Verification |
| --------- | --------------------------------- | ------------------- |
| Python    | requirements.txt / pyproject.toml | poetry.lock         |
| Node.js   | package.json                      | package-lock.json   |
| Go        | go.mod                            | go.sum              |

---

## Dependency Management

Download a dependency:

```bash
go get github.com/gin-gonic/gin
```

Remove unused dependencies:

```bash
go mod tidy
```

View all project dependencies:

```bash
go list -m all
```

---

## Useful Go Commands

### Run Program

```bash
go run main.go
```

### Build Executable

```bash
go build
```

### Format Code

```bash
go fmt ./...
```

### Run Tests

```bash
go test ./...
```

### Download Dependencies

```bash
go mod tidy
```

### Check Environment

```bash
go env
```

---

## Key Learnings

- Go does not use virtual environments like Python.
- Go uses modules for dependency management.
- Every project should have a `go.mod` file.
- Dependencies are managed automatically by Go.
- `go.sum` stores dependency verification hashes.
- Both `go.mod` and `go.sum` should be committed to Git.
- `go mod tidy` keeps dependencies clean.
- Go projects are typically organized using modules and packages.

---

## Commands Learned

```bash
go version
go mod init github.com/buttehtisham123/go-learning
go mod tidy
go run main.go
go build
go fmt ./...
go test ./...
git status
```
