package types

import (
	"context"
	"time"
)

type (
	Log struct {
		ID      uint32 `json:"id"`
		User    string `json:"user"`
		Content string `json:"content"`
		Time    string `json:"time"`
	}

	Logger struct {
		Name          string
		LogChannel    chan *Log
		SignalChannel chan struct{}
		LastID        uint32
		MaxFileSizeMB int
		TimeCycle     time.Duration
		Ctx           context.Context
		Cancel        context.CancelFunc
	}
)
