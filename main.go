package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/leandroandrade/posts-api-mysql/posts/boundary"
	"github.com/leandroandrade/posts-api-mysql/handler"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	sub := router.PathPrefix("/resources").Subrouter()
	sub.Handle("/posts", handler.AppHandler(boundary.GetPosts)).Methods("GET")
	sub.Handle("/posts", handler.AppHandler(boundary.CreatePosts)).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))

}


