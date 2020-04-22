package liblogging

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmlogrus"
)

type (
	// Logger is ...
	Logger struct {
		Out   *logrus.Logger
		Err   *logrus.Logger
		Field *logrus.Entry
	}
)

// New is ...
func New() *Logger {
	tx := apm.DefaultTracer.StartTransaction("name", "type")
	defer tx.End()

	ctx := apm.ContextWithTransaction(context.Background(), tx)
	span, ctx := apm.StartSpan(ctx, "name", "type")
	defer span.End()
	return &Logger{
		Out: &logrus.Logger{
			Formatter: new(logrus.TextFormatter),
			Out:       os.Stdout,
			Level:     logrus.InfoLevel,
		},
		Err: &logrus.Logger{
			Formatter: new(logrus.TextFormatter),
			Out:       os.Stderr,
			Level:     logrus.InfoLevel,
		},
		Field: logrus.WithFields(apmlogrus.TraceContext(ctx)),
	}
}
