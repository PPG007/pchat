package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"pchat/utils/env"
	"time"
)

const (
	traceKey = "backtrace"

	MaxStackSize = 4096
)

type Fields = logrus.Fields

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		ForceColors:     true,
		FullTimestamp:   true,
	})
	logrus.SetLevel(logrus.WarnLevel)
	if env.IsDebug() {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetOutput(os.Stdout)
}

func Info(ctx context.Context, msg string, extra logrus.Fields) {
	entry := newTraceEntry(ctx, extra)
	if msg != "" {
		entry.Info(msg)
	} else {
		entry.Print()
	}
}

func Warn(ctx context.Context, msg string, extra logrus.Fields) {
	newTraceEntry(ctx, extra).Warn(msg)
}

func WarnTrace(ctx context.Context, msg string, extra logrus.Fields, trace []byte) {
	newTraceEntry(ctx, extra).WithField(traceKey, string(trace)).Warn(msg)
}

func Error(ctx context.Context, msg string, extra logrus.Fields) {
	newTraceEntry(ctx, extra).Error(msg)
}

func ErrorTrace(ctx context.Context, msg string, extra logrus.Fields, trace []byte) {
	newTraceEntry(ctx, extra).WithField(traceKey, string(trace)).Error(msg)
}

func newTraceEntry(ctx context.Context, extra logrus.Fields) *logrus.Entry {
	entry := logrus.WithFields(extra)
	// TODO:
	return entry
}
