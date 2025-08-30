# emlog
A lightweight logging package for Go with file size and time-based log rotation.

# Installation
```bash
go get github.com/Em1rGuven/emlog
```

## Quickstart 
```go
import "github.com/Em1rGuven/emlog"
```

### Usage
```go
logger := emlog.NewLogger("mylogs", 72, 20) // "mylogs.log" created, 72 hours lifecycle, 20 MB max size
logger.CreateLog("admin", "internal server error") // {"id":1,"user":"admin","content":"internal server error","time":"2025-08-30 12:00:00"}
logger.CloseLogger() // logger is closed (for preventing memory leaks)
```



