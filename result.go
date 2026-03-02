// Package result provides a generic Result[T] type for explicit and composable success/error handling.
// Inspired by Rust’s Result and functional patterns, but fully idiomatic Go.
package result

import (
	"fmt"
)

// Result represents either a success value of type T or an error.
// It is immutable and has no hidden state.
type Result[T any] struct {
	value T
	err   error
}

// Ok creates a successful Result containing the given value.
func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

// Err creates a failed Result containing the given error.
// If err is nil, it panics (to prevent accidental Ok(nil error) mistakes).
func Err[T any](err error) Result[T] {
	if err == nil {
		panic("result: Err called with nil error")
	}
	return Result[T]{err: err}
}

// IsOk returns true if the Result contains a success value.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr returns true if the Result contains an error.
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Value returns the success value. It panics if the Result is an error.
func (r Result[T]) Value() T {
	if r.err != nil {
		panic(fmt.Sprintf("result: called Value on Err result: %v", r.err))
	}
	return r.value
}

// Error returns the error. It panics if the Result is successful.
func (r Result[T]) Error() error {
	if r.err == nil {
		panic("result: called Error on Ok result")
	}
	return r.err
}

// Unwrap returns the success value or panics with the error (Rust-style).
func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}

// UnwrapOr returns the success value or the provided default if the Result is an error.
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.err != nil {
		return defaultValue
	}
	return r.value
}

// AndThen calls the provided function only if the Result is Ok and returns a new Result.
// This enables clean chaining of operations that can fail.
func (r Result[T]) AndThen(f func(T) Result[T]) Result[T] {
	if r.err != nil {
		return r
	}
	return f(r.value)
}

// Map applies a pure transformation to the success value (if present).
// Errors are passed through unchanged.
func (r Result[T]) Map(f func(T) T) Result[T] {
	if r.err != nil {
		return r
	}
	return Ok(f(r.value))
}
