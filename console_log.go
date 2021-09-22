package xlog

import (
	"os"
)

type ConsoleLog struct {
	level int
}

func NewConsoleLogger(level int) Logger {
	return &ConsoleLog{
		level: level,
	}
}

func (c *ConsoleLog) SetLevel(level int) {
	if level < DebugLevel || level > FatalLevel {
		c.level = DebugLevel
	}
	c.level = level
}

func (c *ConsoleLog) Debug(format string, a ...interface{}) {
	if c.level > DebugLevel {
		return
	}
	write_log(os.Stdout, DebugLevel, format, a...)
}

func (c *ConsoleLog) Trace(format string, a ...interface{}) {
	if c.level > TraceLevel {
		return
	}
	write_log(os.Stdout, TraceLevel, format, a...)
}

func (c *ConsoleLog) Info(format string, a ...interface{}) {
	if c.level > InfoLevel {
		return
	}
	write_log(os.Stdout, InfoLevel, format, a...)
}

func (c *ConsoleLog) Warn(format string, a ...interface{}) {
	if c.level > WarnLevel {
		return
	}
	write_log(os.Stdout, WarnLevel, format, a...)
}

func (c *ConsoleLog) Error(format string, a ...interface{}) {
	if c.level > ErrorLevel {
		return
	}
	write_log(os.Stdout, ErrorLevel, format, a...)
}

func (c *ConsoleLog) Fatal(format string, a ...interface{}) {
	if c.level > FatalLevel {
		return
	}
	write_log(os.Stdout, FatalLevel, format, a...)
	os.Exit(1)
}
