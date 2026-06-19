# 03_switch

## What is switch?

A `switch` statement is used when you need to compare one value against multiple possible values.

It is often cleaner and easier to read than a long `if else if else` chain.

---

## Why Use switch?

Instead of:

```go
if role == "admin" {
	fmt.Println("Full Access")
} else if role == "editor" {
	fmt.Println("Limited Access")
} else if role == "viewer" {
	fmt.Println("Read Only Access")
}
```

You can write:

```go
switch role {
case "admin":
	fmt.Println("Full Access")

case "editor":
	fmt.Println("Limited Access")

case "viewer":
	fmt.Println("Read Only Access")
}
```

This is cleaner and easier to maintain.

---

## Basic Syntax

```go
switch expression {
case value1:
	// code

case value2:
	// code

default:
	// code
}
```

---

## Basic Example

```go
day := "Monday"

switch day {
case "Monday":
	fmt.Println("Start of Week")

case "Friday":
	fmt.Println("Weekend Coming")

default:
	fmt.Println("Normal Day")
}
```

Output:

```text
Start of Week
```

---

## How switch Works

Go checks each case from top to bottom.

```go
number := 2

switch number {
case 1:
	fmt.Println("One")

case 2:
	fmt.Println("Two")

case 3:
	fmt.Println("Three")
}
```

Output:

```text
Two
```

When a matching case is found, Go executes it and stops.

---

## default Case

The `default` block runs when no case matches.

```go
role := "guest"

switch role {
case "admin":
	fmt.Println("Admin")

case "editor":
	fmt.Println("Editor")

default:
	fmt.Println("Unknown Role")
}
```

Output:

```text
Unknown Role
```

---

## Multiple Values in One Case

Multiple values can share the same code block.

```go
day := "Sunday"

switch day {
case "Saturday", "Sunday":
	fmt.Println("Weekend")

default:
	fmt.Println("Working Day")
}
```

Output:

```text
Weekend
```

---

## Switch with Numbers

```go
month := 1

switch month {
case 1:
	fmt.Println("January")

case 2:
	fmt.Println("February")

case 3:
	fmt.Println("March")

default:
	fmt.Println("Unknown Month")
}
```

---

## Switch Without Expression

A switch can be used without a variable.

This behaves similarly to an `if else if` chain.

```go
score := 85

switch {
case score >= 90:
	fmt.Println("Grade A+")

case score >= 80:
	fmt.Println("Grade A")

case score >= 70:
	fmt.Println("Grade B")

default:
	fmt.Println("Fail")
}
```

Output:

```text
Grade A
```

---

## fallthrough

Normally Go stops after the first matching case.

```go
number := 1

switch number {
case 1:
	fmt.Println("One")

case 2:
	fmt.Println("Two")
}
```

Output:

```text
One
```

Using `fallthrough` forces the next case to execute.

```go
number := 1

switch number {
case 1:
	fmt.Println("One")
	fallthrough

case 2:
	fmt.Println("Two")
}
```

Output:

```text
One
Two
```

### Important

- `fallthrough` executes the next case even if its condition does not match.
- It only moves to the immediately next case.
- It is rarely used in production Go code.

---

## Go Does Not Need break

In many languages, switch statements require `break`.

Example from other languages:

```javascript
switch (day) {
  case "Monday":
    console.log("Monday");
    break;
}
```

In Go:

```go
switch day {
case "Monday":
	fmt.Println("Monday")
}
```

Go automatically stops after a matching case.

No `break` is required.

---

## Real Backend Examples

### HTTP Method Routing

```go
switch method {
case "GET":
	fmt.Println("Fetch Data")

case "POST":
	fmt.Println("Create Data")

case "DELETE":
	fmt.Println("Delete Data")

default:
	fmt.Println("Unsupported Method")
}
```

---

### User Role Check

```go
switch role {
case "admin":
	fmt.Println("Full Access")

case "editor":
	fmt.Println("Limited Access")

case "viewer":
	fmt.Println("Read Only Access")

default:
	fmt.Println("No Access")
}
```

---

### Package Registry Type

```go
switch registry {
case "npm":
	fmt.Println("Node.js Package")

case "pypi":
	fmt.Println("Python Package")

case "maven":
	fmt.Println("Java Package")

default:
	fmt.Println("Unknown Registry")
}
```

---

### Vulnerability Severity

```go
switch severity {
case "critical":
	fmt.Println("Immediate Action Required")

case "high":
	fmt.Println("Fix Soon")

case "medium":
	fmt.Println("Monitor")

default:
	fmt.Println("Low Priority")
}
```

---

## if vs switch

Use `if` for conditions:

```go
if age >= 18 {
	fmt.Println("Adult")
}
```

Use `switch` for matching values:

```go
switch role {
case "admin":
case "editor":
case "viewer":
}
```

### Easy Rule

```text
if     → Is this condition true?

switch → Which value is this?
```

---

## Common Mistakes

### Forgetting default

Not required, but usually recommended.

```go
switch role {
case "admin":
	fmt.Println("Admin")

default:
	fmt.Println("Unknown")
}
```

---

### Using fallthrough Unnecessarily

Most switch statements do not need `fallthrough`.

```go
fallthrough
```

Use it only when you intentionally want to execute the next case.

---

### Using switch for Simple Conditions

❌ Not ideal

```go
switch age {
case 18:
	fmt.Println("Adult")
}
```

✅ Better

```go
if age >= 18 {
	fmt.Println("Adult")
}
```

---

## Key Points

- `switch` compares one value against many possible values.
- It is cleaner than long `if else if else` chains.
- Go automatically stops after a matching case.
- `default` runs when no case matches.
- Multiple values can be grouped in one case.
- `switch {}` can replace complex `if else if` chains.
- `fallthrough` forces execution of the next case.
- No `break` is normally required.
- Commonly used for roles, statuses, HTTP methods, registry types, and enums.

---
