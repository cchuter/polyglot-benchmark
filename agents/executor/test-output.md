# Test Output: error-handling

## Command
```
go test -v ./...
```

## Results: PASS (5/5)

```
=== RUN   TestNoErrors
--- PASS: TestNoErrors (0.00s)
=== RUN   TestKeepTryOpenOnTransient
--- PASS: TestKeepTryOpenOnTransient (0.00s)
=== RUN   TestFailOpenOnNonTransient
--- PASS: TestFailOpenOnNonTransient (0.00s)
=== RUN   TestCallDefrobAndCloseOnFrobError
--- PASS: TestCallDefrobAndCloseOnFrobError (0.00s)
=== RUN   TestCallCloseOnNonFrobError
--- PASS: TestCallCloseOnNonFrobError (0.00s)
PASS
ok  	erratum	0.005s
```
