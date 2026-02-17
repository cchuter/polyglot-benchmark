# Scope: polyglot-go-counter

## In Scope

- Verifying the existing `counter_test.go` test suite correctly detects all bugs
- Ensuring `counter.go` has the correct package declaration
- Running tests against all four implementations to confirm pass/fail behavior
- Creating `.osmi/` documentation artifacts (GOAL.md, SCOPE.md, plan.md, etc.)
- Committing changes and creating a PR

## Out of Scope

- Modifying `interface.go`, `maker.go`, or any `impl*.go` file
- Changing the `go.mod` file
- Adding new implementation files
- Modifying the `.meta/` or `.docs/` directories
- Performance optimization or benchmarking

## Dependencies

- Go 1.18+ toolchain
- No external packages (standard library only: `testing`, `unicode`, `os`, `log`)
- Exercise infrastructure files: `interface.go`, `maker.go`, `impl1.go`-`impl4.go`
