package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type HTTPServer struct {
	server *http.Server
	router *gin.Engine
}

func NewHTTPServer(host, port string) *HTTPServer {
	router := gin.Default()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router.Handler(),
	}

	return &HTTPServer{
		server: server,
		router: router,
	}
}

func (s *HTTPServer) route() {
	api := s.router.Group("/api")
	api.POST("/ping", s.PingRequest())
}

func (s *HTTPServer) Run(ctx context.Context) error {
	s.route()
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		slog.Info(fmt.Sprintf("http server running on %v", s.server.Addr))
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		return s.server.Shutdown(shutdownCtx)
	})

	return g.Wait()
}
