package xlog

import (
	"testing"
	"time"

	"github.com/binz96/xlog"
)

func TestFileLog(t *testing.T) {
	ti := time.Now()
	today := ti.Format("2006-01-02")
	logger := xlog.NewFileLogger(xlog.DebugLevel, "log/"+today+".txt")
	logger.Debug("第一行日志")
}

func TestConsoleLog(t *testing.T) {
	logger := xlog.NewConsoleLogger(xlog.DebugLevel)
	logger.Debug("第一行日志")
}
