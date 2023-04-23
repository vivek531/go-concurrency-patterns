# go-concurrency-patterns

Notes:

1. Synchronous communication - sender blocks when it sends data to Unbuffered channel until the data is read.
2. Asynchronous communication - sender does not block when it sends data to buffered channel. Unless buffer is at full capacity.
