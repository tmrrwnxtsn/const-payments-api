package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Logger представляет логгер, который поддерживает уровни логирования (DEBUG, INFO, ERROR).
type Logger interface {
	// Debugf использует fmt.Sprintf чтобы зафиксировать сообщение на уровне DEBUG
	Debugf(format string, args ...interface{})
	// Infof использует fmt.Sprintf чтобы зафиксировать сообщение на уровне INFO
	Infof(format string, args ...interface{})
	// Errorf использует fmt.Sprintf чтобы зафиксировать сообщение на уровне ERROR
	Errorf(format string, args ...interface{})
}

type logger struct {
	*logrus.Logger
}

// New создаёт новый логгер с конфигурацией по умолчанию.
func New() Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetReportCaller(true)
	l.SetOutput(os.Stdout)
	return &logger{l}
}
