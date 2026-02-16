# Scope: Pig Latin Exercise (Issue #145)

## In Scope

- Implement the `Sentence` function in `go/exercises/practice/pig-latin/pig_latin.go`
- Optionally implement a helper `Word` function for single-word translation
- Handle all four pig latin translation rules as defined in the issue
- Handle multi-word input by splitting, translating each word, and rejoining
- Pass all existing test cases

## Out of Scope

- Modifying test files (`cases_test.go`, `pig_latin_test.go`)
- Modifying `go.mod`
- Adding new test cases
- Unicode/non-ASCII character handling (tests only use lowercase ASCII)
- Case preservation (tests use lowercase input/output only)
- Punctuation handling (tests don't include punctuation)
- Any changes to other exercises or other language directories

## Dependencies

- Go toolchain (1.18+)
- Standard library only (`regexp`, `strings`)
- No external packages required
