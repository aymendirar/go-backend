package http

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func (s *HTTPServer) ServeHTTP(w *httptest.ResponseRecorder, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func NewTestHTTPServer() *HTTPServer {
	router := gin.Default()
	server := &http.Server{
		Addr:    "",
		Handler: router.Handler(),
	}
	s := &HTTPServer{
		server: server,
		router: router,
	}

	s.route()
	return s
}
