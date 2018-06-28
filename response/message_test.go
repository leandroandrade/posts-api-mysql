package response

import (
	"testing"
	"net/http/httptest"
)

type responseMock struct {
	Message string
}

func TestJSONErr(t *testing.T) {
	var testcases = []struct {
		writer  *httptest.ResponseRecorder
		payload Payload
	}{
		{
			writer:  httptest.NewRecorder(),
			payload: Payload{Code: 200, Message: "Success", Detail: "Operation sucessful"},
		},
		{
			writer:  httptest.NewRecorder(),
			payload: Payload{Code: 400, Message: "Request error", Detail: "Cannot complete your request"},
		},
	}

	for _, test := range testcases {
		JSONErr(test.writer, test.payload)
		if test.payload.Code != test.writer.Code {
			t.Errorf("FAIL: JSONErr code %v, want %v", test.writer.Code, test.payload.Code)
		}
	}
}

func TestJSON(t *testing.T) {
	var testcases = []struct {
		writer *httptest.ResponseRecorder
		body   responseMock
		code   int
	}{
		{
			writer: httptest.NewRecorder(),
			body:   responseMock{Message: "Hello"},
			code:   200,
		},
	}

	for _, test := range testcases {
		JSON(test.writer, test.code, test.body)
		if test.code != test.writer.Code {
			t.Errorf("FAIL: JSONErr code %v, want %v", test.writer.Code, test.code)
		}
	}
}
