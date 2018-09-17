package main

import (
	"github.com/gorilla/mux"
	"github.com/leandroandrade/posts-api-mysql/posts/boundary"
	"github.com/urfave/negroni"
	"github.com/phyber/negroni-gzip/gzip"
	"net/http"
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/posts/service"
	login "github.com/leandroandrade/posts-api-mysql/authentication/boundary"
)

const PathPrefix = "/resources"

func main() {
	postResources := boundary.NewPostHandler(service.NewService())
	loginResources := login.NewLoginResources()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home)
	router.HandleFunc("/token-auth", loginResources.Login).Methods("POST")

	resources := router.PathPrefix(PathPrefix).Subrouter()

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
	negr.Run("")

	http.ListenAndServe(":3000", negr)
}

func Home(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(writer, "Leandro Post API")

}
