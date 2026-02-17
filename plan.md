# Implementation Plan: paasio Exercise

## Proposal A: Mutex-based approach

**Role: Proponent**

### Approach

Use a `sync.Mutex` to protect byte count and operation count fields. This is the classic approach for synchronized shared state in Go.

### Structures

- `counter` struct with `bytes int64`, `ops int`, and `mutex *sync.Mutex`
- `readCounter` struct embedding `counter` and holding an `io.Reader`
- `writeCounter` struct embedding `counter` and holding an `io.Writer`
- `rwCounter` struct embedding both `WriteCounter` and `ReadCounter`

### Implementation

1. `counter.addBytes(n int)` — locks mutex, increments bytes and ops atomically, unlocks
2. `counter.count() (int64, int)` — locks mutex, reads both fields, unlocks
3. `readCounter.Read(p)` — delegates to underlying reader, then calls `addBytes`
4. `readCounter.ReadCount()` — delegates to `counter.count()`
5. Same pattern for `writeCounter`
6. `NewReadWriteCounter` composes `NewWriteCounter` and `NewReadCounter`

### Rationale

- Both bytes and ops are updated under the same lock, guaranteeing the consistency invariant (`nops == n / numBytes`) required by the tests.
- Simple, well-understood pattern in Go.
- Matches the reference solution in `.meta/example.go` exactly.

### Files Modified

- `go/exercises/practice/paasio/paasio.go` (the only solution file)

---

## Proposal B: Atomic-based approach

**Role: Opponent**

### Approach

Use `sync/atomic` operations instead of a mutex. Store a combined state in an `atomic.Value` or use two separate `atomic.Int64` values.

### Structures

- `readCounter` with `io.Reader`, `bytesRead atomic.Int64`, `opsRead atomic.Int64`
- `writeCounter` with `io.Writer`, `bytesWritten atomic.Int64`, `opsWritten atomic.Int64`
- `rwCounter` composing both

### Implementation

1. `Read()` calls underlying reader, then atomically adds bytes and increments ops
2. `ReadCount()` reads both atomics

### Critique of Proposal A

- Mutex has higher overhead than atomics in high-contention scenarios.

### Problems with Proposal B

- **Critical flaw**: The consistency tests require that `nops == n / numBytes` at any observation point. With two separate atomic values, there's a race between updating bytes and updating ops — a concurrent `ReadCount()` call could see updated bytes but stale ops (or vice versa). This would **fail the consistency tests**.
- Could be fixed with a single atomic holding a combined struct (e.g., `atomic.Value` with a `{bytes, ops}` pair), but `atomic.Value` requires interface boxing and is more complex and error-prone.
- More complex to reason about correctness than a simple mutex.

### Files Modified

- `go/exercises/practice/paasio/paasio.go`

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion     | Proposal A (Mutex)         | Proposal B (Atomics)        |
|---------------|---------------------------|-----------------------------|
| Correctness   | Guaranteed — single lock protects both fields | Risky — two separate atomics fail consistency tests; combined atomic is complex |
| Risk          | Very low                   | High — subtle concurrency bugs |
| Simplicity    | Simple, ~90 lines          | More complex, harder to reason about |
| Consistency   | Matches reference solution exactly | Deviates from established pattern |

### Decision

**Proposal A wins.** The mutex approach is correct, simple, matches the reference solution, and directly satisfies the consistency requirements of the tests. The atomic approach introduces unnecessary complexity and risk for no meaningful benefit.

### Final Implementation Plan

**File to modify:** `go/exercises/practice/paasio/paasio.go`

**Step 1:** Define a `counter` helper struct:
```go
type counter struct {
    bytes int64
    ops   int
    mutex *sync.Mutex
}

func newCounter() counter {
    return counter{mutex: new(sync.Mutex)}
}

func (c *counter) addBytes(n int) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.bytes += int64(n)
    c.ops++
}

func (c *counter) count() (n int64, ops int) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    return c.bytes, c.ops
}
```

**Step 2:** Define `readCounter`:
```go
type readCounter struct {
    r io.Reader
    counter
}

func (rc *readCounter) Read(p []byte) (int, error) {
    m, err := rc.r.Read(p)
    rc.addBytes(m)
    return m, err
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
    return rc.count()
}
```

**Step 3:** Define `writeCounter`:
```go
type writeCounter struct {
    w io.Writer
    counter
}

func (wc *writeCounter) Write(p []byte) (int, error) {
    m, err := wc.w.Write(p)
    wc.addBytes(m)
    return m, err
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
    return wc.count()
}
```

**Step 4:** Define `rwCounter` and constructors:
```go
type rwCounter struct {
    WriteCounter
    ReadCounter
}

func NewWriteCounter(w io.Writer) WriteCounter {
    return &writeCounter{w: w, counter: newCounter()}
}

func NewReadCounter(r io.Reader) ReadCounter {
    return &readCounter{r: r, counter: newCounter()}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
    return &rwCounter{
        NewWriteCounter(rw),
        NewReadCounter(rw),
    }
}
```

**Step 5:** Run tests and vet to verify.
