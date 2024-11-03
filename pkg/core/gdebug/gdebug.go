package gdebug

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"slices"
	"strings"
)

var debug = strings.Split(os.Getenv("GOTK4_DEBUG"), ",")

func HasKey(key string) bool {
	return slices.Contains(debug, key) || slices.Contains(debug, "all")
}

func NewDebugLogger(key string) *slog.Logger {
	if !HasKey(key) {
		return slog.New(noopLogHandler{})
	}
	return mustDebugLogger(key)
}

func NewDebugLoggerNullable(key string) *slog.Logger {
	if !HasKey(key) {
		return nil
	}
	return mustDebugLogger(key)
}

func mustDebugLogger(name string) *slog.Logger {
	if HasKey("to-console") {
		return slog.With("gotk4_module", name)
	}

	f, err := os.CreateTemp(os.TempDir(), fmt.Sprintf("gotk4-%s-%d-*", name, os.Getpid()))
	if err != nil {
		log.Panicln("cannot create temp", name, "file:", err)
	}

	slog.Info(
		"gotk4: intern: enabled debug file",
		"file", f.Name())

	return slog.New(
		slog.NewTextHandler(f, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)
}

type noopLogHandler struct{}

var _ slog.Handler = noopLogHandler{}

func (noopLogHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (noopLogHandler) Handle(context.Context, slog.Record) error { return nil }
func (noopLogHandler) WithAttrs([]slog.Attr) slog.Handler        { return noopLogHandler{} }
func (noopLogHandler) WithGroup(string) slog.Handler             { return noopLogHandler{} }
