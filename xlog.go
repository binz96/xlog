package xlog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

const (
	DebugLevel = iota
	TraceLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var levelString = map[int]string{
	DebugLevel: "DEBUG",
	TraceLevel: "TRACE",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	FatalLevel: "FATAL",
}

type Logger interface {
	SetLevel(level int)
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

func write_log(f *os.File, level int, format string, a ...interface{}) {
	t := time.Now()
	fmt.Fprintf(f, "[%s] [%s] ", levelString[level], t.Format("2006.01.02 15:04:05"))
	callerpc, callerfile, callerline, ok := runtime.Caller(2)
	if ok {
		fmt.Fprintf(f, "[%s-%d: %s] ", path.Base(callerfile), callerline, path.Base(runtime.FuncForPC(callerpc).Name()))
	}
	fmt.Fprintf(f, format, a...)
	fmt.Fprintf(f, "\n")
}
