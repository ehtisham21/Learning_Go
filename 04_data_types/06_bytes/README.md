# 06 - Bytes in Go

## What is a Byte?

A byte is the basic unit used to store data in memory.

In Go:

```go
byte == uint8
```

This means:

```go
var a byte = 65
var b uint8 = 65
```

Both are exactly the same.

Go provides the name `byte` because it is easier to understand when working with raw data.

---

## Why Do We Need Bytes?

Computers store everything as bytes.

Examples:

* Text
* Files
* Images
* Videos
* JSON Data
* HTTP Requests
* Network Data

Everything is eventually stored as bytes.

---

## Byte Declaration

### Using var

```go
var b byte = 65
```

### Short Declaration

```go
b := byte(65)
```

---

## Printing a Byte

Example:

```go
b := byte(65)

fmt.Println(b)
```

Output:

```text
65
```

A byte stores a numeric value.

---

## Printing a Byte as a Character

Use `%c`.

Example:

```go
b := byte(65)

fmt.Printf("%c\n", b)
```

Output:

```text
A
```

Because:

```text
A = 65
```

---

## Checking the Type

```go
b := byte(65)

fmt.Printf("%T\n", b)
```

Output:

```text
uint8
```

Remember:

```text
byte is an alias for uint8
```

---

## Strings Are Stored as Bytes

Example:

```go
text := "Go"
```

Internally:

```text
G = 71
o = 111
```

---

## Convert String to Bytes

Example:

```go
text := "Go"

data := []byte(text)

fmt.Println(data)
```

Output:

```text
[71 111]
```

Explanation:

```text
G = 71
o = 111
```

---

## Convert Bytes Back to String

Example:

```go
data := []byte{71, 111}

fmt.Println(string(data))
```

Output:

```text
Go
```

---

## Accessing Bytes in a String

Example:

```go
word := "Go"

fmt.Println(word[0])
fmt.Println(word[1])
```

Output:

```text
71
111
```

These are byte values.

To print them as characters:

```go
fmt.Printf("%c\n", word[0])
fmt.Printf("%c\n", word[1])
```

Output:

```text
G
o
```

---

## Byte Count

Use `len()`.

Example:

```go
word := "Go"

fmt.Println(len(word))
```

Output:

```text
2
```

Because the string uses 2 bytes.

---

## Unicode Example

Example:

```go
word := "你"
```

Byte count:

```go
fmt.Println(len(word))
```

Output:

```text
3
```

Because:

```text
你 uses 3 bytes in UTF-8
```

---

## Actual Bytes

To see the actual bytes:

```go
fmt.Println([]byte(word))
```

Output:

```text
[228 189 160]
```

Explanation:

```text
len(word)
```

tells us:

```text
How many bytes?
```

Output:

```text
3
```

---

```go
[]byte(word)
```

tells us:

```text
What are those bytes?
```

Output:

```text
[228 189 160]
```

---

## Byte Slice

A collection of bytes is called a byte slice.

Type:

```go
[]byte
```

Example:

```go
data := []byte("Hello")

fmt.Println(data)
```

Output:

```text
[72 101 108 108 111]
```

---

## Why Is []byte Important?

Many Go libraries use:

```go
[]byte
```

Examples:

### File Handling

```go
data, err := os.ReadFile("file.txt")
```

Returns:

```go
[]byte
```

---

### JSON Processing

```go
json.Marshal()
json.Unmarshal()
```

Often work with:

```go
[]byte
```

---

### HTTP Requests

Request and response bodies are commonly handled as bytes.

---

### Hashing

```go
sha256.Sum256([]byte(password))
```

---

### Encryption

Most encryption libraries require:

```go
[]byte
```

---

## Byte vs Rune vs String

### String

Stores text.

Example:

```go
"Hello"
```

---

### Rune

Represents one character.

Example:

```go
'H'
```

---

### Byte

Represents raw storage data.

Example:

```go
72
```

---

## Example

```go
text := "Go"
```

String:

```text
"Go"
```

Runes:

```text
'G'
'o'
```

Bytes:

```text
71
111
```

---

## Key Points

* `byte` is an alias for `uint8`.
* A byte stores values from 0 to 255.
* Strings are internally stored as bytes.
* `[]byte` means a slice of bytes.
* `len()` returns the number of bytes.
* `[]byte(text)` returns the actual bytes.
* `word[i]` returns a byte.
* `%c` prints a byte as a character.
* File handling, networking, JSON, hashing, and encryption commonly use bytes.

## Easy Summary

```text
string = text

rune   = one character

byte   = raw storage unit
```

Example:

```go
text := "Go"
```

```text
string -> "Go"

runes  -> 'G', 'o'

bytes  -> 71, 111
```

Use:

* `string` to store text
* `rune` to work with characters
* `byte` to work with raw data
