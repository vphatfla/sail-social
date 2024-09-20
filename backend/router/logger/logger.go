package logger

import (
	"log/slog"
	"time"

	"github.com/go-chi/httplog/v2"
)

func Logger() *httplog.Logger {
	// Logger
	logger := httplog.NewLogger("booster-hub-backend", httplog.Options{
		LogLevel: slog.LevelDebug,
		// JSON:             true,
		Concise: true,
		// RequestHeaders:   true,
		// ResponseHeaders:  true,
		MessageFieldName: "message",
		LevelFieldName:   "severity",
		TimeFieldFormat:  time.RFC3339,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		// QuietDownRoutes: []string{
		// 	"/",
		// 	"/ping",
		// },
		// QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})

	return logger
}
