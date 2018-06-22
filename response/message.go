package response

import (
	"net/http"
	"encoding/json"
)

type Payload struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

func JSONErr(writer http.ResponseWriter, payload Payload) {
	write(writer, payload.Code, payload)
}

func JSON(writer http.ResponseWriter, statuscode int, v interface{}) {
	write(writer, statuscode, v)
}

func write(writer http.ResponseWriter, statucode int, v interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statucode)

	json.NewEncoder(writer).Encode(v)
}
