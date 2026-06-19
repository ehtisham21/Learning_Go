# Assignment Operators in Go

## Overview

Assignment operators are used to assign values to variables or update existing values.

They help make code shorter, cleaner, and easier to read.

---

## Assignment Operators

| Operator | Meaning             | Example  |
| -------- | ------------------- | -------- |
| `=`      | Assign              | `x = 10` |
| `+=`     | Add and Assign      | `x += 5` |
| `-=`     | Subtract and Assign | `x -= 5` |
| `*=`     | Multiply and Assign | `x *= 2` |
| `/=`     | Divide and Assign   | `x /= 2` |
| `%=`     | Modulus and Assign  | `x %= 3` |

---

# Basic Assignment (`=`)

Assigns a value to a variable.

```go
score := 100

score = 200

fmt.Println(score)
```

Output:

```text
200
```

The previous value is replaced with the new value.

---

# Add and Assign (`+=`)

Instead of:

```go
score = score + 10
```

You can write:

```go
score += 10
```

Example:

```go
score := 100

score += 10

fmt.Println(score)
```

Output:

```text
110
```

---

# Subtract and Assign (`-=`)

Instead of:

```go
balance = balance - 200
```

You can write:

```go
balance -= 200
```

Example:

```go
balance := 1000

balance -= 200

fmt.Println(balance)
```

Output:

```text
800
```

---

# Multiply and Assign (`*=`)

Instead of:

```go
price = price * 2
```

You can write:

```go
price *= 2
```

Example:

```go
price := 50

price *= 2

fmt.Println(price)
```

Output:

```text
100
```

---

# Divide and Assign (`/=`)

Instead of:

```go
amount = amount / 4
```

You can write:

```go
amount /= 4
```

Example:

```go
amount := 100

amount /= 4

fmt.Println(amount)
```

Output:

```text
25
```

---

# Modulus and Assign (`%=`)

Instead of:

```go
num = num % 3
```

You can write:

```go
num %= 3
```

Example:

```go
num := 10

num %= 3

fmt.Println(num)
```

Output:

```text
1
```

---

# Why Use Assignment Operators?

They reduce repetition and make code cleaner.

Without assignment operators:

```go
score = score + 10
score = score - 5
score = score * 2
```

With assignment operators:

```go
score += 10
score -= 5
score *= 2
```

The result is the same, but the code is easier to read.

---

# Real-World Examples

## Bank Deposit

```go
balance := 5000

balance += 1000

fmt.Println(balance)
```

Output:

```text
6000
```

---

## Shopping Discount

```go
price := 1000

price -= 200

fmt.Println(price)
```

Output:

```text
800
```

---

## Salary Increase

```go
salary := 50000

salary *= 2

fmt.Println(salary)
```

Output:

```text
100000
```

---

## Bill Split

```go
bill := 1000

bill /= 4

fmt.Println(bill)
```

Output:

```text
250
```

---

# Difference Between `=` and `:=`

## `:=`

Creates a new variable.

```go
score := 100
```

---

## `=`

Updates an existing variable.

```go
score = 200
```

---

Example:

```go
score := 100

score = 200
```

The variable is created once using `:=` and updated later using `=`.

---

# Common Beginner Mistake

❌ Wrong

```go
x := 10

x := x + 5
```

Error:

```text
no new variables on left side of :=
```

Because `x` already exists.

✅ Correct

```go
x := 10

x = x + 5
```

Or:

```go
x += 5
```

---

# Equivalent Operations

| Shortcut | Equivalent  |
| -------- | ----------- |
| `x += 5` | `x = x + 5` |
| `x -= 5` | `x = x - 5` |
| `x *= 2` | `x = x * 2` |
| `x /= 2` | `x = x / 2` |
| `x %= 3` | `x = x % 3` |

---

# Key Points

- `=` assigns or updates a value.
- `+=` adds and updates.
- `-=` subtracts and updates.
- `*=` multiplies and updates.
- `/=` divides and updates.
- `%=` gets the remainder and updates.
- Assignment operators reduce repetition.
- Use `:=` only when creating a variable.
- Use `=` and assignment operators when updating a variable.

---

# Summary

| Operator | Purpose             |
| -------- | ------------------- |
| `=`      | Assign a value      |
| `+=`     | Add and update      |
| `-=`     | Subtract and update |
| `*=`     | Multiply and update |
| `/=`     | Divide and update   |
| `%=`     | Modulus and update  |

Assignment operators make code shorter, cleaner, and easier to maintain.
