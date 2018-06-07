package restful

import (
	stdlog "log"
	"os"
)

// StdLogger corresponds to a minimal subset of the interface satisfied by stdlib log.Logger
type StdLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

var Logger StdLogger

func init() {
	// default Logger
	SetLogger(stdlog.New(os.Stderr, "[restful] ", stdlog.LstdFlags|stdlog.Lshortfile))
}

// Print delegates to the Logger
func Print(v ...interface{}) {
	Logger.Print(v...)
}

// Printf delegates to the Logger
func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}
