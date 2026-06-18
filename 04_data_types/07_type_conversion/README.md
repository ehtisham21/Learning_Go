# 07 - Type Conversion in Go

## What is Type Conversion?

Type conversion means changing a value from one data type to another.

Example:

```go
age := 25

price := float64(age)
```

Here:

```text
int -> float64
```

---

## Why Do We Need Type Conversion?

Go is a strongly typed language.

Different data types cannot be mixed directly.

Example:

```go
age := 25
price := 10.99

fmt.Println(age + price)
```

This gives an error because:

```text
age   = int
price = float64
```

Go requires explicit conversion.

Correct:

```go
fmt.Println(float64(age) + price)
```

---

## Basic Type Conversion

Syntax:

```go
targetType(value)
```

Example:

```go
float64(age)
int(price)
```

---

## Integer to Float

```go
age := 25

price := float64(age)
```

Result:

```text
25.0
```

---

## Float to Integer

```go
price := 10.99

num := int(price)
```

Result:

```text
10
```

Important:

```text
Decimal part is removed.
```

It is not rounded.

---

## When NOT to Use strconv

Use normal conversion for numeric types.

Examples:

```go
float64(age)
int(price)
int64(age)
```

---

## What is strconv?

`strconv` stands for:

```text
String Conversion
```

It is used when converting between:

```text
string ↔ int
string ↔ float64
string ↔ bool
```

---

## String to Int

```go
text := "25"

num, _ := strconv.Atoi(text)
```

Result:

```text
25
```

---

## Int to String

```go
age := 25

text := strconv.Itoa(age)
```

Result:

```text
"25"
```

---

## String to Float

```go
price := "10.99"

value, _ := strconv.ParseFloat(price, 64)
```

Result:

```text
10.99
```

---

## Float to String

```go
text := strconv.FormatFloat(10.99, 'f', 2, 64)
```

Result:

```text
"10.99"
```

---

## String to Bool

```go
isAdmin, _ := strconv.ParseBool("true")
```

Result:

```text
true
```

---

## Bool to String

```go
text := strconv.FormatBool(true)
```

Result:

```text
"true"
```

---

## Rune to String

```go
letter := 'A'

fmt.Println(string(letter))
```

Output:

```text
A
```

No `strconv` needed.

---

## Byte to String

```go
b := byte(65)

fmt.Println(string(b))
```

Output:

```text
A
```

No `strconv` needed.

---

## String to Byte Slice

```go
text := "Hello"

data := []byte(text)
```

Output:

```text
[72 101 108 108 111]
```

---

## Byte Slice to String

```go
data := []byte{72, 101, 108, 108, 111}

text := string(data)
```

Output:

```text
Hello
```

---

## Common Mistake

Many beginners write:

```go
age := 25

text := string(age)
```

Expecting:

```text
"25"
```

But this is wrong.

`string(age)` treats `age` as a Unicode value.

Correct:

```go
text := strconv.Itoa(age)
```

---

## When to Use strconv

Use `strconv` whenever a string is involved.

Examples:

```go
"25"  -> 25
25    -> "25"

"10.99" -> 10.99
10.99   -> "10.99"

"true" -> true
true   -> "true"
```

---

## When NOT to Use strconv

Use normal conversion for numeric types.

Examples:

```go
float64(age)
int(price)
int64(age)
uint(age)
```

---

## Key Points

* Type conversion changes one data type into another.
* Use `targetType(value)` for numeric conversions.
* Go does not automatically convert types.
* Use `strconv` when converting between strings and other types.
* Use `strconv.Atoi()` for string → int.
* Use `strconv.Itoa()` for int → string.
* Use `strconv.ParseFloat()` for string → float64.
* Use `strconv.FormatFloat()` for float64 → string.
* Use `strconv.ParseBool()` for string → bool.
* Use `strconv.FormatBool()` for bool → string.
* Do not use `string(25)` to get `"25"`.

## Easy Summary

```text
Numeric ↔ Numeric
        ↓
Normal Conversion

int -> float64
float64 -> int
int -> int64
```

```text
String ↔ Anything
        ↓
strconv

"25" -> 25
25 -> "25"

"10.99" -> 10.99
10.99 -> "10.99"
```
