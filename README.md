# go-concurrency-patterns

Synchronous communication - sender blocks when it sends data to Unbuffered channel until the data is read.
Asynchronous communication - sender does not block when it sends data to buffered channel. Unless buffer is at full capacity.