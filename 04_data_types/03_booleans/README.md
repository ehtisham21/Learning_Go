# 03 - Booleans in Go

## What is a Boolean?

A boolean is a data type that can store only two values:

```go
true
false
```

Booleans are commonly used to represent:

* Yes / No
* On / Off
* Success / Failure
* True / False

Examples:

```go
isLoggedIn := true
isAdmin := false
```

---

## Boolean Type

In Go, the boolean type is:

```go
bool
```

Example:

```go
var isActive bool = true
```

---

## Boolean Declaration

### Using var

```go
var isActive bool = true
```

### Type Inference

```go
var isActive = true
```

Go automatically determines the type.

### Short Variable Declaration

```go
isActive := true
```

This is the most common style inside functions.

---

## Zero Value

The default value of a boolean is:

```go
false
```

Example:

```go
var isAdmin bool

fmt.Println(isAdmin)
```

Output:

```go
false
```

---

## Printing Boolean Values

```go
isLoggedIn := true
isAdmin := false

fmt.Println(isLoggedIn)
fmt.Println(isAdmin)
```

Output:

```go
true
false
```

---

## Comparison Operators

Comparison operators return boolean values.

### Equal To

```go
10 == 10
```

Output:

```go
true
```

### Not Equal To

```go
10 != 5
```

Output:

```go
true
```

### Greater Than

```go
20 > 10
```

Output:

```go
true
```

### Less Than

```go
5 < 10
```

Output:

```go
true
```

### Greater Than or Equal To

```go
10 >= 10
```

Output:

```go
true
```

### Less Than or Equal To

```go
5 <= 10
```

Output:

```go
true
```

---

## Logical Operators

Logical operators combine multiple boolean values.

### AND (&&)

Both conditions must be true.

```go
age := 20
hasID := true

fmt.Println(age >= 18 && hasID)
```

Output:

```go
true
```

Example:

```go
true && false
```

Output:

```go
false
```

---

### OR (||)

At least one condition must be true.

```go
true || false
```

Output:

```go
true
```

---

### NOT (!)

Reverses a boolean value.

```go
isLoggedIn := true

fmt.Println(!isLoggedIn)
```

Output:

```go
false
```

---

## Using Booleans with If Statements

Booleans are commonly used in conditions.

Example:

```go
isLoggedIn := true

if isLoggedIn {
	fmt.Println("Welcome!")
}
```

Output:

```go
Welcome!
```

Another example:

```go
age := 16

if age >= 18 {
	fmt.Println("Adult")
}
```

Nothing is printed because the condition is false.

---

## Checking Variable Types

Use `%T` to check the type of a variable.

```go
isActive := true

fmt.Printf("%T\n", isActive)
```

Output:

```go
bool
```

---

## Good Boolean Variable Names

Use names that clearly describe the condition.

Good examples:

```go
isLoggedIn := true
isAdmin := false
hasPermission := true
isEmailVerified := false
```

Avoid unclear names such as:

```go
flag := true
x := false
```

---

## Key Points

* Boolean values can only be `true` or `false`.
* The Go boolean type is `bool`.
* The zero value of a boolean is `false`.
* Comparison operators return boolean values.
* `&&` means AND.
* `||` means OR.
* `!` means NOT.
* Booleans are heavily used in `if` statements.
* Use meaningful names such as `isLoggedIn`, `hasPermission`, and `isAdmin`.
* Use `%T` to check the variable type.
