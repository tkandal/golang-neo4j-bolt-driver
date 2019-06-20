package buffer

import (
	"bytes"
	"sync"
)

// Pool is a pool where to get a temporary byte-buffer, which is not garbage-collected.
// Useful for programs that need to respond very fast, and re-use a lot of buffers.
var Pool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
