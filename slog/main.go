package main

import (
	"log/slog"
	"math/rand"
	"os"
	"runtime"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stderr,&slog.HandlerOptions{
		Level: slog.LevelDebug,
	}).WithAttrs([]slog.Attr{
		slog.String("app_version","v1.0.0.0"),
	}))
	slog.SetDefault(logger)

	slog.Info("Golang rock",
		slog.String("version",runtime.Version()),
		slog.Int("Random number", rand.Int()),
		slog.Group("OS Info",
				slog.String("OS",runtime.GOOS),
				slog.Int("CPUs",runtime.NumCPU()),
				slog.String("arch", runtime.GOARCH),
			),
	)
	slog.Error("Gopher has stumbled!")
	slog.Warn("Gopher's getting dizzy")
	slog.Debug("Debugging")
}