package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	traceKey = "backtrace"
)

type Fields = logrus.Fields

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     true,
	})
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
