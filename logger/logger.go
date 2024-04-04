package logger

import (
	"GophKeeper-Server/config"
	"log/slog"
	"os"
)

type Logger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
}

func NewLogger(cfg *config.Config) (Logger, error) {
	l := &Slogger{
		l: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
	}
	return l, nil
}

type Slogger struct {
	l *slog.Logger
}

func (s *Slogger) Error(msg string, args ...any) {
	s.l.Error(msg, args...)
}

func (s *Slogger) Info(msg string, args ...any) {
	s.l.Info(msg, args...)
}

func (s *Slogger) Debug(msg string, args ...any) {
	s.l.Debug(msg, args...)
}

// type Slogger struct {
// 	l *slog.Logger
// }

// func NewSlogger(cfg *config.Config) *Slogger {
// }

// func ConfigHandler(cfg *config.Config) (*slog.Handler, error) {
// 	lt := "text"
// 	ltInterface, ok := cfg.Logger["type"]
// 	if ok {
// 		lt, ok = ltInterface.(string)
// 		if !ok {
// 			return nil, fmt.Errorf("logger type must be a string")
// 		}
// 	}
// 	switch lt {
// 	case "text":
// 		return TextHandler(cfg)
// 	case "json":
// 		return JSONHandler(cfg)
// 	default:
// 		return nil, fmt.Errorf("unknown logger type: %s", lt)
// 	}
// }

// func JSONHandler(cfg *config.Config) (*slog.JSONHandler, error) {
// 	slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: })
// }

// func TextHandler(cfg *config.Config) (*slog.Handler, error) {
// 	slog.NewTextHandler()
// }
