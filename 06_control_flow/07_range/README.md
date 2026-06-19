# 07_range

## What is range?

`range` is Go's foreach-style loop.

It is used to go through each item in a collection one by one.

Example:

```go
names := []string{
	"Ali",
	"Ahmed",
	"Sara",
}

for index, value := range names {
	fmt.Println(index, value)
}
```

Output:

```text
0 Ali
1 Ahmed
2 Sara
```

---

# Why Use range?

Without `range`:

```go
for i := 0; i < len(names); i++ {
	fmt.Println(i, names[i])
}
```

With `range`:

```go
for index, value := range names {
	fmt.Println(index, value)
}
```

The `range` version is shorter and easier to read.

---

# Basic Syntax

```go
for index, value := range collection {
	// code
}
```

Think of it as:

```text
For each item in the collection:
Give me its position (index)
Give me its value
```

---

# Understanding Each Part

```go
for index, value := range names {
	fmt.Println(index, value)
}
```

### for

Starts a loop.

```go
for
```

means:

```text
Repeat code multiple times
```

---

### range

```go
range names
```

means:

```text
Go through each item in names
one by one
```

---

### index

```go
index
```

stores the position of the item.

Example:

```text
Ali   -> index 0
Ahmed -> index 1
Sara  -> index 2
```

---

### value

```go
value
```

stores the actual item.

Example:

```text
Ali
Ahmed
Sara
```

---

### :=

```go
index, value :=
```

creates the variables:

```go
index
value
```

for the loop.

---

# Step-by-Step Example

```go
names := []string{
	"Ali",
	"Ahmed",
	"Sara",
}
```

Visual:

```text
Index    Value
-----    -----
0        Ali
1        Ahmed
2        Sara
```

Loop:

```go
for index, value := range names {
	fmt.Println(index, value)
}
```

---

### First Iteration

```go
index = 0
value = "Ali"
```

Output:

```text
0 Ali
```

---

### Second Iteration

```go
index = 1
value = "Ahmed"
```

Output:

```text
1 Ahmed
```

---

### Third Iteration

```go
index = 2
value = "Sara"
```

Output:

```text
2 Sara
```

---

# Ignore the Index

Sometimes you only need the value.

```go
for _, value := range names {
	fmt.Println(value)
}
```

Output:

```text
Ali
Ahmed
Sara
```

The underscore:

```go
_
```

means:

```text
I don't need this value
Ignore it
```

---

# Only Use the Index

Sometimes you only need the position.

```go
for index := range names {
	fmt.Println(index)
}
```

Output:

```text
0
1
2
```

---

# What Does range Actually Return?

For slices and arrays:

```go
index, value
```

Example:

```text
0 Ali
1 Ahmed
2 Sara
```

The first variable gets the index.

The second variable gets the value.

---

# Easy Comparison

Traditional loop:

```go
for i := 0; i < len(names); i++ {
	fmt.Println(i, names[i])
}
```

Range loop:

```go
for index, value := range names {
	fmt.Println(index, value)
}
```

Both produce:

```text
0 Ali
1 Ahmed
2 Sara
```

The range version is easier to write.

---

# Using Existing Variables

Normally:

```go
for index, value := range names {
}
```

creates variables.

You can also use existing variables.

```go
var index int
var value string

for index, value = range names {
	fmt.Println(index, value)
}
```

Here:

```go
=
```

updates existing variables.

---

# Most Common Pattern

This is the pattern you will use most often:

```go
for _, value := range names {
	fmt.Println(value)
}
```

Because most of the time you care about the item, not its index.

---

# Common Beginner Mistakes

### Forgetting _

Wrong:

```go
for index, value := range names {
	fmt.Println(value)
}
```

Error:

```text
index declared and not used
```

Correct:

```go
for _, value := range names {
	fmt.Println(value)
}
```

---

### Confusing index and value

Wrong idea:

```text
index = Ali
value = 0
```

Correct:

```text
index = 0
value = Ali
```

---

### Confusing := and =

```go
for index, value := range names {
}
```

creates variables.

```go
for index, value = range names {
}
```

uses existing variables.

---

# Easy Memory Trick

```go
for index, value := range names {
}
```

Read it as:

```text
For each item in names,
give me:

1. Its position (index)
2. Its actual value (value)
```

Example:

```text
names
│
├── [0] Ali
├── [1] Ahmed
└── [2] Sara
```

Loop gives:

```text
index=0 value=Ali
index=1 value=Ahmed
index=2 value=Sara
```

---

# Key Points

* `range` is Go's foreach loop.
* It loops through a collection one item at a time.
* `index` stores the position.
* `value` stores the actual item.
* `_` ignores unused values.
* `:=` creates loop variables.
* `=` uses existing variables.
* Most common pattern:

```go
for _, value := range items {
	fmt.Println(value)
}
```

* Use `range` when you want to process each item in a collection.
