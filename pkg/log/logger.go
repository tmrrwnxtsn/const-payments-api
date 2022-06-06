// Package log предоставляет возможности структурированного логгирования на разных уровнях.
package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Logger представляет логгер с уровнями логирования debug, info, error.
type Logger interface {
	// Debug использует fmt.Sprint, чтобы зафиксировать сообщение на уровне debug.
	Debug(args ...interface{})
	// Info использует fmt.Sprint, чтобы зафиксировать сообщение на уровне info.
	Info(args ...interface{})
	// Error использует fmt.Sprint, чтобы зафиксировать сообщение на уровне error.
	Error(args ...interface{})

	// Debugf использует fmt.Sprintf чтобы зафиксировать сообщение на уровне debug.
	Debugf(format string, args ...interface{})
	// Infof использует fmt.Sprintf чтобы зафиксировать сообщение на уровне info
	Infof(format string, args ...interface{})
	// Errorf использует fmt.Sprintf чтобы зафиксировать сообщение на уровне error.
	Errorf(format string, args ...interface{})

	// SetLoggingLevel устанавливает уровень логгирования (debug, info, error).
	SetLoggingLevel(loggingLevel string) error
}

type logger struct {
	*logrus.Logger
}

// New создаёт новый логгер. По умолчанию уровень логгирования – info. Можно изменить методом Logger.SetLoggingLevel.
func New() Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetOutput(os.Stdout)
	return &logger{l}
}

// SetLoggingLevel устанавливает уровень логгирования.
func (l *logger) SetLoggingLevel(loggingLevel string) error {
	lvl, err := logrus.ParseLevel(loggingLevel)
	if err != nil {
		return err
	}
	l.SetLevel(lvl)
	return nil
}
