package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/leandroandrade/posts-api-mysql/posts/boundary"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	sub := router.PathPrefix("/resources").Subrouter()
	sub.HandleFunc("/posts", boundary.GetPosts).Methods("GET")
	sub.HandleFunc("/posts", boundary.CreatePosts).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))

}
