package emlog

import (
	"sync/atomic"
)

func (l *Logger) incrementLastID() {
	atomic.AddUint32(&l.LastID, 1)
}
