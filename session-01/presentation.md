# Go 1.21: What's New and Why You Should Care

---

## Go 1.21 Highlights

* Three new built-in functions: min, max, and clear.
* Improved type inference for generic functions.
* Profile Guided Optimization (PGO) feature for faster and more efficient builds.
* New packages for structured logging, slices, maps, and comparison utilities.
* A new port for WebAssembly System Interface (WASI).
* Backward and Forward Language Compatibility

---

## New built-in functions (1)

### Min/Max

Compute the smallest (or largest, for max) value of a fixed number of given arguments

```go
package main

import "fmt"

func main() {
  fmt.Println(min(2, 3), max(2.1, 3.14), min("Hello", "World"))
}
```

---

## New built-in functions (2)

### Why builtin?

So why not using like this?
```go
package main

import "fmt"

func Min[T ~int | ~float32 | ~float64 | ~string](a, b T) T {
  if a < b {
    return a
  }

  return b
}

func main() {
  fmt.Println(Min(2, 3))
}
```

---

## New built-in functions (3)

### Why builtin?

So why not using `math.Min`/`math.Max`
```go
package main

import (
  "fmt"
  "math"
)

func main() {
  inf, nan := math.Inf(1), math.NaN()

  // 1.
  fmt.Println(math.Min(-inf, nan), min(-inf, nan))
  fmt.Println(math.Max(inf, nan), max(inf, nan))

  // 2.
  // fmt.Println(math.Max("Hi", "Hello"))
}
```

---

## New built-in functions (4)

### Clear
Deletes all elements from a map or zeroes all elements of a slice.

```go
package main

import (
  "fmt"
)

func main() {
  m := map[float64]int{
    1.0:  1,
    2.0:  2,
    3.14: 3,
  }

  fmt.Println(m, len(m))
  clear(m)
  fmt.Println(m, len(m))
}
```

---

## New built-in functions (5)

### Clear vs Delete

```go
package main

import (
  "fmt"
  "math"
)

func main() {
  nan := math.NaN()
  m := map[float64]int{
    1.0:  1,
    2.0:  2,
    3.14: 3,
    nan:  4,
  }

  fmt.Println(m, len(m))
  // 1. Delete some elements from the map.
  delete(m, 1.0)
  delete(m, nan)
  fmt.Println(m, len(m))

  // 2. Clear the map.
  clear(m)
  fmt.Println(m, len(m))
}
```

---

## Improved type inference for generic functions. (1)

### Type inference

> A use of a generic function may omit some or all type arguments if
  they can be inferred from the context within which the function is used,
  including the constraints of the function's type parameters.

```go
package main

import "fmt"

func isEven[T ~int](a T) bool {
    return a % 2 == 0
}

func main() {
    // Without type inference
    fmt.Println(isEven[int](2), isEven[int](3))
}
```

---

## Improved type inference for generic functions. (2)

### Type inference

> A use of a generic function may omit some or all type arguments if
  they can be inferred from the context within which the function is used,
  including the constraints of the function's type parameters.

```go
package main

import "fmt"

func isEven[T ~int](a T) bool {
    return a % 2 == 0
}

func main() {
    // With type inference
    fmt.Println(isEven(2), isEven(3))
}
```

---

## Improved type inference for generic functions. (3)

### Partial Type inference

```go
package main

import (
  "golang.org/x/exp/slices"
)

func isEven[T ~int](a T) bool {
  return a%2 == 0
}

func main() {
  s := []int{2, 3, 5}

    // 1. Go 1.20
  slices.IndexFunc(s, isEven[int])

    // 2. Go 1.21
  slices.IndexFunc(s, isEven) // Compile error on Go 1.20.
}
```

---

## Profile Guided Optimization (PGO) 

### PGO
> Profile-guided optimization (PGO), also known as feedback-directed optimization (FDO),
is a compiler optimization technique that feeds information (a profile) from representative
runs of the application back into to the compiler for the next build of the application,
which uses that information to make more informed optimization decisions.

---

## Profile Guided Optimization (PGO) 

### PGO
> Profile-guided optimization (PGO), also known as feedback-directed optimization (FDO),
is a compiler optimization technique that feeds information (a profile) from representative
runs of the application back into to the compiler for the next build of the application,
which uses that information to make more informed optimization decisions.

### What's new in Go 1.21

The -pgo build flag now defaults to -pgo=auto.
> If a file named default.pgo is present in the main package's directory,
the go command will use it to enable profile-guided optimization for building
the corresponding program.

