# 04 - Strings in Go

## What is a String?

A string is a sequence of characters used to store text data.

Examples:

```go
"Hello"
"Go Language"
"Ehtisham"
"Backend Developer"
"hello@example.com"
```

Strings are one of the most commonly used data types in Go.

---

## String Type

In Go, the string type is:

```go
string
```

Example:

```go
var name string = "Ehtisham"
```

---

## Why Use Strings?

Strings are used to store:

- Names
- Email addresses
- Messages
- URLs
- File paths
- JSON data
- API responses

Example:

```go
username := "admin"
email := "admin@example.com"
```

---

## String Declaration

### Using var

```go
var name string = "Ehtisham"
```

### Type Inference

```go
var name = "Ehtisham"
```

Go automatically infers the type.

### Short Variable Declaration

```go
name := "Ehtisham"
```

This is the most common style inside functions.

---

## Zero Value

The zero value of a string is an empty string.

```go
""
```

Example:

```go
var country string

fmt.Println(country)
```

Output:

```text

```

Nothing is printed because the string is empty.

---

## Checking the Type

Use `%T` to check the variable type.

```go
name := "Ehtisham"

fmt.Printf("%T\n", name)
```

Output:

```text
string
```

---

## String Length

Use the `len()` function to get the length of a string.

Example:

```go
name := "Ehtisham"

fmt.Println(len(name))
```

Output:

```text
8
```

---

## String Concatenation

Concatenation means joining strings together.

Use the `+` operator.

Example:

```go
firstName := "Ehtisham"
lastName := "Butt"

fullName := firstName + " " + lastName

fmt.Println(fullName)
```

Output:

```text
Ehtisham Butt
```

---

## Multi-line Strings

Go supports raw strings using backticks.

Example:

```go
message := `Hello
Welcome to Go
Learning Strings`
```

Output:

```text
Hello
Welcome to Go
Learning Strings
```

Raw strings preserve formatting exactly as written.

---

## Escape Characters

Escape characters start with a backslash (`\`).

### New Line

```go
fmt.Println("Hello\nWorld")
```

Output:

```text
Hello
World
```

### Tab

```go
fmt.Println("Name:\tEhtisham")
```

Output:

```text
Name:    Ehtisham
```

### Double Quotes

```go
fmt.Println("He said \"Hello\"")
```

Output:

```text
He said "Hello"
```

### Backslash

```go
fmt.Println("C:\\Users\\Admin")
```

Output:

```text
C:\Users\Admin
```

---

## String Comparison

Strings can be compared using comparison operators.

### Equal To

```go
"go" == "go"
```

Output:

```text
true
```

### Not Equal To

```go
"go" != "python"
```

Output:

```text
true
```

Example:

```go
fmt.Println("go" == "go")
fmt.Println("go" == "python")
fmt.Println("go" != "python")
```

---

## Accessing Characters

Strings can be accessed using indexes.

Example:

```go
language := "Go"

fmt.Println(language[0])
fmt.Println(language[1])
```

Output:

```text
71
111
```

These values are ASCII byte values.

We will learn proper character handling later with runes.

---

## String Formatting

Use `fmt.Printf()` to format output.

Example:

```go
name := "Ehtisham"
age := 25

fmt.Printf("Name: %s, Age: %d\n", name, age)
```

Output:

```text
Name: Ehtisham, Age: 25
```

---

## Common Format Specifiers

| Specifier | Description   |
| --------- | ------------- |
| `%s`      | String        |
| `%d`      | Integer       |
| `%f`      | Float         |
| `%t`      | Boolean       |
| `%T`      | Variable Type |

Example:

```go
fmt.Printf("%s\n", name)
fmt.Printf("%d\n", age)
fmt.Printf("%T\n", name)
```

---

## String Immutability

Strings in Go are immutable.

This means a string cannot be modified directly.

Invalid:

```go
name := "Go"

name[0] = 'P'
```

This causes a compile-time error.

Instead, assign a new string:

```go
name = "Pro Go"
```

---

## Key Points

- Strings store text data.
- The Go string type is `string`.
- Strings are written inside double quotes.
- The zero value of a string is an empty string (`""`).
- Use `len()` to get string length.
- Use `+` to join strings.
- Backticks create multi-line raw strings.
- Escape characters start with `\`.
- Strings can be compared using `==` and `!=`.
- Strings are immutable.
- `%s` is used for string formatting.
- Characters accessed by index return byte values.
