package http

import (
	"testing"

	"github.com/aymendirar/go-backend/src/http"
	"github.com/aymendirar/go-backend/test/util"
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
