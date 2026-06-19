# 05 - Runes in Go

## What is a Rune?

A rune represents **one character** in Go.

Examples:

```go
'A'
'B'
'ا'
'你'
'😀'
```

A rune can store:

* English letters
* Urdu letters
* Arabic letters
* Chinese characters
* Emojis

---

## Why Do We Need Runes?

Go strings are stored as bytes.

For English text, this is usually fine.

Example:

```go
word := "Go"
```

Each character uses 1 byte.

But some characters need multiple bytes.

Example:

```go
word := "你好"
```

Although there are only 2 characters, they use 6 bytes internally.

Runes help us work with the actual characters instead of the raw bytes.

---

## Rune Type

The rune type is:

```go
rune
```

Internally:

```go
rune == int32
```

Go uses `rune` because it is easier to understand than `int32` when working with characters.

---

## Rune Declaration

### Using var

```go
var letter rune = 'A'
```

### Type Inference

```go
var letter = 'A'
```

### Short Variable Declaration

```go
letter := 'A'
```

---

## Single Quotes vs Double Quotes

### Rune

```go
'A'
```

Represents:

```text
One character
```

### String

```go
"A"
```

Represents:

```text
Text/String
```

Example:

```go
char := 'A'
text := "A"
```

---

## Printing Runes

Example:

```go
letter := 'A'

fmt.Println(letter)
```

Output:

```text
65
```

Why?

Because Go prints the Unicode value.

To print the actual character:

```go
fmt.Printf("%c\n", letter)
```

Output:

```text
A
```

---

## Unicode Values

Every character has a Unicode number.

Examples:

| Character | Unicode Value |
| --------- | ------------- |
| A         | 65            |
| B         | 66            |
| ا         | 1575          |
| 你         | 20320         |
| 😀        | 128512        |

Example:

```go
fmt.Println('A')
```

Output:

```text
65
```

---

## Working with Different Languages

Runes support characters from all languages.

Example:

```go
urdu := 'ا'
chinese := '你'
emoji := '😀'

fmt.Printf("%c\n", urdu)
fmt.Printf("%c\n", chinese)
fmt.Printf("%c\n", emoji)
```

Output:

```text
ا
你
😀
```

---

## Checking the Type

```go
letter := 'A'

fmt.Printf("%T\n", letter)
```

Output:

```text
int32
```

Remember:

```text
rune is an alias for int32
```

---

## Iterating Through a String

When looping through a string using `range`, Go returns runes.

Example:

```go
word := "Hello"

for _, ch := range word {
	fmt.Printf("%c ", ch)
}
```

Output:

```text
H e l l o
```

Here `ch` is a rune.

---

## Unicode String Example

```go
word := "你好"

for _, ch := range word {
	fmt.Printf("%c ", ch)
}
```

Output:

```text
你 好
```

Runes allow Go to correctly handle Unicode characters.

---

## Byte Count vs Character Count

Example:

```go
fmt.Println(len("Go"))
```

Output:

```text
2
```

Example:

```go
fmt.Println(len("你好"))
```

Output:

```text
6
```

Because:

```text
你 = 3 bytes
好 = 3 bytes
```

Total:

```text
6 bytes
```

But there are only:

```text
2 characters
```

To count characters:

```go
utf8.RuneCountInString("你好")
```

Output:

```text
2
```

---

## Rune to String Conversion

Example:

```go
letter := 'A'

fmt.Println(string(letter))
```

Output:

```text
A
```

---

## When Should I Use Runes?

Use runes when:

* Working with individual characters
* Processing text character by character
* Supporting multiple languages
* Working with emojis
* Validating characters

Example:

```go
for _, ch := range text {
	// process each character
}
```

---

## Key Points

* A rune represents one Unicode character.
* Runes use single quotes (`'A'`).
* Strings use double quotes (`"A"`).
* `rune` is an alias for `int32`.
* `%c` prints the character.
* `%d` prints the Unicode value.
* `range` over a string returns runes.
* `len()` counts bytes, not characters.
* Use `utf8.RuneCountInString()` to count characters.
* Runes are useful for Unicode text, multiple languages, and emojis.

## Easy Summary

```go
'A'
```

= One character (rune)

```go
"A"
```

= Text (string)

```go
"Hello"
```

= String

```go
'H'
```

= Rune

Use **strings** to store text.

Use **runes** when working with individual characters.
