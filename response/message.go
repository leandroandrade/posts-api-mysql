package response

import (
	"net/http"
	"encoding/json"
)

type Message struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"value,omitempty"`
}

func JSON(writer http.ResponseWriter, message Message) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(message.Code)

	json.NewEncoder(writer).Encode(message)
}
