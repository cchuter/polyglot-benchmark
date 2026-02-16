# Implementation Plan: paasio

## Overview

Implement IO statistics wrappers in `go/exercises/practice/paasio/paasio.go`. The solution needs three types and three constructor functions, all with thread-safe counting using `sync.Mutex`.

## File to Modify

- `go/exercises/practice/paasio/paasio.go` — the only file to change

## Architecture

### Shared counter type

A `counter` struct holds `bytes int64`, `ops int`, and a `*sync.Mutex`. Two methods:
- `addBytes(n int)` — locks, increments bytes by n and ops by 1, unlocks
- `count() (int64, int)` — locks, returns bytes and ops, unlocks

This ensures `ReadCount`/`WriteCount` return atomically consistent snapshots.

### readCounter struct

```go
type readCounter struct {
    r io.Reader
    counter
}
```

- `Read(p []byte) (int, error)` — delegates to `r.Read(p)`, then calls `addBytes(n)` with the returned byte count
- `ReadCount() (int64, int)` — delegates to `counter.count()`

### writeCounter struct

```go
type writeCounter struct {
    w io.Writer
    counter
}
```

- `Write(p []byte) (int, error)` — delegates to `w.Write(p)`, then calls `addBytes(n)` with the returned byte count
- `WriteCount() (int64, int)` — delegates to `counter.count()`

### rwCounter struct

```go
type rwCounter struct {
    WriteCounter
    ReadCounter
}
```

Embeds both a `WriteCounter` and a `ReadCounter`, satisfying the `ReadWriteCounter` interface through composition.

### Constructor functions

1. `NewReadCounter(r io.Reader) ReadCounter` — returns `&readCounter{r: r, counter: newCounter()}`
2. `NewWriteCounter(w io.Writer) WriteCounter` — returns `&writeCounter{w: w, counter: newCounter()}`
3. `NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter` — returns `&rwCounter{NewWriteCounter(rw), NewReadCounter(rw)}`

### Helper

`newCounter() counter` — returns `counter{mutex: new(sync.Mutex)}`

## Ordering

1. Write the complete implementation in `paasio.go`
2. Run `go vet ./...` in the exercise directory
3. Run `go test ./...` in the exercise directory
4. Fix any issues if needed

## Rationale

- Using `sync.Mutex` rather than `sync/atomic` because the tests require atomically consistent snapshots of both bytes AND ops together (a single lock covers both fields)
- Embedding `counter` in read/write counter types avoids code duplication
- Composing `rwCounter` from `WriteCounter` + `ReadCounter` interfaces means read and write have independent locks, which is correct per the interface contract
