package logger

import (
	"GophKeeper-Server/config"
	"fmt"
	"log/slog"
	"os"
)

type Logger interface {
	Error(msg any)
	Errorf(msg string, args ...interface{})
	Info(msg any)
	Infof(msg string, args ...interface{})
	Debug(msg any)
	Debugf(msg string, args ...interface{})
}

type Slogger struct {
	l *slog.Logger
}

func NewSlogger(cfg *config.Config) *Slogger {
}

func ConfigHandler(cfg *config.Config) (*slog.Handler, error) {
	lt := "text"
	ltInterface, ok := cfg.Logger["type"]
	if ok {
		lt, ok = ltInterface.(string)
		if !ok {
			return nil, fmt.Errorf("logger type must be a string")
		}
	}
	switch lt {
	case "text":
		return TextHandler(cfg)
	case "json":
		return JSONHandler(cfg)
	default:
		return nil, fmt.Errorf("unknown logger type: %s", lt)
	}
}

func JSONHandler(cfg *config.Config) (*slog.JSONHandler, error) {
	slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: })
}

func TextHandler(cfg *config.Config) (*slog.Handler, error) {
	slog.NewTextHandler()
}
