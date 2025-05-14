package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func NewResponse(code int, body []byte) Response {
	ans := Response{StatusCode: code}
	if bytes.HasPrefix(bytes.TrimSpace(body), []byte("{")) {
		json.Unmarshal(body, &ans)
	}
	return ans
}

type Response struct {
	StatusCode int     `json:"-"`
	Errors     []Error `json:"_errors"`
	RequestId  string  `json:"_request_id"`
}

func (e Response) Failed() bool {
	return len(e.Errors) > 0
}

func (e Response) Error() string {
	var buf strings.Builder

	buf.WriteString(fmt.Sprintf("[HTTP %d]", e.StatusCode))

	for i := 0; i < len(e.Errors); i++ {
		buf.WriteString(" ")
		if i > 0 {
			buf.WriteString("|| ")
		}

		buf.WriteString(e.Errors[i].Error())
	}

	return buf.String()
}
