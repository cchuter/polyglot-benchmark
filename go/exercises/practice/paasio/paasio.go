package paasio

import (
	"io"
	"sync"
)

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

func (c *counter) count() (int64, int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
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
	return &readCounter{r: r, counter: newCounter()}
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{w: w, counter: newCounter()}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{
		NewWriteCounter(rw),
		NewReadCounter(rw),
	}
}
