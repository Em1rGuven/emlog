package emlog

import (
	"context"
	"fmt"
	"github.com/Em1rGuven/emlog/types"
	"math"
	"os"
	"time"
)

type Logger struct {
	*types.Logger
}

func NewLogger(name string, reset time.Duration, maxSize int) *Logger {
	if reset == -1 {
		reset = math.MaxInt32
	} else if reset < 0 {
		reset = 72 // default
	}

	if maxSize == -1 {
		maxSize = math.MaxInt32
	} else if maxSize < 0 {
		maxSize = 50 // default
	}

	fileName := fmt.Sprintf("%s.log", name)
	f, _ := os.Create(fileName)
	defer func() { _ = f.Close() }()
	ctx, cancel := context.WithCancel(context.Background())
	newLogger := &Logger{
		&types.Logger{
			Name:          name,
			LogChannel:    make(chan *types.Log, 500),
			SignalChannel: make(chan struct{}, 500),
			LastID:        0,
			MaxFileSizeMB: maxSize,
			TimeCycle:     reset * time.Hour,
			Ctx:           ctx,
			Cancel:        cancel,
		},
	}
	go newLogger.LoggerMaintenance()
	go newLogger.ProcessLogs()

	return newLogger
}

func (l *Logger) CloseLogger() {
	l.Cancel()
	close(l.LogChannel)
}

func (l *Logger) CreateLog(user, content string) {
	l.incrementLastID()

	l.LogChannel <- &types.Log{
		ID:      l.LastID,
		User:    user,
		Content: content,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
}
