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

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{r: r}
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{w: w}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{
		WriteCounter: NewWriteCounter(rw),
		ReadCounter:  NewReadCounter(rw),
	}
}
