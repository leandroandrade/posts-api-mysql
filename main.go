package main

import (
	"github.com/gorilla/mux"
	"github.com/leandroandrade/posts-api-mysql/posts/boundary"
	"github.com/urfave/negroni"
	"github.com/phyber/negroni-gzip/gzip"
	"net/http"
	"fmt"
)

func main() {
	negr := negroni.Classic()
	negr.Use(gzip.Gzip(gzip.BestSpeed))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home)

	sub := router.PathPrefix("/resources").Subrouter()
	sub.Path("/posts").Queries("size", "{size}", "page", "{page}").
		HandlerFunc(boundary.FindPostsPagination).
		Methods("GET")

	sub.HandleFunc("/posts", boundary.CreatePosts).Methods("POST")
	sub.HandleFunc("/posts/{id:[0-9]+}", boundary.DeletePost).Methods("DELETE")
	sub.HandleFunc("/posts/{id:[0-9]+}", boundary.UpdatePost).Methods("PUT")
	sub.HandleFunc("/posts/{id:[0-9]+}", boundary.GetPostByID).Methods("GET")

	negr.UseHandler(router)
	negr.Run(":3000")

}

func Home(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(writer, "Leandro Post API")

}
