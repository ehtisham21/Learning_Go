# Logical Operators in Go

## Overview

Logical operators are used to combine or modify boolean (`bool`) values.

They help us work with multiple conditions and make decisions in our programs.

Logical operators are commonly used with:

- `if` statements
- loops
- validations
- authentication checks
- permission checks

---

## Logical Operators

| Operator      | Name | Meaning                             |
| ------------- | ---- | ----------------------------------- |
| `&&`          | AND  | Both conditions must be true        |
| `pipe symbol` | OR   | At least one condition must be true |
| `!`           | NOT  | Reverses a boolean value            |

---

# AND Operator (`&&`)

The AND operator returns `true` only when both conditions are true.

## Syntax

```go
condition1 && condition2
```

## Example

```go
age := 25

fmt.Println(age >= 18 && age <= 60)
```

Output:

```text
true
```

Both conditions are true:

```text
25 >= 18 → true
25 <= 60 → true
```

Result:

```text
true && true → true
```

---

## Login Example

```go
isLoggedIn := true
isAdmin := true

fmt.Println(isLoggedIn && isAdmin)
```

Output:

```text
true
```

---

## AND Truth Table

| A     | B     | A && B |
| ----- | ----- | ------ |
| true  | true  | true   |
| true  | false | false  |
| false | true  | false  |
| false | false | false  |

---

# OR Operator (`||`)

The OR operator returns `true` if at least one condition is true.

## Syntax

```go
condition1 || condition2
```

## Example

```go
isAdmin := false
isOwner := true

fmt.Println(isAdmin || isOwner)
```

Output:

```text
true
```

Because one condition is true.

---

## Access Permission Example

```go
isManager := true
isAdmin := false

fmt.Println(isManager || isAdmin)
```

Output:

```text
true
```

---

## OR Truth Table

| A     | B     | A     |     | B   |
| ----- | ----- | ----- | --- | --- |
| true  | true  | true  |
| true  | false | true  |
| false | true  | true  |
| false | false | false |

---

# NOT Operator (`!`)

The NOT operator reverses a boolean value.

## Syntax

```go
!condition
```

## Example

```go
isLoggedIn := true

fmt.Println(!isLoggedIn)
```

Output:

```text
false
```

---

Another example:

```go
isLoggedIn := false

fmt.Println(!isLoggedIn)
```

Output:

```text
true
```

---

## NOT Truth Table

| A     | !A    |
| ----- | ----- |
| true  | false |
| false | true  |

---

# Combining Conditions

Logical operators are often combined with comparison operators.

```go
age := 25
salary := 70000

fmt.Println(age >= 18 && salary > 50000)
```

Output:

```text
true
```

---

Another example:

```go
age := 16
hasPermission := true

fmt.Println(age >= 18 || hasPermission)
```

Output:

```text
true
```

Because one condition is true.

---

# Using Parentheses

Parentheses make complex conditions easier to read.

```go
result := (age >= 18 && age <= 60) && salary > 50000
```

This is easier to understand than writing everything on one line without grouping.

---

# Short-Circuit Evaluation

Go stops evaluating conditions when it already knows the final result.

## AND Example

```go
false && something
```

Go immediately returns:

```text
false
```

The second condition is not checked.

---

## OR Example

```go
true || something
```

Go immediately returns:

```text
true
```

The second condition is not checked.

This behavior is called **short-circuit evaluation**.

---

# Real-World Examples

## Driving License Check

```go
age := 20
hasLicense := true

fmt.Println(age >= 18 && hasLicense)
```

Output:

```text
true
```

---

## Discount Eligibility

```go
isMember := true
hasCoupon := false

fmt.Println(isMember || hasCoupon)
```

Output:

```text
true
```

---

## Show Login Button

```go
isLoggedIn := false

fmt.Println(!isLoggedIn)
```

Output:

```text
true
```

Meaning:

```text
Show Login Button
```

---

# Common Beginner Mistake

❌ Invalid

```go
age >= 18 and age <= 60
```

Go does not use:

```text
and
or
not
```

✅ Correct

```go
age >= 18 && age <= 60
age >= 18 || hasPermission
!isLoggedIn
```

---

# Key Points

- Logical operators work with boolean values.
- `&&` requires all conditions to be true.
- `||` requires at least one condition to be true.
- `!` reverses a boolean value.
- Commonly used with comparison operators.
- Frequently used in `if` statements and loops.
- Go uses `&&`, `||`, and `!` instead of `and`, `or`, and `not`.
- Go performs short-circuit evaluation for better performance.

---
