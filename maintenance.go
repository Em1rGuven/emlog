package emlog

import (
	"os"
	"time"
)

func (l *Logger) LoggerMaintenance() {
	file, err := os.OpenFile(l.Name+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()

	ticker := time.NewTicker(l.TimeCycle)
	for {
		select {
		case <-ticker.C:
			_ = file.Truncate(0)
		case <-l.SignalChannel:
			_ = file.Truncate(0)
		case <-l.Ctx.Done():
			return
		}
	}
}
