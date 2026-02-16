# Changes

## paasio.go — Implement IO statistics wrapper

- Added `counter` struct with mutex-protected byte/op tracking
- Implemented `readCounter` wrapping `io.Reader` with `Read` and `ReadCount` methods
- Implemented `writeCounter` wrapping `io.Writer` with `Write` and `WriteCount` methods
- Implemented `rwCounter` composing `ReadCounter` and `WriteCounter` interfaces
- Added constructor functions: `NewReadCounter`, `NewWriteCounter`, `NewReadWriteCounter`
