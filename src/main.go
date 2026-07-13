package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aymendirar/go-backend/src/http"
)

func main() {
	if err := _main(); err != nil {
		slog.Error("main error", "error", err)
		os.Exit(1)
	}
}

func _main() error {
	slog.Info("backend starting up...")
	http := http.NewHTTPServer()
	http.Run()

	shutdownCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-shutdownCtx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	http.Stop(ctx)
	slog.Info("exiting...")
	return nil
}
