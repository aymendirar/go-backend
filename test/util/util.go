package util

import (
	"bytes"
	"encoding/json"
	"io"
)

type TestCase struct {
	Input    any
	Expected any
}

func JSONRequest(payload map[string]any) io.Reader {
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(payload)
	return &buf
}
