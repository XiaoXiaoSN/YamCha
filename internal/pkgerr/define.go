package pkgerr

import (
	"bytes"
	"fmt"
)

type _error struct {
	StatusCode int
	ErrorCode  string
	Message    string
}

func (e *_error) Error() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "[%s] %s", e.ErrorCode, e.Message)
	return buf.String()
}

// Error Define
var (
	ErrBadRequest = &_error{StatusCode: 400, ErrorCode: "400000", Message: "Invailed Input"}
	ErrNotFound   = &_error{StatusCode: 404, ErrorCode: "404000", Message: "Resource Not Found"}
)
