## Basics

### Packages

- all Go programs are comprised of packages.
- packages start running in the `main` package
- conventionally, package names are the same as the last element of the import path
  - e.g. `math/rand` package comprises files beginning with `package rand` statement

### Imports

- imports can be grouped ("factored" import statement—good practice/style) or individually line-by-line

```go
package main

// valid, okay practice
import "fmt"
import "math"

// better, more common approach
import (
  "fmt"
  "math"
)
```

### Exported Names

- names<sup>1</sup> are only exported if it begins with a capital letter
  - think public attributes
  - e.g.: `Pi` via the `math` package
- can only refer to exported names from imported packages; attempting to use an "unexported" name will result in an error
  - think of "unexported" as private

```go
package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println(math.Pi) // 3.141592653589793
}
```

<sup>1</sup> I assume this refers to any:

- function
- type
- interface
- etc.

### Functions

- take zero or more arguments
- args go in parenthesis, with their types stated _**after**_ the arg name
- return type should also be stated following argument declaration

```go
func add(x int, y int) int {
  return x + y
}
```

- if two or more consecutive named function params/args share a type, all but the last types may be omitted

```go
func add(x, y int) int {
  return x + y
}
```

### Multiple Results

- functions can return any number of results, all of which need to be typed accordingly

```go
func swap(x, y string) (string, string) {
  return y, x
}
```

### Named Results

- return values may be named; if so, treated as regular variables defined in function declaration line
- named returns are used to document the meaning of said values
- `return` without any arguments returns the named return values
  - this is called a **_"naked" return_**
- naked returns should only be used in short functions; can harm readability in longer functions

```go
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y - sum - x
  return // compiled? to `return x, y`
}
```

### Variables

- use `var` statements to declare variables; similar to function args, the type is last
- `var` statement can be at both the package or function level

```go
package main

import "fmt"

// package level
var c, python, java bool

func main() {
  // function level
  var i int
  fmt.Println(i, c, python, java)
}
```

### Initialized Variables

- `var` declaration can include initializers, one per variable
- if initializer is present, the type can be omitted; variable will infer the type based on the initializer

```go
package main

import "fmt"

var i, j int = 1, 2
var k = 3.14

func main() {
  var c, python, java = true, false, "no!"
  fmt.Println(i, j, k, c, python, java)                     // 1 2 3.14 true false no!
  fmt.Printf("%T %T %T %T %T %T", i, j, k, c, python, java) // int int bool bool string
}
```

### Short Variable Declarations

- inside functions, the `:=` short assignment operator can be used in place of a `var` declaration w/ implicit type
  - not available outside functions because "globally" every statement begins with a keyword (`package`, `import`, `func`, `var`, etc.)

```go
func main() {
  var i, j int = 1, 2
  k := 3.14

  fmt.Println(i, j, k)
}
```

### Basic Types

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

- `int` should be used when needing an integer value unless there's a specific reason to use a sized or unsigned integer type

### Zero Values

- variables declared w/o an explicit initial value are given their _zero value_
- zero value refers to:
  - `0` for numeric types
  - `false` for booleans
  - `""` (empty string) for strings

### Type Conversions

- `T(v)` expression denotes converting value `v` to type `T`
- example numeric conversions:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// alternatively…
i := 42
f := float64(i)
u := uint(f)
```

### Constants

- constants are declared like variables using the `const` keyword
- constants can be character, string, boolean, or numeric values
- constants cannot be used with the `:=` shorthand syntax
