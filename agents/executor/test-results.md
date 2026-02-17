# Paasio Test Results

## Build Result: PASS

Build completed successfully with no errors.

## Test Output

```
=== RUN   TestMultiThreaded
--- PASS: TestMultiThreaded (0.00s)
=== RUN   TestWriteWriter
--- PASS: TestWriteWriter (0.00s)
=== RUN   TestWriteReadWriter
--- PASS: TestWriteReadWriter (0.00s)
=== RUN   TestReadReader
--- PASS: TestReadReader (0.37s)
=== RUN   TestReadReadWriter
--- PASS: TestReadReadWriter (0.27s)
=== RUN   TestReadTotalReader
    paasio_test.go:117: Calling Read() for 8000*50=400000 bytes
--- PASS: TestReadTotalReader (0.14s)
=== RUN   TestReadTotalReadWriter
    paasio_test.go:117: Calling Read() for 8000*50=400000 bytes
--- PASS: TestReadTotalReadWriter (0.07s)
=== RUN   TestWriteTotalWriter
    paasio_test.go:156: Calling Write() with 8000*50=400000 bytes
--- PASS: TestWriteTotalWriter (0.08s)
=== RUN   TestWriteTotalReadWriter
    paasio_test.go:156: Calling Write() with 8000*50=400000 bytes
--- PASS: TestWriteTotalReadWriter (0.06s)
=== RUN   TestReadCountConsistencyReader
--- PASS: TestReadCountConsistencyReader (0.07s)
=== RUN   TestReadCountConsistencyReadWriter
--- PASS: TestReadCountConsistencyReadWriter (0.04s)
=== RUN   TestWriteCountConsistencyWriter
--- PASS: TestWriteCountConsistencyWriter (0.04s)
=== RUN   TestWriteCountConsistencyReadWriter
--- PASS: TestWriteCountConsistencyReadWriter (0.04s)
PASS
ok  	paasio	2.253s
```

## Test Result: PASS

## Summary

- **Tests passed: 13/13**
- **Tests failed: 0**
- **Race detector: enabled (-race flag), no races detected**
- **CPU: 2**
- **Total time: 2.253s**
