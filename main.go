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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home)

	resources := router.PathPrefix("/resources").Subrouter()

	posts := resources.PathPrefix("/posts").Subrouter()
	posts.Methods("GET").Queries("size", "{size}", "page", "{page}").
		HandlerFunc(boundary.FindPostsPagination)
	posts.Methods("POST").HandlerFunc(boundary.CreatePosts)

	post := resources.PathPrefix("/posts/{id}").Subrouter()
	post.Methods("PUT").HandlerFunc(boundary.UpdatePost)
	post.Methods("DELETE").HandlerFunc(boundary.DeletePost)
	post.Methods("GET").HandlerFunc(boundary.GetPostByID)

	negr := negroni.Classic()
	negr.Use(gzip.Gzip(gzip.BestSpeed))

	negr.UseHandler(router)
	negr.Run(":3000")
}

func Home(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(writer, "Leandro Post API")

}
