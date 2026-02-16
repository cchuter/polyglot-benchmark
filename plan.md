# Implementation Plan: paasio (Issue #227)

## Branch 1: Shared Counter Struct with Embedding (Minimal, DRY)

Use a shared `counter` struct that handles mutex-protected byte/ops tracking, then embed it into `readCounter` and `writeCounter`. Compose `ReadWriteCounter` from the two independent counters.

### Files to Modify
- `go/exercises/practice/paasio/paasio.go`

### Approach
1. Define a `counter` struct with `bytes int64`, `ops int`, `mutex *sync.Mutex`
2. Add `newCounter()` constructor, `addBytes(n int)`, and `count() (int64, int)` methods
3. Define `readCounter` struct embedding `counter` and holding `r io.Reader`
4. Define `writeCounter` struct embedding `counter` and holding `w io.Writer`
5. Define `rwCounter` struct embedding `WriteCounter` and `ReadCounter`
6. Implement `NewReadCounter`, `NewWriteCounter`, `NewReadWriteCounter`

### Rationale
This is the canonical Exercism solution pattern. Minimal code, DRY via embedding, straightforward.

### Evaluation
- **Feasibility**: High — directly mirrors the reference solution
- **Risk**: Very low — proven pattern
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~90 lines, 1 file

---

## Branch 2: Separate Read/Write Structs Without Shared Counter (Explicit)

Each struct (`readCounter`, `writeCounter`) has its own inline mutex and counters without a shared base type. The `rwCounter` embeds both.

### Files to Modify
- `go/exercises/practice/paasio/paasio.go`

### Approach
1. Define `readCounter` with `r io.Reader`, `bytes int64`, `ops int`, `mu sync.Mutex`
2. Define `writeCounter` with `w io.Writer`, `bytes int64`, `ops int`, `mu sync.Mutex`
3. Each has its own `Read`/`Write` and `ReadCount`/`WriteCount` methods
4. Define `rwCounter` embedding `*readCounter` and `*writeCounter`
5. Implement `NewReadWriteCounter` by creating both and composing

### Rationale
More explicit, no shared abstraction. Easier to understand at a glance without needing to understand embedding of the `counter` type.

### Evaluation
- **Feasibility**: High
- **Risk**: Low, but more code duplication
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~100 lines, 1 file, some duplication

---

## Branch 3: Atomic Operations (Performance-Oriented)

Use `sync/atomic` instead of `sync.Mutex` for counter updates. Pack bytes and ops into a single atomic struct value to ensure consistency.

### Files to Modify
- `go/exercises/practice/paasio/paasio.go`

### Approach
1. Use `atomic.Int64` for bytes and `atomic.Int32` for ops (Go 1.19+ types)
2. However, the consistency tests require that bytes and ops are updated atomically *together* — `atomic` alone can't guarantee this without a mutex or `atomic.Value` with a struct
3. Could use `atomic.Value` storing a `{bytes, ops}` struct, with CAS loop for updates

### Rationale
Lock-free performance under high concurrency.

### Evaluation
- **Feasibility**: Medium — `atomic.Int64` requires Go 1.19, but `go.mod` specifies Go 1.18. Would need `sync/atomic` primitives. CAS loop is complex.
- **Risk**: High — consistency requirement (bytes and ops must be updated atomically together) makes this significantly harder. CAS loops are error-prone.
- **Alignment**: Could satisfy criteria but with much higher implementation risk
- **Complexity**: ~120+ lines, more complex logic, potential Go version issues

---

## Selected Plan

**Branch 1: Shared Counter Struct with Embedding** is selected.

### Rationale
- It is the simplest approach with the least code
- It perfectly matches the reference solution, reducing risk to near zero
- The DRY `counter` type eliminates duplication while remaining readable
- Go 1.18 compatible (only uses `sync.Mutex`)
- Branch 2 is viable but duplicates code unnecessarily
- Branch 3 has Go version compatibility issues and high complexity for the consistency requirement

### Detailed Implementation

**File: `go/exercises/practice/paasio/paasio.go`**

```go
package paasio

import (
	"io"
	"sync"
)

// counter provides thread-safe byte and operation counting.
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

// readCounter wraps an io.Reader with byte/operation counting.
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

// writeCounter wraps an io.Writer with byte/operation counting.
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

// rwCounter composes read and write counters.
type rwCounter struct {
	WriteCounter
	ReadCounter
}

// NewWriteCounter returns a WriteCounter wrapping the given writer.
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{
		w:       w,
		counter: newCounter(),
	}
}

// NewReadCounter returns a ReadCounter wrapping the given reader.
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{
		r:       r,
		counter: newCounter(),
	}
}

// NewReadWriteCounter returns a ReadWriteCounter wrapping the given read-writer.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{
		NewWriteCounter(rw),
		NewReadCounter(rw),
	}
}
```

### Ordering of Changes
1. Write the complete solution to `paasio.go`
2. Run `go test -cpu 2 ./...` in the exercise directory
3. Verify all tests pass
