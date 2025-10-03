package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

type transport struct {
	transport http.RoundTripper
	client    Client
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	// Input.
	if t.client.LoggingIsSetTo(LogDetailed) {
		reqData, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			t.client.Log(ctx, LogDetailed, fmt.Sprintf(logReqMsg, prettyPrintJsonLines(reqData)))
		} else {
			t.client.Log(ctx, "", fmt.Sprintf("[ERROR] API Request error: %#v", err))
		}
	}

	resp, err := t.transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	// Output.
	if t.client.LoggingIsSetTo(LogDetailed) {
		respData, err := httputil.DumpResponse(resp, true)
		if err == nil {
			t.client.Log(ctx, LogDetailed, fmt.Sprintf(logRespMsg, prettyPrintJsonLines(respData)))
		} else {
			t.client.Log(ctx, "", fmt.Sprintf("[ERROR] API Response error: %#v", err))
		}
	}

	return resp, nil
}

// NewTransport creates a wrapper around a *http.RoundTripper,
// designed to be used for the `Transport` field of http.Client.
//
// This logs each pair of HTTP request/response that it handles.
// The logging is done via Go standard library `log` package.
//
// Deprecated: This will log the content of every http request/response
// at `[DEBUG]` level, without any filtering. Any sensitive information
// will appear as-is in your logs. Please use NewSubsystemLoggingHTTPTransport instead.
func NewTransport(t http.RoundTripper, client Client) *transport {
	return &transport{t, client}
}

// prettyPrintJsonLines iterates through a []byte line-by-line,
// transforming any lines that are complete json into pretty-printed json.
func prettyPrintJsonLines(b []byte) string {
	parts := strings.Split(string(b), "\n")
	for i, p := range parts {
		if b := []byte(p); json.Valid(b) {
			var out bytes.Buffer
			_ = json.Indent(&out, b, "", " ") // already checked for validity
			parts[i] = out.String()
		}
	}
	return strings.Join(parts, "\n")
}

const logReqMsg = `
---[ REQUEST ]---------------------------------------
%s
-----------------------------------------------------`

const logRespMsg = `
---[ RESPONSE ]--------------------------------------
%s
-----------------------------------------------------`
