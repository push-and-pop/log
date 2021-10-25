package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// const (
// 	nocolor   = "0"
// 	red       = "31"
// 	green     = "38;5;48"
// 	yellow    = "33"
// 	gray      = "38;5;251"
// 	graybold  = "1;38;5;251"
// 	lightgray = "38;5;243"
// 	cyan      = "1;36"
// )

const (
	Trace_   = "Trace: "
	Info_    = "Info: "
	Warning_ = "Warning: "
	Error_   = "Error: "
	Fatal_   = "Fatal: "
)

const (
	DateFormat = "2006-01-02 15:04:05"
)

type Logger interface {
	Trace(format string, msg ...interface{}) error
	Info(format string, msg ...interface{}) error
	Warning(format string, msg ...interface{}) error
	Error(format string, msg ...interface{}) error
	Fatal(format string, msg ...interface{}) error
}

var defaultLog Logger = SetDefaultLog()

type logger struct {
	mu  sync.Mutex
	out io.Writer
	buf []byte
}

func New(out ...io.Writer) Logger {
	return &logger{
		out: complexWriter(out...),
	}
}

func SetDefaultLog(w ...io.Writer) Logger {
	if len(w) == 0 {
		return New(os.Stdout)
	}
	return New(w...)
}

//复合writer
func complexWriter(w ...io.Writer) io.Writer {
	return io.MultiWriter(w...)
}

func (l *logger) Trace(format string, msg ...interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf = l.buf[:0]
	l.buf = formatf(Trace_, format+"\n", msg...)
	return l.outPut()
}
func (l *logger) Info(format string, msg ...interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf = l.buf[:0]
	l.buf = formatf(Info_, format+"\n", msg...)
	return l.outPut()
}

func (l *logger) Warning(format string, msg ...interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf = l.buf[:0]
	l.buf = formatf(Warning_, format+"\n", msg...)
	return l.outPut()
}

func (l *logger) Error(format string, msg ...interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf = l.buf[:0]
	l.buf = formatf(Error_, format+"\n", msg...)
	return l.outPut()
}

func (l *logger) Fatal(format string, msg ...interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf = l.buf[:0]
	l.buf = formatf(Fatal_, format+"\n", msg...)
	return l.outPut()
}

func (l *logger) outPut() error {

	_, err := l.out.Write(l.buf)
	return err

}

func timeNow() string {
	return time.Now().Format(DateFormat)
}

func formatf(Type string, format string, msg ...interface{}) []byte {
	return []byte(fmt.Sprintf(timeNow()+" "+Type+format, msg...))
}

func Trace(format string, msg ...interface{}) error {
	return defaultLog.Trace(format, msg...)
}
func Info(format string, msg ...interface{}) error {
	return defaultLog.Info(format, msg...)
}

func Warning(format string, msg ...interface{}) error {
	return defaultLog.Warning(format, msg...)
}

func Error(format string, msg ...interface{}) error {
	return defaultLog.Error(format, msg...)
}

func Fatal(format string, msg ...interface{}) error {
	return defaultLog.Fatal(format, msg...)
}
