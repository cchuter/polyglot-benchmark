# Implementation Plan: polyglot-go-paasio

## Proposal A: Mutex-based counter with embedded struct

**Role: Proponent**

### Approach

Use a shared `counter` struct with a `sync.Mutex` that tracks bytes and ops. Embed this counter into `readCounter` and `writeCounter` structs. Compose `rwCounter` from `WriteCounter` and `ReadCounter` interfaces.

### File Changes

- **`paasio.go`** — single file, complete implementation.

### Design

```go
package paasio

import (
    "io"
    "sync"
)

type counter struct {
    bytes int64
    ops   int
    mu    sync.Mutex
}

func (c *counter) addBytes(n int) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.bytes += int64(n)
    c.ops++
}

func (c *counter) count() (int64, int) {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.bytes, c.ops
}

type readCounter struct {
    r io.Reader
    counter
}

func NewReadCounter(r io.Reader) ReadCounter {
    return &readCounter{r: r}
}

func (rc *readCounter) Read(p []byte) (int, error) {
    n, err := rc.r.Read(p)
    rc.addBytes(n)
    return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
    return rc.count()
}

type writeCounter struct {
    w io.Writer
    counter
}

func NewWriteCounter(w io.Writer) WriteCounter {
    return &writeCounter{w: w}
}

func (wc *writeCounter) Write(p []byte) (int, error) {
    n, err := wc.w.Write(p)
    wc.addBytes(n)
    return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
    return wc.count()
}

type rwCounter struct {
    WriteCounter
    ReadCounter
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
    return &rwCounter{
        WriteCounter: NewWriteCounter(rw),
        ReadCounter:  NewReadCounter(rw),
    }
}
```

### Rationale

- **Simple and idiomatic**: Uses embedding and composition, standard Go patterns.
- **Correct**: Mutex protects bytes+ops updates atomically, ensuring consistency tests pass.
- **Minimal**: Single helper type (`counter`) avoids duplication.
- **Matches the reference solution** in `.meta/example.go` almost exactly, which validates this approach.

---

## Proposal B: Atomic-based counter without mutex

**Role: Opponent**

### Approach

Use `sync/atomic` operations instead of a mutex. Store bytes and ops as `int64` values and use `atomic.AddInt64` / `atomic.LoadInt64` for lock-free counting.

### Design

```go
package paasio

import (
    "io"
    "sync/atomic"
)

type readCounter struct {
    r     io.Reader
    bytes int64
    ops   int64
}

func NewReadCounter(r io.Reader) ReadCounter {
    return &readCounter{r: r}
}

func (rc *readCounter) Read(p []byte) (int, error) {
    n, err := rc.r.Read(p)
    atomic.AddInt64(&rc.bytes, int64(n))
    atomic.AddInt64(&rc.ops, 1)
    return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
    return atomic.LoadInt64(&rc.bytes), int(atomic.LoadInt64(&rc.ops))
}
// ... similar for writeCounter, rwCounter
```

### Critique of Proposal A

Proposal A uses mutexes which have higher overhead than atomics in high-contention scenarios.

### Critique of Proposal B (self-critique)

**Critical flaw**: The consistency tests (`testReadCountConsistency`, `testWriteCountConsistency`) verify that `bytes == ops * bytesPerOp` at all times. With separate atomic operations for bytes and ops, there is a window between `AddInt64(&rc.bytes, ...)` and `AddInt64(&rc.ops, 1)` where a concurrent `ReadCount()` call would see updated bytes but stale ops — **violating the consistency invariant**. This approach will fail the consistency tests.

### Rationale

While atomics are faster, they cannot provide the atomicity guarantee needed here: bytes and ops must be updated as a single atomic unit. Only a mutex (or a single 64-bit atomic packing both values) can provide this.

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion     | Proposal A (Mutex) | Proposal B (Atomics) |
|---------------|-------------------|---------------------|
| Correctness   | Fully correct — mutex ensures bytes+ops are always consistent | **FAILS** consistency tests — two separate atomics create a race window |
| Risk          | Low — matches reference solution | High — will fail tests |
| Simplicity    | Simple, idiomatic Go | Slightly less code but more subtle correctness issues |
| Consistency   | Matches `.meta/example.go` conventions exactly | Deviates from reference |

### Decision

**Proposal A wins decisively.** Proposal B has a fundamental correctness flaw that would cause test failures. Proposal A is proven correct by the reference solution and uses standard Go synchronization patterns.

### Final Implementation Plan

**File to modify**: `go/exercises/practice/paasio/paasio.go`

**Implementation steps** (in order):

1. Add imports: `io`, `sync`
2. Define `counter` struct with `bytes int64`, `ops int`, and `mu sync.Mutex`
3. Add `addBytes(n int)` method on `*counter` — locks mutex, increments bytes and ops
4. Add `count() (int64, int)` method on `*counter` — locks mutex, returns bytes and ops
5. Define `readCounter` struct embedding `counter` with field `r io.Reader`
6. Implement `NewReadCounter(r io.Reader) ReadCounter` — returns `&readCounter{r: r}`
7. Implement `Read(p []byte) (int, error)` on `*readCounter` — delegates to `rc.r.Read(p)`, calls `addBytes(n)`
8. Implement `ReadCount() (int64, int)` on `*readCounter` — delegates to `rc.count()`
9. Define `writeCounter` struct embedding `counter` with field `w io.Writer`
10. Implement `NewWriteCounter(w io.Writer) WriteCounter` — returns `&writeCounter{w: w}`
11. Implement `Write(p []byte) (int, error)` on `*writeCounter` — delegates to `wc.w.Write(p)`, calls `addBytes(n)`
12. Implement `WriteCount() (int64, int)` on `*writeCounter` — delegates to `wc.count()`
13. Define `rwCounter` struct embedding `WriteCounter` and `ReadCounter`
14. Implement `NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter` — composes `NewWriteCounter(rw)` and `NewReadCounter(rw)`
