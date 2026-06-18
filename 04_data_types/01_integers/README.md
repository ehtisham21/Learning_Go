# 01 - Integers in Go

## What are Integers?

Integers are whole numbers that do not contain decimal points.

Examples:

```go
10
25
100
-5
0
```

Integers are commonly used for:

* Age
* Quantity
* Scores
* Counters
* IDs
* Number of users

---

## Integer Types in Go

Go provides several integer types.

| Type  | Description                             |
| ----- | --------------------------------------- |
| int   | Most commonly used integer type         |
| int8  | Small integer                           |
| int16 | Medium integer                          |
| int32 | Larger integer                          |
| int64 | Very large integer                      |
| uint  | Unsigned integer (positive values only) |

For most programs, use:

```go
var age int = 25
```

or

```go
age := 25
```

---

## Integer Declaration

### Using var

```go
var age int = 25
```

### Type Inference

```go
var age = 25
```

Go automatically determines that the type is `int`.

### Short Variable Declaration

```go
age := 25
```

This is the most common style inside functions.

---

## Signed Integers

Signed integers can store:

* Positive numbers
* Negative numbers
* Zero

Example:

```go
var temperature int = -5
var score int = 100
```

---

## Unsigned Integers

Unsigned integers can only store positive values and zero.

Example:

```go
var users uint = 50
```

Valid values:

```go
0
1
100
500
```

Invalid value:

```go
var users uint = -10
```

This will produce a compile-time error.

---

## Zero Value

If an integer variable is declared without a value, Go automatically assigns the zero value.

```go
var marks int
```

Output:

```go
0
```

---

## Arithmetic Operations

### Addition

```go
a + b
```

### Subtraction

```go
a - b
```

### Multiplication

```go
a * b
```

### Division

```go
a / b
```

### Modulus (Remainder)

```go
a % b
```

Example:

```go
a := 10
b := 3

fmt.Println(a + b) // 13
fmt.Println(a - b) // 7
fmt.Println(a * b) // 30
fmt.Println(a / b) // 3
fmt.Println(a % b) // 1
```

---

## Integer Division

When both operands are integers, Go returns an integer result.

Example:

```go
fmt.Println(10 / 3)
```

Output:

```go
3
```

The decimal portion is discarded.

---

## Increment Operator

Increase a value by one.

```go
count := 5
count++
```

Output:

```go
6
```

---

## Decrement Operator

Decrease a value by one.

```go
count--
```

Output:

```go
4
```

---

## Checking Variable Types

Use `%T` to see the type of a variable.

```go
age := 25

fmt.Printf("%T\n", age)
```

Output:

```go
int
```

---

## Key Points

* Integers store whole numbers.
* `int` is the most commonly used integer type.
* Signed integers can be positive or negative.
* Unsigned integers (`uint`) cannot store negative values.
* The zero value of an integer is `0`.
* Integer division removes the decimal portion.
* `%` returns the remainder.
* `++` increases a value by 1.
* `--` decreases a value by 1.
* `:=` is the preferred way to declare variables inside functions.
* `%T` displays the variable type.
