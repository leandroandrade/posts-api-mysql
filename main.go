package main

import (
	"github.com/gorilla/mux"
	"github.com/leandroandrade/posts-api-mysql/posts/boundary"
	"github.com/urfave/negroni"
	"github.com/phyber/negroni-gzip/gzip"
	"net/http"
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/posts/service"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home)

	postResources := boundary.NewPostHandler(service.NewService())

	resources := router.PathPrefix("/resources").Subrouter()

	posts := resources.PathPrefix("/posts").Subrouter()
	posts.Methods("GET").Queries("size", "{size}", "page", "{page}").
		HandlerFunc(postResources.FindPostsPagination)
	posts.Methods("POST").HandlerFunc(postResources.CreatePosts)

	post := resources.PathPrefix("/posts/{id}").Subrouter()
	post.Methods("PUT").HandlerFunc(postResources.UpdatePost)
	post.Methods("DELETE").HandlerFunc(postResources.DeletePost)
	post.Methods("GET").HandlerFunc(postResources.GetPostByID)

	negr := negroni.Classic()
	negr.Use(gzip.Gzip(gzip.BestSpeed))

	negr.UseHandler(router)
	negr.Run(":3000")
}

func Home(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(writer, "Leandro Post API")

}
