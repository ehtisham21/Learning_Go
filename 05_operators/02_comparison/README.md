# Comparison Operators in Go

## Overview

Comparison operators are used to compare two values.

The result of every comparison is a boolean value:

```go
true
```

or

```go
false
```

Comparison operators are commonly used in:

* `if` statements
* `switch` statements
* loops
* validations
* decision making

---

## Comparison Operators

| Operator | Description              | Example    |
| -------- | ------------------------ | ---------- |
| `==`     | Equal To                 | `10 == 10` |
| `!=`     | Not Equal To             | `10 != 5`  |
| `>`      | Greater Than             | `10 > 5`   |
| `<`      | Less Than                | `5 < 10`   |
| `>=`     | Greater Than or Equal To | `18 >= 18` |
| `<=`     | Less Than or Equal To    | `18 <= 18` |

---

## Equal To (`==`)

Checks whether two values are equal.

```go
age := 25

fmt.Println(age == 25)
```

Output:

```text
true
```

---

## Not Equal To (`!=`)

Checks whether two values are different.

```go
age := 25

fmt.Println(age != 18)
```

Output:

```text
true
```

---

## Greater Than (`>`)

Checks if the left value is greater than the right value.

```go
fmt.Println(20 > 10)
```

Output:

```text
true
```

---

## Less Than (`<`)

Checks if the left value is smaller than the right value.

```go
fmt.Println(5 < 10)
```

Output:

```text
true
```

---

## Greater Than or Equal To (`>=`)

Checks if a value is greater than or equal to another value.

```go
age := 18

fmt.Println(age >= 18)
```

Output:

```text
true
```

---

## Less Than or Equal To (`<=`)

Checks if a value is less than or equal to another value.

```go
age := 18

fmt.Println(age <= 18)
```

Output:

```text
true
```

---

## Comparing Strings

Strings can also be compared.

```go
name := "Ehtisham"

fmt.Println(name == "Ehtisham")
```

Output:

```text
true
```

### Case Sensitive

```go
fmt.Println("Go" == "go")
```

Output:

```text
false
```

Because uppercase and lowercase letters are different.

---

## Comparing Variables

```go
a := 10
b := 20

fmt.Println(a == b)
fmt.Println(a != b)
fmt.Println(a > b)
fmt.Println(a < b)
```

Output:

```text
false
true
false
true
```

---

## Storing Comparison Results

Since comparisons return boolean values, we can store them in variables.

```go
age := 25

isAdult := age >= 18

fmt.Println(isAdult)
```

Output:

```text
true
```

---

## Real-World Examples

### Age Validation

```go
age := 20

fmt.Println(age >= 18)
```

Output:

```text
true
```

---

### Password Check

```go
password := "admin123"

fmt.Println(password == "admin123")
```

Output:

```text
true
```

---

### Product Availability

```go
stock := 15

fmt.Println(stock > 0)
```

Output:

```text
true
```

---

## Assignment vs Comparison

A common beginner mistake is confusing:

### Assignment

```go
x = 20
```

This assigns a value.

### Comparison

```go
x == 20
```

This compares values and returns `true` or `false`.

---

## Invalid Comparison

Go does not allow comparing different types.

```go
num := 10

fmt.Println(num == "10")
```

This causes a compilation error because `int` and `string` are different types.

---

## Key Points

* Comparison operators always return a `bool`.
* Results are either `true` or `false`.
* `=` is used for assignment.
* `==` is used for comparison.
* Strings can be compared.
* String comparison is case-sensitive.
* Different data types cannot be compared directly.
* Comparison operators are commonly used with `if`, `switch`, and loops.

---

## Summary

| Operator | Meaning                  |
| -------- | ------------------------ |
| `==`     | Equal To                 |
| `!=`     | Not Equal To             |
| `>`      | Greater Than             |
| `<`      | Less Than                |
| `>=`     | Greater Than or Equal To |
| `<=`     | Less Than or Equal To    |

Comparison operators help programs make decisions by evaluating conditions and returning `true` or `false`.
