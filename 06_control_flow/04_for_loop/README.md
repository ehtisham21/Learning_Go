# 04_for_loop

## What is a for Loop?

A `for` loop is used to repeat a block of code multiple times.

Go has only one loop keyword:

```go
for
```

Unlike many languages, Go does not have separate `while` or `do while` loops.

The `for` keyword is used for all looping scenarios.

---

## Why Use Loops?

Without a loop:

```go
fmt.Println(1)
fmt.Println(2)
fmt.Println(3)
fmt.Println(4)
fmt.Println(5)
```

With a loop:

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

Output:

```text
1
2
3
4
5
```

---

## Basic Syntax

```go
for initialization; condition; update {
	// code
}
```

Example:

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

---

## Understanding the Three Parts

### Initialization

Runs once before the loop starts.

```go
i := 1
```

---

### Condition

Checked before every iteration.

```go
i <= 5
```

If true, the loop continues.

If false, the loop stops.

---

### Update

Runs after every iteration.

```go
i++
```

Equivalent to:

```go
i = i + 1
```

---

## Counting Backwards

```go
for i := 5; i >= 1; i-- {
	fmt.Println(i)
}
```

Output:

```text
5
4
3
2
1
```

---

## Incrementing by More Than One

```go
for i := 0; i <= 10; i += 2 {
	fmt.Println(i)
}
```

Output:

```text
0
2
4
6
8
10
```

---

## While-Style Loop

Go does not have a `while` keyword.

Instead:

```go
count := 1

for count <= 5 {
	fmt.Println(count)
	count++
}
```

This behaves like a `while` loop in other languages.

---

## Infinite Loop

```go
for {
	fmt.Println("Running...")
}
```

This loop never stops.

It is commonly used for:

- Web servers
- Workers
- Message consumers
- Background services

---

## break

The `break` statement immediately stops a loop.

Example:

```go
for i := 1; i <= 10; i++ {

	if i == 5 {
		break
	}

	fmt.Println(i)
}
```

Output:

```text
1
2
3
4
```

---

## continue

The `continue` statement skips the current iteration.

Example:

```go
for i := 1; i <= 5; i++ {

	if i == 3 {
		continue
	}

	fmt.Println(i)
}
```

Output:

```text
1
2
4
5
```

---

## Variable Scope

Variables declared inside a loop only exist inside that loop.

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}

fmt.Println(i)
```

Produces:

```text
undefined: i
```

---

## Looping Through a Slice

Traditional approach:

```go
packages := []string{
	"django",
	"requests",
	"flask",
}

for i := 0; i < len(packages); i++ {
	fmt.Println(packages[i])
}
```

Output:

```text
django
requests
flask
```

---

## Backend Example: Retry Logic

```go
for attempt := 1; attempt <= 3; attempt++ {
	fmt.Println("Retry:", attempt)
}
```

Output:

```text
Retry: 1
Retry: 2
Retry: 3
```

---

## Backend Example: Processing Data

```go
packageVersions := []string{
	"1.0.0",
	"1.1.0",
	"2.0.0",
}

for i := 0; i < len(packageVersions); i++ {
	fmt.Println("Processing:", packageVersions[i])
}
```

---

## Common Mistakes

### Wrong Condition

❌ Wrong

```go
for i := 1; i >= 5; i++ {
	fmt.Println(i)
}
```

The condition is false immediately.

The loop never runs.

---

### Forgetting the Update

❌ Wrong

```go
i := 1

for i <= 5 {
	fmt.Println(i)
}
```

This creates an infinite loop because `i` never changes.

---

### Using a Variable Outside Its Scope

❌ Wrong

```go
for i := 1; i <= 5; i++ {
}

fmt.Println(i)
```

Produces:

```text
undefined: i
```

---

## Real Go Usage

While traditional loops are important to understand, Go developers often use:

```go
for _, item := range items {
	fmt.Println(item)
}
```

This is Go's foreach-style loop and is the most common way to iterate over slices, arrays, maps, and strings.

You will learn this in:

```text
07_range
```

---

## Key Points

- Go has only one loop keyword: `for`.
- Go does not have `while` or `do while`.
- Traditional syntax:

```go
for init; condition; update {
}
```

- Initialization runs once.
- Condition is checked before every iteration.
- Update runs after every iteration.
- `for condition {}` acts like a while loop.
- `for {}` creates an infinite loop.
- `break` exits a loop.
- `continue` skips the current iteration.
- Variables declared inside a loop only exist inside that loop.
- Loops are heavily used for processing data, retries, workers, and APIs.

---
