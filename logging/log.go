package logging

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func Setup(debug *bool) {
	var level = slog.LevelInfo
	if debug != nil && *debug {
		level = slog.LevelDebug
	}

	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:     level,
			AddSource: true,
		}),
	))
}
