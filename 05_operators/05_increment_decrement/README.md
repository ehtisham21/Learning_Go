# Increment and Decrement Operators in Go

## Overview

Increment and decrement operators are used to increase or decrease a value by `1`.

They are commonly used in:

- Counters
- Loops
- Tracking values
- Game scores
- Visitor counts

---

## Operators

| Operator | Meaning        | Equivalent |
| -------- | -------------- | ---------- |
| `++`     | Increment by 1 | `x += 1`   |
| `--`     | Decrement by 1 | `x -= 1`   |

---

# Increment Operator (`++`)

The increment operator increases a value by `1`.

Instead of:

```go
count = count + 1
```

You can write:

```go
count++
```

---

## Example

```go
count := 1

count++

fmt.Println(count)
```

Output:

```text
2
```

---

## Multiple Increments

```go
count := 1

count++
count++
count++

fmt.Println(count)
```

Output:

```text
4
```

Calculation:

```text
1 → 2 → 3 → 4
```

---

# Decrement Operator (`--`)

The decrement operator decreases a value by `1`.

Instead of:

```go
count = count - 1
```

You can write:

```go
count--
```

---

## Example

```go
count := 5

count--

fmt.Println(count)
```

Output:

```text
4
```

---

## Multiple Decrements

```go
count := 5

count--
count--
count--

fmt.Println(count)
```

Output:

```text
2
```

Calculation:

```text
5 → 4 → 3 → 2
```

---

# Real-World Examples

## Website Visitors

```go
visitors := 100

visitors++

fmt.Println(visitors)
```

Output:

```text
101
```

A new visitor increases the count by 1.

---

## Game Lives

```go
lives := 3

lives--

fmt.Println(lives)
```

Output:

```text
2
```

A player loses one life.

---

## Product Stock

```go
stock := 10

stock--

fmt.Println(stock)
```

Output:

```text
9
```

One item was sold.

---

# Increment and Decrement in Loops

You will often see increment operators inside loops.

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

Output:

```text
0
1
2
3
4
```

The variable `i` increases by 1 after each iteration.

---

# Important Go Rule

Go handles increment and decrement differently from many other languages.

In Go:

```go
x++
```

and

```go
x--
```

are **statements**, not expressions.

---

## Valid

```go
x := 5

x++

fmt.Println(x)
```

Output:

```text
6
```

---

## Invalid

```go
y := x++
```

Error.

Because `x++` does not return a value.

---

## Invalid

```go
fmt.Println(x++)
```

Error.

Because `x++` cannot be used inside another expression.

---

# Statement vs Expression

## Expression

Produces a value.

Example:

```go
10 + 5
```

Produces:

```text
15
```

---

```go
age >= 18
```

Produces:

```text
true
```

---

## Statement

Performs an action.

Example:

```go
x++
```

Updates the variable but does not return a value.

---

# No Prefix Increment in Go

Some languages allow:

```cpp
++x
```

Go does not.

❌ Invalid

```go
++x
```

---

Only this is allowed:

```go
x++
```

---

# No Prefix Decrement in Go

❌ Invalid

```go
--x
```

---

Only this is allowed:

```go
x--
```

---

# Alternative Using Assignment Operators

Increment:

```go
count += 1
```

Equivalent to:

```go
count++
```

---

Decrement:

```go
count -= 1
```

Equivalent to:

```go
count--
```

---

# Common Beginner Mistakes

❌ Invalid

```go
y := x++
```

❌ Invalid

```go
fmt.Println(x++)
```

❌ Invalid

```go
++x
```

❌ Invalid

```go
--x
```

---

✅ Correct

```go
x++
x--
```

---

# Key Points

- `++` increases a value by 1.
- `--` decreases a value by 1.
- Commonly used in loops and counters.
- `x++` is valid.
- `x--` is valid.
- `++x` is not supported in Go.
- `--x` is not supported in Go.
- `x++` and `x--` are statements, not expressions.
- They cannot be assigned to variables or used inside function calls.

---

# Summary

| Operator | Meaning        | Equivalent |
| -------- | -------------- | ---------- |
| `++`     | Increment by 1 | `x += 1`   |
| `--`     | Decrement by 1 | `x -= 1`   |

Examples:

```go
count++
count--
```

These operators provide a simple and readable way to increase or decrease values by one.
