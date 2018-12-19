# mmap

[![GoDoc](https://godoc.org/github.com/dawidd6/mmap?status.svg)](https://godoc.org/github.com/dawidd6/mmap)

Package mmap implements a concurrent-safe map.

Simply the basic operations on maps but guarded with RWMutex.

### Example

```go
package main

import (
    "fmt"
    "mmap"
)

func main() {
    m := mmap.New()

    m.Set("key", "value")

    fmt.Println(m.Has("key")) // true
    fmt.Println(m.Get("key")) // value
    fmt.Println(m.Count()) // 1
    fmt.Println(m.Items()) // [value]

    m.Iterate(func(key interface{}, value interface{}) {
        fmt.Println(key, value) // key value
    })

    m.Remove("key")
}
```