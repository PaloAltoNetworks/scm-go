package config_setup

import (
	"encoding/json"
	"time"
)

// normalizeRFC3339 rewrites named string fields in a JSON object from common
// non-RFC3339 datetime formats to RFC3339 so that the standard encoding/json
// time.Time unmarshaller can handle responses from non-conformant backends
// (e.g. backends that serialise time.Time with Go's default String() format).
func normalizeRFC3339(data []byte, fields ...string) []byte {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return data
	}
	fallbackFormats := []string{
		"2006-01-02 15:04:05 +0000 UTC",
		"2006-01-02 15:04:05 +0000 MST",
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02T15:04:05 +0000 UTC",
	}
	modified := false
	for _, field := range fields {
		v, ok := raw[field]
		if !ok {
			continue
		}
		var s string
		if err := json.Unmarshal(v, &s); err != nil {
			continue
		}
		// Already valid RFC3339 — leave it untouched.
		if _, err := time.Parse(time.RFC3339, s); err == nil {
			continue
		}
		for _, layout := range fallbackFormats {
			if t, err := time.Parse(layout, s); err == nil {
				if b, err := json.Marshal(t.UTC().Format(time.RFC3339)); err == nil {
					raw[field] = b
					modified = true
				}
				break
			}
		}
	}
	if !modified {
		return data
	}
	if b, err := json.Marshal(raw); err == nil {
		return b
	}
	return data
}
