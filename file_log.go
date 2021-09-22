package xlog

import (
	"os"
)

type FileLog struct {
	level int
	path  string
	file  *os.File
}

func NewFileLogger(level int, path string) Logger {
	logger := FileLog{
		level: level,
		path:  path,
	}
	var err error
	logger.file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return &logger
}

func (f *FileLog) SetLevel(level int) {
	if level < DebugLevel || level > FatalLevel {
		f.level = DebugLevel
	}
	f.level = level
}

func (f *FileLog) Debug(format string, a ...interface{}) {
	if f.level > DebugLevel {
		return
	}
	write_log(f.file, DebugLevel, format, a...)
}

func (f *FileLog) Trace(format string, a ...interface{}) {
	if f.level > TraceLevel {
		return
	}
	write_log(f.file, TraceLevel, format, a...)
}

func (f *FileLog) Info(format string, a ...interface{}) {
	if f.level > InfoLevel {
		return
	}
	write_log(f.file, InfoLevel, format, a...)
}

func (f *FileLog) Warn(format string, a ...interface{}) {
	if f.level > WarnLevel {
		return
	}
	write_log(f.file, WarnLevel, format, a...)
}

func (f *FileLog) Error(format string, a ...interface{}) {
	if f.level > ErrorLevel {
		return
	}
	write_log(f.file, ErrorLevel, format, a...)
}

func (f *FileLog) Fatal(format string, a ...interface{}) {
	if f.level > FatalLevel {
		return
	}
	write_log(f.file, FatalLevel, format, a...)
	f.file.Close()
	os.Exit(1)
}
