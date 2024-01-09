package api

import (
	"context"
	"net/url"
)

type Client interface {
	GetHost() string
	LoggingIsSetTo(string) bool
	Log(context.Context, string, string)
	Do(context.Context, string, string, url.Values, interface{}, interface{}, ...error) ([]byte, error)
}
