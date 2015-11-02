package ylog

import (
	"testing"
)

func TestLog(t *testing.T) {
	Init(".")
	SetDebug(true)
	Debug("debug")
	Error("error")
	Warning("warning")
	Info("info")
	Close()
}
