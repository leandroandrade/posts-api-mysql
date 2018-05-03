package handler

import (
	"net/http"
	"encoding/json"
)

type AppHandler func(http.ResponseWriter, *http.Request) *AppError

func (fn AppHandler) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	if err := fn(writer, r); err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(err.Code)

		json.NewEncoder(writer).Encode(err)
	}
}

type AppError struct {
	Error   error  `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}
