package emlog

import (
	"encoding/json"
	"os"
)

func (l *Logger) ProcessLogs() {
	file, err := os.OpenFile(l.Name+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()

	for v := range l.LogChannel {
		payload, err := json.Marshal(v)
		if err != nil {
			return
		}
		content := append(payload, '\n')
		_, err = file.Write(content)
		if err != nil {
			return
		}

		info, err := file.Stat()
		if err != nil {
			return
		} else if info.Size() > int64(l.MaxFileSizeMB<<20) {
			l.SignalChannel <- struct{}{}
		}
	}
}
