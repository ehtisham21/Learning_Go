# 02 - Floats in Go

## What are Floats?

Floats are numbers that contain decimal points.

Examples:

```go
10.5
3.14
99.99
-12.75
0.5
```

Floats are used when a value requires decimal precision.

Common examples:

* Product prices
* Ratings
* Percentages
* Measurements
* Scientific calculations

---

## Float Types in Go

Go provides two floating-point types.

| Type    | Description                  |
| ------- | ---------------------------- |
| float32 | 32-bit floating-point number |
| float64 | 64-bit floating-point number |

Example:

```go
var temperature float32 = 36.5
var price float64 = 10.99
```

---

## Which Float Type Should You Use?

In most applications, use:

```go
float64
```

Reasons:

* Better precision
* Default float type in Go
* Reduces rounding issues

Example:

```go
price := 10.99
```

Go automatically infers the type as:

```go
float64
```

---

## Float Declaration

### Using var

```go
var price float64 = 10.99
```

### Type Inference

```go
var price = 10.99
```

Go automatically determines the type.

### Short Variable Declaration

```go
price := 10.99
```

This is the most common style inside functions.

---

## Zero Value

The default value of a float is:

```go
0
```

Example:

```go
var amount float64

fmt.Println(amount)
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

Example:

```go
a := 10.5
b := 2.5

fmt.Println(a + b)
fmt.Println(a - b)
fmt.Println(a * b)
fmt.Println(a / b)
```

Output:

```go
13
8
26.25
4.2
```

---

## Float Division vs Integer Division

### Integer Division

```go
fmt.Println(10 / 3)
```

Output:

```go
3
```

The decimal portion is discarded.

### Float Division

```go
fmt.Println(10.0 / 3.0)
```

Output:

```go
3.3333333333333335
```

Float division keeps the decimal values.

---

## Checking Variable Types

Use `%T` to see a variable's type.

Example:

```go
price := 10.99

fmt.Printf("%T\n", price)
```

Output:

```go
float64
```

---

## Formatting Float Output

By default, floats may display many decimal places.

Example:

```go
pi := 3.14159265359

fmt.Printf("%.2f\n", pi)
```

Output:

```go
3.14
```

Common formats:

| Format | Description          |
| ------ | -------------------- |
| %f     | Default float output |
| %.2f   | 2 decimal places     |
| %.4f   | 4 decimal places     |

Examples:

```go
fmt.Printf("%f\n", pi)
fmt.Printf("%.2f\n", pi)
fmt.Printf("%.4f\n", pi)
```

---

## Floating-Point Precision

Computers store floating-point numbers in binary format.

Because of this, some decimal values cannot be represented exactly.

Example:

```go
fmt.Println(0.1 + 0.2)
```

Output:

```go
0.30000000000000004
```

This behavior is normal and occurs in many programming languages.

---

## Type Conversion

Go does not allow arithmetic between different numeric types.

Invalid:

```go
age := 25
price := 10.99

fmt.Println(age + price)
```

Error:

```go
mismatched types int and float64
```

Correct:

```go
fmt.Println(float64(age) + price)
```

Output:

```go
35.99
```

---

## Key Points

* Floats store decimal numbers.
* Go provides `float32` and `float64`.
* `float64` is the preferred type in most applications.
* The default inferred float type is `float64`.
* Float division keeps decimal values.
* Integer division removes decimal values.
* Use `%T` to check variable types.
* Use `%.2f` to format decimal places.
* Floating-point precision issues are normal.
* Different numeric types must be converted before being used together.
