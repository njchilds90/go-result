# go-result

[![Go Reference](https://pkg.go.dev/badge/github.com/njchilds90/go-result.svg)](https://pkg.go.dev/github.com/njchilds90/go-result)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A small, zero-dependency, generic `Result[T]` type for explicit and composable error handling in Go.

Inspired by Rust’s `Result`, this package provides a type-safe alternative to the traditional `(T, error)` return pattern when you want cleaner chaining and functional-style composition.

---

## ✨ Features

- ✅ Generic and fully type-safe (`Result[T]`)
- 📦 Zero dependencies
- 🔗 Chainable operations (`AndThen`, `Map`)
- ⚠️ Explicit success/error state
- 🧪 Easy to test and reason about
- 🤖 Deterministic behavior (AI-friendly and pipeline-friendly)

---

## 📦 Installation

```bash
go get github.com/njchilds90/go-result
```

---

## 🧠 Why Use Result?

Go’s standard pattern:

```go
value, err := doSomething()
if err != nil {
    return err
}
```

This is explicit and idiomatic — but can become repetitive in multi-step flows.

`Result[T]` lets you compose operations cleanly:

```go
final := result.Ok(10).
    AndThen(double).
    AndThen(double)
```

Errors automatically propagate without nested conditionals.

---

## 🚀 Quick Example

```go
package main

import (
	"errors"
	"fmt"

	"github.com/njchilds90/go-result"
)

func double(x int) result.Result[int] {
	return result.Ok(x * 2)
}

func fail(x int) result.Result[int] {
	return result.Err[int](errors.New("failed"))
}

func main() {
	r := result.Ok(5).
		AndThen(double).
		AndThen(double)

	if r.IsOk() {
		fmt.Println("Success:", r.Value()) // 20
	}

	r2 := result.Ok(5).
		AndThen(double).
		AndThen(fail)

	if r2.IsErr() {
		fmt.Println("Error:", r2.Error())
	}
}
```

---

## 📘 API Overview

### Creating Results

```go
result.Ok[T](value T)
result.Err[T](err error)
```

- `Err` panics if `err == nil`
- `Ok` always produces a success value

---

### Inspecting Results

```go
r.IsOk() bool
r.IsErr() bool
r.Value() T
r.Error() error
r.Unwrap() T
r.UnwrapOr(default T) T
```

- `Value()` and `Unwrap()` panic if called on an error result  
- `Error()` panics if called on a success result  

This is intentional — misuse should fail fast.

---

### Composing Results

```go
r.Map(func(T) T) Result[T]
r.AndThen(func(T) Result[T]) Result[T]
```

- `Map` transforms a successful value
- `AndThen` chains operations that may fail
- Errors propagate automatically

---

## 🧪 Example Test

```go
func TestChain(t *testing.T) {
	double := func(x int) result.Result[int] {
		return result.Ok(x * 2)
	}

	r := result.Ok(5).AndThen(double).AndThen(double)

	if !r.IsOk() || r.Value() != 20 {
		t.Fatalf("unexpected result: %v", r)
	}
}
```

---

## 📌 Design Philosophy

- Keep it small and predictable
- No hidden allocations or reflection
- No opinionated logging or wrapping
- Panic only on misuse (not on normal error flow)
- Preserve Go’s explicitness while enabling composition

---

## 🤝 When To Use (And When Not To)

### Good Fit

- Multi-step pipelines
- Functional-style transformations
- Domain/business logic composition
- Internal application layers

### Prefer `(T, error)` When

- Building public Go APIs meant to feel idiomatic
- Interoperating heavily with existing Go libraries
- Simplicity matters more than composition

---

## 🛠 Suggested Future Enhancements

- `MapErr(func(error) error)`
- `Fold(onOk, onErr)`
- `OrElse(func(error) Result[T])`
- `Must()` helper for tests
- `String()` method for debugging
- Optional `Option[T]` companion package

---

## 📄 License

MIT License — see LICENSE file.

---

## 🙌 Contributions

Issues and pull requests are welcome.  
Keep the API minimal, composable, and idiomatic.
