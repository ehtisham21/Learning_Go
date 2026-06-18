# Arithmetic Operators in Go

## Overview

Arithmetic operators are used to perform mathematical calculations in Go.

They allow you to add, subtract, multiply, divide, and find the remainder between numbers.

---

## Arithmetic Operators

| Operator | Description         | Example  |
| -------- | ------------------- | -------- |
| `+`      | Addition            | `10 + 5` |
| `-`      | Subtraction         | `10 - 5` |
| `*`      | Multiplication      | `10 * 5` |
| `/`      | Division            | `10 / 5` |
| `%`      | Modulus (Remainder) | `10 % 3` |

---

## Addition

Adds two values together.

```go
a := 10
b := 5

fmt.Println(a + b)
```

Output:

```text
15
```

### String Concatenation

The `+` operator can also join strings.

```go
firstName := "Ehtisham"
lastName := "Butt"

fmt.Println(firstName + " " + lastName)
```

Output:

```text
Ehtisham Butt
```

---

## Subtraction

Subtracts one value from another.

```go
fmt.Println(10 - 3)
```

Output:

```text
7
```

---

## Multiplication

Multiplies two values.

```go
fmt.Println(10 * 3)
```

Output:

```text
30
```

---

## Division

Divides one value by another.

```go
fmt.Println(10 / 2)
```

Output:

```text
5
```

### Integer Division

When both values are integers, Go returns an integer result.

```go
fmt.Println(10 / 3)
```

Output:

```text
3
```

The decimal part is removed.

---

### Float Division

Convert values to `float64` to get decimal results.

```go
fmt.Println(float64(10) / float64(3))
```

Output:

```text
3.3333333333333335
```

---

## Modulus Operator

The modulus operator `%` returns the remainder after division.

```go
fmt.Println(10 % 3)
```

Output:

```text
1
```

Because:

```text
10 ÷ 3 = 3 remainder 1
```

---

## Even and Odd Numbers

A common use of `%` is checking whether a number is even.

```go
num := 8

fmt.Println(num % 2 == 0)
```

Output:

```text
true
```

Rule:

```go
num % 2 == 0
```

means the number is even.

---

## Operator Precedence

Go follows normal mathematical rules.

Order of execution:

1. `()`
2. `*`, `/`, `%`
3. `+`, `-`

Example:

```go
result := 10 + 5*2
fmt.Println(result)
```

Output:

```text
20
```

Because multiplication happens first.

```text
10 + (5 × 2)
= 20
```

Using parentheses:

```go
result := (10 + 5) * 2
fmt.Println(result)
```

Output:

```text
30
```

---

## Real-World Example

```go
price := 100
tax := 20

total := price + tax

fmt.Println(total)
```

Output:

```text
120
```

---

## Key Points

- Use `+` for addition and joining strings.
- Use `-` for subtraction.
- Use `*` for multiplication.
- Use `/` for division.
- Use `%` to get the remainder.
- `int / int` returns an integer.
- Use `float64()` for decimal division.
- `%` is commonly used to check even and odd numbers.
- Parentheses `()` have the highest priority in calculations.
