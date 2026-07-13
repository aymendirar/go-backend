package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	server *http.Server
	router *gin.Engine
}

const ADDRESS string = ":8000"

func NewHTTPServer() *HTTPServer {
	router := gin.Default()
	server := &http.Server{
		Addr:    ADDRESS,
		Handler: router.Handler(),
	}

	return &HTTPServer{
		server: server,
		router: router,
	}
}

func NewTestHTTPServer() *HTTPServer {
	router := gin.Default()
	server := &http.Server{
		Addr:    ADDRESS,
		Handler: router.Handler(),
	}
	s := &HTTPServer{
		server: server,
		router: router,
	}

	s.route()
	return s
}

func (s *HTTPServer) route() {
	s.router.POST("/ping", s.PingRequest())
}

func (s *HTTPServer) Run() {
	s.route()
	go func() error {
		return s.server.ListenAndServe()
	}()
	slog.Info("http server running...")
}

func (s *HTTPServer) Stop(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", "error", err)
	}
}
