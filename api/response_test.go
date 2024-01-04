package api

import (
	"net/http"
	"testing"
)

func TestParseResponse(t *testing.T) {
	data := []byte(`
{
    "_errors":[
        {
            "code":"API_I00035",
            "message":"Invalid Request Payload ",
            "details":["\"folder\" is required"]
        }
    ],
    "_request_id":"167e22a1-6250-42f8-8e36-74b8e52f24bb"
}`)

	ans := NewResponse(http.StatusBadRequest, data)

	if !ans.Failed() {
		t.Fatalf("No errors present")
	}

	if len(ans.Errors) != 1 {
		t.Fatalf("Errors is len %d, not 1", len(ans.Errors))
	}

	if ans.Errors[0].Code != "API_I00035" {
		t.Errorf("Code is %q", ans.Errors[0].Code)
	}
	if ans.Errors[0].Message != "Invalid Request Payload " {
		t.Errorf("Message is %q", ans.Errors[0].Message)
	}
	if ans.Errors[0].Details == nil {
		t.Errorf("Details is nil")
	}
}
