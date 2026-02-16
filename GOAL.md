# Goal: polyglot-go-error-handling

## Problem Statement

Implement the `Use` function in `go/exercises/practice/error-handling/error_handling.go` that demonstrates Go error handling and resource management patterns including defer, panic, and recover.

## What Needs to Be Built

A function `Use(opener ResourceOpener, input string) error` that:

1. Opens a resource using the provided `ResourceOpener` function
2. Calls `Frob(input)` on the opened resource
3. Closes the resource in all cases (success, error, panic)
4. Properly handles errors and panics

## Acceptance Criteria

1. **Happy path**: When opener succeeds and Frob doesn't panic, return nil and ensure Close is called exactly once
2. **Transient errors**: When opener returns a `TransientError`, keep retrying until it succeeds or returns a non-transient error
3. **Non-transient open errors**: When opener returns a non-`TransientError` error, return that error immediately
4. **FrobError panic**: When `Frob` panics with a `FrobError`, call `Defrob(frobError.defrobTag)` on the resource, then Close, and return the error
5. **Non-FrobError panic**: When `Frob` panics with a non-`FrobError` error, call Close (but not Defrob) and return the error
6. **Close exactly once**: Resource's `Close()` must be called exactly once whenever a resource was successfully opened
7. All tests in `error_handling_test.go` pass via `go test ./...`
8. Code passes `go vet ./...`

## Key Constraints

- Solution must be in package `erratum`
- Solution file is `error_handling.go`
- Must use Go's `defer`, `recover`, and named return values
- Types `Resource`, `ResourceOpener`, `FrobError`, and `TransientError` are defined in `common.go`
