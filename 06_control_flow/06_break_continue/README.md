# 06_break_continue

## What are break and continue?

`break` and `continue` are used to control loop execution.

- `break` stops the loop.
- `continue` skips the current iteration and moves to the next one.

---

## break

`break` immediately exits the current loop.

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

When `i == 5`, the loop stops completely.

---

## continue

`continue` skips the current iteration.

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

When `i == 3`, Go skips `fmt.Println(i)` and moves to the next iteration.

---

## Difference Between break and continue

| Keyword    | Meaning                     |
| ---------- | --------------------------- |
| `break`    | Stop the loop completely    |
| `continue` | Skip current iteration only |

Example with `break`:

```go
for i := 1; i <= 5; i++ {
	if i == 3 {
		break
	}

	fmt.Println(i)
}
```

Output:

```text
1
2
```

Example with `continue`:

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

## Search Example Using break

Use `break` when you find what you are looking for and do not need to continue.

```go
users := []string{"Ali", "Ahmed", "Sara"}
target := "Ahmed"

for _, user := range users {
	if user == target {
		fmt.Println("User found:", user)
		break
	}

	fmt.Println("Checking:", user)
}
```

Output:

```text
Checking: Ali
User found: Ahmed
```

---

## Skip Invalid Data Using continue

Use `continue` when you want to skip invalid data.

```go
versions := []string{"1.0.0", "", "2.0.0"}

for _, version := range versions {
	if version == "" {
		continue
	}

	fmt.Println("Processing:", version)
}
```

Output:

```text
Processing: 1.0.0
Processing: 2.0.0
```

---

## break in Nested Loops

In nested loops, `break` stops only the loop where it is written.

```go
for i := 1; i <= 3; i++ {
	for j := 1; j <= 5; j++ {
		if j == 3 {
			break
		}

		fmt.Println(i, j)
	}
}
```

Output:

```text
1 1
1 2
2 1
2 2
3 1
3 2
```

Here, `break` stops only the inner loop.

The outer loop continues.

---

## continue in Nested Loops

In nested loops, `continue` skips only the current iteration of the loop where it is written.

```go
for i := 1; i <= 2; i++ {
	for j := 1; j <= 4; j++ {
		if j == 2 {
			continue
		}

		fmt.Println(i, j)
	}
}
```

Output:

```text
1 1
1 3
1 4
2 1
2 3
2 4
```

Here, only `j == 2` is skipped.

---

## Labeled break

Use labeled `break` when you want to stop an outer loop from inside an inner loop.

```go
OuterBreak:
for i := 1; i <= 3; i++ {
	for j := 1; j <= 5; j++ {
		if j == 3 {
			break OuterBreak
		}

		fmt.Println(i, j)
	}
}
```

Output:

```text
1 1
1 2
```

Here, both loops stop.

---

## Labeled continue

Use labeled `continue` when you want to skip to the next iteration of an outer loop.

```go
OuterContinue:
for i := 1; i <= 3; i++ {
	for j := 1; j <= 3; j++ {
		if j == 2 {
			continue OuterContinue
		}

		fmt.Println(i, j)
	}
}
```

Output:

```text
1 1
2 1
3 1
```

When `j == 2`, Go jumps to the next `i`.

---

## Backend Example: Retry Until Success

```go
successAt := 3

for attempt := 1; attempt <= 5; attempt++ {
	fmt.Println("Attempt:", attempt)

	if attempt == successAt {
		fmt.Println("Success, stopping retries")
		break
	}
}
```

Output:

```text
Attempt: 1
Attempt: 2
Attempt: 3
Success, stopping retries
```

---

## Backend Example: Skip Invalid Packages

```go
packages := []string{"django", "", "flask", "", "requests"}

for _, pkg := range packages {
	if pkg == "" {
		continue
	}

	fmt.Println("Valid package:", pkg)
}
```

Output:

```text
Valid package: django
Valid package: flask
Valid package: requests
```

---

## Common Mistakes

### Thinking break stops all loops

`break` only stops the current loop.

Use labeled `break` if you want to stop outer loops.

---

### Thinking continue stops the loop

`continue` does not stop the loop.

It only skips the current iteration.

---

### Writing important code after continue

```go
if version == "" {
	continue
}

fmt.Println("Processing:", version)
```

For empty versions, anything after `continue` is skipped.

---

## Key Points

- `break` stops the current loop immediately.
- `continue` skips the current iteration.
- `break` exits the loop.
- `continue` keeps the loop running.
- In nested loops, both affect only the loop they are inside.
- Use labeled `break` to stop outer loops.
- Use labeled `continue` to continue outer loops.
- Use `break` when the work is done.
- Use `continue` when data should be skipped.

---
