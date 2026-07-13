package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingRequest struct {
	Ping string `json:"ping"`
}

type PingResponse struct {
	Response string `json:"response"`
}

func (s *HTTPServer) PingRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request PingRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, PingResponse{Response: "try sending JSON"})
			return
		}

		if request.Ping != "pong" {
			c.JSON(http.StatusBadRequest, PingResponse{Response: "incorrect field"})
			return
		}

		c.JSON(http.StatusOK, PingResponse{Response: "ok"})
	}
}

func (s *HTTPServer) Ping(r string) bool {
	return r == "pong"
}
