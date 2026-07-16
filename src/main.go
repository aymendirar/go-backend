package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/aymendirar/go-backend/src/http"
	"github.com/aymendirar/go-backend/src/util"
)

func main() {
	if err := _main(); err != nil {
		slog.Error("main error", "error", err)
		os.Exit(1)
	}
}

func _main() error {
	slog.Info("backend starting up...")

	env, err := util.LoadEnv()
	if err != nil {
		return err
	}

	shutdownCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	s := http.NewHTTPServer(env.HOST, env.PORT)
	if err := s.Run(shutdownCtx); err != nil {
		return err
	}
	slog.Info("exiting...")
	return nil
}
