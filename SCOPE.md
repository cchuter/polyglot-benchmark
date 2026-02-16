# Scope: Dominoes Chain Solver (Issue #121)

## In Scope

- Implement `MakeChain` function in `go/exercises/practice/dominoes/dominoes.go`
- Define `type Domino [2]int` in the same file
- Algorithm must handle: empty input, single domino, multi-domino chains, reversal, backtracking, disconnected graph detection
- All 12 test cases in `cases_test.go` and `dominoes_test.go` must pass

## Out of Scope

- Modifying test files (`dominoes_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Performance optimization beyond passing the benchmark
- Adding new test cases
- Changes to any other exercise or language directory

## Dependencies

- Go 1.18+ toolchain
- No external packages (stdlib only)
- Test files are read-only and pre-existing
