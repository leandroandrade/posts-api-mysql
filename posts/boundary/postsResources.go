package boundary

import (
	"net/http"
	"github.com/leandroandrade/posts-api-mysql/posts/service"
	"encoding/json"
)

func GetPosts(writer http.ResponseWriter, request *http.Request) {
	posts := service.FindAll()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(posts)
	writer.WriteHeader(http.StatusOK)
}

func CreatePosts(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("CREATE POSTS"))
}
