# 05 - Imports

## Introduction

Imports allow a Go program to use code from other packages.

A package contains related functions, variables, constants, and types. By importing a package, we can access and use its functionality in our program.

Without imports, we would only be able to use the code written in the current file or package.

---

# Why Do We Need Imports?

Suppose we want to print something on the screen:

```go
fmt.Println("Hello Go")
```

The `Println()` function belongs to the `fmt` package.

To use it, we must import the package first:

```go
import "fmt"
```

Complete example:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go")
}
```

---

# Import Syntax

## Single Import

When importing only one package:

```go
import "fmt"
```

Example:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```

---

## Multiple Imports

When importing multiple packages:

```go
import (
	"fmt"
	"strings"
	"time"
)
```

Example:

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("golang"))
}
```

Output:

```text
GOLANG
```

---

# Common Standard Library Packages

Go provides a rich Standard Library.

## fmt

Used for formatted input and output.

Import:

```go
import "fmt"
```

Common functions:

```go
fmt.Println()
fmt.Printf()
fmt.Sprintf()
```

Example:

```go
fmt.Println("Hello Go")
```

---

## strings

Used for string manipulation.

Import:

```go
import "strings"
```

Example:

```go
strings.ToUpper("golang")
strings.ToLower("GOLANG")
strings.Contains("golang", "go")
```

Output:

```text
GOLANG
golang
true
```

---

## math

Used for mathematical operations.

Import:

```go
import "math"
```

Example:

```go
math.Sqrt(64)
math.Pow(2, 3)
```

Output:

```text
8
8
```

---

## time

Used for working with dates and time.

Import:

```go
import "time"
```

Example:

```go
time.Now()
```

Example output:

```text
2026-06-17 12:00:00 +0000 UTC
```

---

# Using Imported Packages

Functions from imported packages are accessed using:

```go
packageName.FunctionName()
```

Examples:

```go
fmt.Println()
strings.ToUpper()
math.Sqrt()
time.Now()
```

Pattern:

```text
package.function
```

---

# Complete Example

```go
package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {

	fmt.Println("=== fmt Package ===")

	fmt.Println("Hello from Go Imports")

	fmt.Println("\n=== strings Package ===")

	fmt.Println(strings.ToUpper("golang"))
	fmt.Println(strings.ToLower("GOLANG"))
	fmt.Println(strings.Contains("golang", "go"))

	fmt.Println("\n=== math Package ===")

	fmt.Println("Square Root of 64:", math.Sqrt(64))
	fmt.Println("2 Power 3:", math.Pow(2, 3))

	fmt.Println("\n=== time Package ===")

	fmt.Println("Current Time:", time.Now())

	name := "ehtisham"

	fmt.Println("\n=== Package Functions ===")

	fmt.Println("Original Name:", name)
	fmt.Println("Uppercase Name:", strings.ToUpper(name))

	fmt.Println("Square Root:", math.Sqrt(25))

	fmt.Println("Current Year:", time.Now().Year())
}
```

---

# Importing Your Own Packages

Project Structure:

```text
go-learning/
│
├── go.mod
├── main.go
│
└── mathutils/
    └── add.go
```

### add.go

```go
package mathutils

func Add(a int, b int) int {
	return a + b
}
```

### main.go

```go
package main

import (
	"fmt"
	"github.com/ehtisham21/go-learning/mathutils"
)

func main() {
	fmt.Println(mathutils.Add(5, 3))
}
```

Output:

```text
8
```

---

# Import Alias

Sometimes package names are long.

Example:

```go
import mu "github.com/ehtisham21/go-learning/mathutils"
```

Use:

```go
fmt.Println(mu.Add(5, 3))
```

Here:

```text
mathutils → mu
```

---

# Blank Import

A blank import uses `_`.

Example:

```go
import _ "github.com/lib/pq"
```

This imports a package only for its initialization code.

You will commonly see this when working with database drivers.

---

# Dot Import

Example:

```go
import . "fmt"
```

Then:

```go
Println("Hello")
```

instead of:

```go
fmt.Println("Hello")
```

Although valid, dot imports are generally discouraged because they reduce code readability.

---

# Unused Imports

Go does not allow unused imports.

Example:

```go
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello")
}
```

Error:

```text
"strings" imported and not used
```

To fix this:

* Use the package.
* Or remove the import.

---

# Import Rules

## Rule 1

Package declaration comes first.

```go
package main
```

---

## Rule 2

Imports come immediately after package declaration.

```go
package main

import "fmt"
```

---

## Rule 3

Imports must be used.

Unused imports cause compilation errors.

---

## Rule 4

Use grouped imports when importing multiple packages.

```go
import (
	"fmt"
	"strings"
	"time"
)
```

---

# Key Takeaways

* Imports allow a program to use functionality from other packages.
* Every imported package must be used.
* Standard library packages provide powerful built-in functionality.
* Functions are accessed using:

```go
packageName.FunctionName()
```

* Custom packages can be imported using the module path.
* Import aliases can shorten long package names.
* Blank imports are used for initialization-only packages.
* Dot imports should generally be avoided.
* Go enforces clean imports by disallowing unused packages.

---

# Summary

In this section, you learned:

* What imports are.
* Why imports are required.
* Single and multiple import syntax.
* Common standard library packages.
* How to use package functions.
* How to import your own packages.
* Import aliases.
* Blank imports.
* Dot imports.
* Unused import errors.

Imports are one of the core building blocks of Go and are used in almost every Go program.
