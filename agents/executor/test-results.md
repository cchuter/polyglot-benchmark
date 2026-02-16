# Paasio Test Results

## Build

```
$ go build ./...
```

**Result:** Build succeeded with no errors.

## Tests

```
$ go test -cpu 2 -v ./...
=== RUN   TestMultiThreaded
--- PASS: TestMultiThreaded (0.00s)
=== RUN   TestWriteWriter
--- PASS: TestWriteWriter (0.00s)
=== RUN   TestWriteReadWriter
--- PASS: TestWriteReadWriter (0.00s)
=== RUN   TestReadReader
--- PASS: TestReadReader (0.07s)
=== RUN   TestReadReadWriter
--- PASS: TestReadReadWriter (0.06s)
=== RUN   TestReadTotalReader
    paasio_test.go:117: Calling Read() for 8000*50=400000 bytes
--- PASS: TestReadTotalReader (0.02s)
=== RUN   TestReadTotalReadWriter
    paasio_test.go:117: Calling Read() for 8000*50=400000 bytes
--- PASS: TestReadTotalReadWriter (0.01s)
=== RUN   TestWriteTotalWriter
    paasio_test.go:156: Calling Write() with 8000*50=400000 bytes
--- PASS: TestWriteTotalWriter (0.01s)
=== RUN   TestWriteTotalReadWriter
    paasio_test.go:156: Calling Write() with 8000*50=400000 bytes
--- PASS: TestWriteTotalReadWriter (0.01s)
=== RUN   TestReadCountConsistencyReader
--- PASS: TestReadCountConsistencyReader (0.01s)
=== RUN   TestReadCountConsistencyReadWriter
--- PASS: TestReadCountConsistencyReadWriter (0.01s)
=== RUN   TestWriteCountConsistencyWriter
--- PASS: TestWriteCountConsistencyWriter (0.01s)
=== RUN   TestWriteCountConsistencyReadWriter
--- PASS: TestWriteCountConsistencyReadWriter (0.01s)
PASS
ok  	paasio	0.209s
```

## Summary

- **Total tests:** 13
- **Passed:** 13
- **Failed:** 0
- **Status:** ALL PASS
