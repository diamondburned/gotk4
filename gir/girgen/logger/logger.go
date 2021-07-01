// Package logger provides abstractions over girgen's logging.
package logger

import (
	"log"

	"github.com/fatih/color"
)

// Level vaguely describes the importance of each log entry.
type Level uint8

const (
	// Debug describes logging information that wouldn't be useful for anything
	// except for during debugging.
	Debug Level = iota
	// Skip describes an event that is only used when something is skipped.  The
	// reason for skipping should not be written here and should instead be in
	// Debug.
	Skip
	// Unusual describes an event that is unusual but is not fatal.
	Unusual
	// Error describes an event that is erroneous and will likely not produce a
	// valid output.
	Error
)

func (lvl Level) prefix() string {
	switch lvl {
	case Debug:
		return "debug:"
	case Skip:
		return "skipped:"
	case Unusual:
		return "unusuality:"
	case Error:
		return "error:"
	default:
		return ""
	}
}

func (lvl Level) colorf(f string, v ...interface{}) string {
	switch lvl {
	case Skip:
		return color.BlueString(f, v...)
	case Unusual:
		return color.YellowString(f, v...)
	case Error:
		return color.New(color.Bold, color.FgHiRed).Sprintf(f, v...)
	case Debug:
		fallthrough
	default:
		return color.New(color.Faint).Sprintf(f, v...)
	}
}

// LineLogger describes anything that can log itself.
type LineLogger interface {
	Logln(Level, ...interface{})
}

// NoopLogger is a logger that doesn't log anything. It is used as a placeholder
// to pass into functions that need logging.
var NoopLogger = noopLogger{}

type noopLogger struct{}

func (noop noopLogger) Logln(Level, ...interface{}) {}

// Prefix prepends the given prefixes into the given value list.
func Prefix(list []interface{}, p interface{}) []interface{} {
	list = append(list, nil)
	copy(list[1:], list)
	list[0] = p
	return list
}

// // Logln Logs using the Logger.
// func (g *Generator) Logln(level LogLevel, v ...interface{}) {
// 	if g.Logger == nil || g.LogLevel > level {
// 		return
// 	}

// 	prefix := level.prefix()
// 	if prefix != "" {
// 		if g.Color {
// 			prefix = level.colorf(prefix)
// 		}
// 		v = append(v, nil)
// 		copy(v[1:], v)
// 		v[0] = prefix
// 	}

// 	g.Logger.Println(v...)
// }

// Stdlog renders the given log entry into the stdlib logger.
func Stdlog(logger *log.Logger, minlevel, level Level, v ...interface{}) {
	if logger == nil || minlevel > level {
		return
	}

	prefix := level.prefix()
	if prefix != "" {
		prefix = level.colorf(prefix)
		v = Prefix(v, prefix)
	}

	logger.Println(v...)
}
