package response

import (
	"net/http"
	"encoding/json"
)

type Message struct {
	Code             int    `json:"code,omitempty"`
	MessageUser      string `json:"messageUser,omitempty"`
	MessageDeveloper string `json:"messageDeveloper,omitempty"`
}

func JSONErr(writer http.ResponseWriter, message Message) {
	write(writer, message.Code, message)
}

func JSON(writer http.ResponseWriter, statuscode int, v interface{}) {
	write(writer, statuscode, v)
}

func write(writer http.ResponseWriter, statucode int, v interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statucode)

	json.NewEncoder(writer).Encode(v)
}
