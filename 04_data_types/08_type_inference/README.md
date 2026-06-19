# 08 - Type Inference in Go

## What is Type Inference?

Type inference means Go automatically determines the data type of a variable from the value assigned to it.

Example:

```go
var age = 25
```

Go sees:

```text
25
```

and automatically decides:

```text
int
```

---

## Why Use Type Inference?

Without type inference:

```go
var name string = "Ehtisham"
var age int = 25
var price float64 = 99.99
```

With type inference:

```go
var name = "Ehtisham"
var age = 25
var price = 99.99
```

Less code and easier to read.

---

## How Type Inference Works

Go looks at the value on the right side and determines the type automatically.

Example:

```go
var age = 25
```

Go infers:

```text
int
```

---

Example:

```go
var price = 99.99
```

Go infers:

```text
float64
```

---

Example:

```go
var isAdmin = true
```

Go infers:

```text
bool
```

---

Example:

```go
var name = "Ehtisham"
```

Go infers:

```text
string
```

---

## Checking the Type

Use `%T`.

Example:

```go
var age = 25

fmt.Printf("%T\n", age)
```

Output:

```text
int
```

---

## Type Inference Examples

### Integer

```go
var age = 25
```

Type:

```text
int
```

---

### Float

```go
var price = 99.99
```

Type:

```text
float64
```

---

### String

```go
var name = "Go"
```

Type:

```text
string
```

---

### Boolean

```go
var isActive = true
```

Type:

```text
bool
```

---

### Rune

```go
var letter = 'A'
```

Type:

```text
rune (int32)
```

---

## Short Variable Declaration

Go provides a shorter syntax:

```go
age := 25
```

This is equivalent to:

```go
var age = 25
```

Go still performs type inference.

---

Example:

```go
name := "Ehtisham"
price := 99.99
isAdmin := true
```

Go automatically determines all types.

---

## Checking Types of Short Declarations

```go
score := 95

fmt.Printf("%T\n", score)
```

Output:

```text
int
```

---

## Type Inference with Byte Slices

Example:

```go
data := []byte("Hello")
```

Go infers:

```text
[]byte
```

Check type:

```go
fmt.Printf("%T\n", data)
```

Output:

```text
[]uint8
```

Remember:

```text
byte = uint8
```

---

## Explicit Type vs Type Inference

### Explicit Type

```go
var age int64 = 25
```

You tell Go exactly which type to use.

---

### Type Inference

```go
age := 25
```

Go chooses:

```text
int
```

automatically.

---

## When to Use Explicit Types

Use explicit types when you need a specific type.

Example:

```go
var userID int64 = 1001
```

Without explicit type:

```go
userID := 1001
```

Go would infer:

```text
int
```

instead of:

```text
int64
```

---

## Benefits of Type Inference

* Less typing
* Cleaner code
* Easier to read
* Common Go style
* Reduces repetition

---

## Does Type Inference Make Go Dynamic?

No.

Go is still a statically typed language.

Example:

```go
age := 25
```

Go infers:

```text
int
```

After that:

```go
age = "hello"
```

This causes a compile error.

Because:

```text
age is already int
```

---

## Key Points

* Type inference means Go automatically determines the type.
* `var age = 25` → `int`
* `var price = 99.99` → `float64`
* `var name = "Go"` → `string`
* `var isActive = true` → `bool`
* `var letter = 'A'` → `rune`
* `:=` is the most common way to use type inference inside functions.
* Go remains statically typed.
* Use explicit types when a specific type is required.

## Easy Summary

```go
age := 25
```

Go reads:

```text
25
```

and automatically decides:

```text
int
```

This automatic type detection is called **Type Inference**.

### Memory Trick

```text
var age = 25
        ↓
Go decides the type

age := 25
        ↓
Go decides the type
```

You provide the value, and Go figures out the type.
