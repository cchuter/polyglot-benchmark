# Scope: polyglot-go-dominoes (Issue #335)

## In Scope

- Implement `Domino` type and `MakeChain` function in `go/exercises/practice/dominoes/dominoes.go`
- Handle all edge cases: empty input, single domino, disconnected graphs, backtracking
- Ensure all 11 test cases pass
- Ensure `go vet` passes

## Out of Scope

- Modifying test files (`dominoes_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Performance optimization beyond passing the benchmark
- Adding additional test cases
- Changes to any other exercise or language directory

## Dependencies

- Go 1.18+ (as specified in `go.mod`)
- No external packages required (stdlib only)
- Test infrastructure already exists in `dominoes_test.go` and `cases_test.go`
