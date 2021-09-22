## xlog v0.01
> very simple and naive, to be written better
> TODO: 异步写日志

## usage
```bash
$ go get -u github.com/binz96/xlog
```

```go
package main

import (
    "fmt"

    "github.com/binz96/xlog"
)

func main() {
    // logger := xlog.NewConsoleLogger(xlog.Debug)
    logger := xlog.NewFileLogger(xlog.DebugLevel, "./log.txt")
    logger.Debug("Hello, xlog~")
}
```