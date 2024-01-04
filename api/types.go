package api

func String(v string) *string {
	return &v
}

func Int(v int64) *int64 {
	return &v
}

func Float(v float64) *float64 {
	return &v
}

func Bool(v bool) *bool {
	return &v
}
