# go-result

[![Go Reference](https://pkg.go.dev/badge/github.com/njchilds90/go-result.svg)](https://pkg.go.dev/github.com/njchilds90/go-result)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Generic Result[T] type for explicit, composable success/error handling in Go.

## Features

- Fully generic and type-safe
- Zero external dependencies
- Pure value type (no hidden state)
- Chainable methods (`AndThen`, `Map`)
- Clear panic-on-misuse behavior for safety
- Structured, deterministic, and AI-agent friendly

## Installation

```bash
go get github.com/njchilds90/go-result
Usage
Gopackage main

import (
	"fmt"

	"github.com/njchilds90/go-result"
)

func main() {
	// Simple success path
	r := result.Ok(42)
	fmt.Println("Value:", r.Value()) // 42

	// Chaining operations that can fail
	double := func(x int) result.Result[int] {
		return result.Ok(x * 2)
	}

	final := result.Ok(10).
		AndThen(double).
		AndThen(double)

	if final.IsOk() {
		fmt.Println("Final:", final.Value()) // 40
	}
}
See the GoDoc for every exported function, including full examples.
License
MIT
