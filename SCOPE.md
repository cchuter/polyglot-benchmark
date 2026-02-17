# Scope: polyglot-go-pig-latin

## In Scope

- Implement the `Sentence` function in `go/exercises/practice/pig-latin/pig_latin.go`
- Handle all four Pig Latin translation rules
- Handle multi-word phrases (split on spaces, translate each word, rejoin)
- Pass all 22 test cases defined in `cases_test.go`

## Out of Scope

- Modifying test files (`pig_latin_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Handling uppercase letters (all test inputs are lowercase)
- Handling punctuation or special characters
- Performance optimization beyond passing benchmarks
- Changes to any other exercise directories

## Dependencies

- Go standard library only (`strings` package)
- No external packages required
- Go 1.18+ (as specified in go.mod)
