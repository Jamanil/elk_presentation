package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

const (
	enableJSONLogs = true
	logsMaxDelay   = time.Second
)

func main() {
	var logger *slog.Logger

	logFile, err := os.OpenFile("logs/log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(fmt.Errorf("failed to create log file: %w", err))
	}

	logWriter := io.MultiWriter(os.Stdout, logFile)

	if enableJSONLogs {
		logger = slog.New(slog.NewJSONHandler(logWriter, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(logWriter, nil))
	}

	ctx := context.Background()

	for {
		time.Sleep(time.Duration(rand.Int64N(logsMaxDelay.Nanoseconds())))

		logger.Log(ctx,
			slog.Level(gofakeit.RandomInt([]int{
				int(slog.LevelDebug),
				int(slog.LevelInfo),
				int(slog.LevelWarn),
				int(slog.LevelError),
			})),
			gofakeit.HackerPhrase())
	}
}