### More on the next session
For now see [this](https://go.dev/doc/pgo)

---

## Structured Logging with slog (1)

### What & Why (1)
* Structured logs use key-value pairs so they can be parsed, filtered, searched, and analyzed quickly and reliably.
* Ranked high in the annual survey.

```go
log.Printf(`{"message": %q, "count": %d}`, msg, count)
```

---

## Structured Logging with slog (2)

### What & Why (2)
* Ease of use
    * pleasant enough that users will prefer it to existing packages in new code.
* High performance
* Integration with runtime tracing

---

## Structured Logging with slog (3)

### What & Why (3)
* Common "backend"
    * Get consistent logging across all its dependencies.
    * Every common logging framework will provide a shim from their own backend to a slog's handler.
    * The Go logging community can work together to build high-quality backends that all can share.

---

## Structured Logging with slog (4)

### Example (1)

```go
package main

import "log/slog"

func main() {
  slog.Info("hello, world", "user", "john")
}

// Sample output:
// 2023/08/23 20:12:22 INFO hello, world user=john
```

---

## Structured Logging with slog (5)

### Example (2)

```go
package main

import (
  "log/slog"
  "os"
)

func main() {
  logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
  slog.SetDefault(logger)

  slog.Info("hello, world", "user", "john")
}

// Sample output:
// {"time":"2023/08/23 20:12:22","level":"INFO","msg":"hello, world","user":"john"}
```

---

## Structured Logging with slog (6)

### More on the next session
* structured logging, slices, maps, and comparison utilities.
* A new port for WebAssembly System Interface (WASI).

---
## Backward Compatibility

Go's emphasis on backwards compatibility is one of its key strengths.

> It is intended that programs written to the Go 1 specification will continue to
compile and run correctly, unchanged, over the lifetime of that specification.

---
## Backward Compatibility

Go's emphasis on backwards compatibility is one of its key strengths.

Cannot maintain strict compatibility:

* Changing sort algorithms 
* Fixing clear bugs
* Existing code depends on the old algorithm or the buggy behavior

---
## Backward Compatibility

Go's emphasis on backwards compatibility is one of its key strengths.

Cannot maintain strict compatibility:

* Changing sort algorithms 
* Fixing clear bugs
* Existing code depends on the old algorithm or the buggy behavior

> Keeping older Go programs executing the same way even when built with newer Go distributions


---
## Backward Compatibility

### GODEBUG
A GODEBUG setting is a key=value pair that controls the execution of certain parts of a Go program.
> GODEBUG=http2client=0,http2server=0

GODEBUG settings added for compatibility will be maintained for a minimum of two years.

---
## Backward Compatibility

### Default GODEBUG Values

GODEBUG settings come from three sources if the environment variable omits them:
1. The defaults for the Go toolchain used to build the program
1. Amended to match the Go version listed in go.mod
1. Then overridden by explicit //go:debug lines in the program

> //go:debug panicnil=1

### GODEBUG History
See https://go.dev/doc/godebug#history

---
## Backward Compatibility

### Example 1 (go1.21, go.mod 1.20)

go.mod:
```go
module example

go 1.20
```

main.go:
```go
package main

import (
  "fmt"
)

func main() {
  defer func() {
    r := recover()
    fmt.Printf("type: %T\nvalue: %v\nis nil: %v\n", r, r, r == nil)
  }()

  panic(nil)
}
```

```bash
─❯ go version
go version go1.21.0 darwin/arm64

─❯ go run .
type: <nil>
value: <nil>
is nil: true
```

---
## Backward Compatibility

### Example 2 (go1.21, go.mod 1.21)

go.mod:
```go
module example

go 1.21
```

main.go:
```go
package main

// same as before
...
```

```bash
─❯ go version
go version go1.21.0 darwin/arm64

─❯ go run .
type: *runtime.PanicNilError
value: panic called with nil argument
is nil: false
```

---
## Backward Compatibility

### Example 3 (go1.21, go.mod 1.21, go:debug ...)

go.mod:
```go
module example

go 1.21
```

main.go:
```go
//go:debug panicnil=1
package main

// same as before
...
```

```bash
─❯ go run .
type: <nil>
value: <nil>
is nil: true

─❯ go list -f '{{.DefaultGODEBUG}}' .
panicnil=1

─❯ go build .

─❯ go version -m ./example | grep GODEBUG
	build	DefaultGODEBUG=panicnil=1
```
