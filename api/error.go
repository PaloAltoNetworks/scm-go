package api

import (
	"errors"
	"fmt"
	"strings"
)

var ObjectNotFoundError = errors.New("object not found")
var TooManyRetriesError = errors.New("too many retries")
var UnknownHostError = errors.New("unknown host: cannot determine URI")

type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func (s Error) Error() string {
	var buf strings.Builder

	if s.Code != "" {
		buf.WriteString(s.Code)
	}

	if s.Message != "" {
		if buf.String() != "" {
			buf.WriteString(" ")
		}
		buf.WriteString(s.Message)
	}

	if s.Details != nil {
		if buf.String() != "" {
			buf.WriteString(" - ")
		}
		buf.WriteString(fmt.Sprintf("%v", s.Details))
	}

	return buf.String()
}
