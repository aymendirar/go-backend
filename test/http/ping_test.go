package http

import (
	h "net/http"
	"net/http/httptest"
	"testing"

	"github.com/aymendirar/go-backend/src/http"
	"github.com/aymendirar/go-backend/test/util"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	s := http.NewTestHTTPServer()
	cases := []util.TestCase{
		{
			Input:    "ping",
			Expected: false,
		},
		{
			Input:    "pong",
			Expected: true,
		},
		{
			Input:    "prawn",
			Expected: false,
		},
	}

	for _, c := range cases {
		if result := s.Ping(c.Input.(string)); result != c.Expected.(bool) {
			t.Errorf("failed case: %v", c)
		}
	}
}

func TestPingRoute(t *testing.T) {
	s := http.NewTestHTTPServer()

	w := httptest.NewRecorder()
	req, _ := h.NewRequest(
		"POST", "/api/ping",
		util.JSONRequest(map[string]any{
			"ping": "pong",
		}))
	s.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"response":"ok"}`, w.Body.String())
}
