package emlog

import (
	"sync/atomic"
)

func (l *Logger) incrementLastID() uint32 {
	return atomic.AddUint32(&l.LastID, 1)
}
