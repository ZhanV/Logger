package logger

import (
	"fmt"
	"testing"
	"time"
)

func TestNewFileLogger(t *testing.T) {

	err := Init()
	if err != nil {
		fmt.Printf("init log failure : %s", err)
		return
	}

	Debug("这是一条DEBUG消息: %s","hahahaha")
	Info("这是一条Info消息: %s","hahahaha")
	Warn("这是一条Warn消息: %s","hahahaha")
	Error("这是一条Error消息: %s","hahahaha")

	time.Sleep(3*time.Second)

}
