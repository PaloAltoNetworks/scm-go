package api

import (
	"context"
)

type Logger func(context.Context, string, ...map[string]interface{})
