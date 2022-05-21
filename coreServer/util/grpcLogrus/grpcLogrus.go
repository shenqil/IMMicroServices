package grpcLogrus

import (
	"github.com/sirupsen/logrus"
)

type Log struct {
	logger logrus.Logger
}

// NewLogrus
func NewLogrus(logger *logrus.Logger) *Log {
	return &Log{
		logger: *logger,
	}
}

// Info returns
func (l *Log) Info(args ...interface{}) {
	l.logger.Info(args)
}

// Infoln returns
func (l *Log) Infoln(args ...interface{}) {
	l.logger.Info(args...)
}

// Infof returns
func (l *Log) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warning returns
func (l *Log) Warning(args ...interface{}) {
	l.logger.Warn(args...)
}

// Warningln returns
func (l *Log) Warningln(args ...interface{}) {
	l.logger.Warn(args...)
}

// Warningf returns
func (l *Log) Warningf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Error returns
func (l *Log) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// Errorln returns
func (l *Log) Errorln(args ...interface{}) {
	l.logger.Error(args...)
}

// Errorf returns
func (l *Log) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatal returns
func (l *Log) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalln returns
func (l *Log) Fatalln(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalf logs to fatal level
func (l *Log) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (l *Log) V(v int) bool {
	return false
}
