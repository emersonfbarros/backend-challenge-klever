package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	err    *log.Logger
	writer io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		err:    log.New(writer, "ERROR: ", logger.Flags()),
		writer: writer,
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
