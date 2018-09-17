package boundary

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/leandroandrade/posts-api-mysql/posts/service"
	"strings"
	"github.com/leandroandrade/posts-api-mysql/posts/model"
	"encoding/json"
)

func TestCreateEmptyBody(t *testing.T) {
	expected := 400

	request, _ := http.NewRequest("POST", "/resources/posts", strings.NewReader(""))
	writer := httptest.NewRecorder()

	postHandler := NewPostHandler(service.NewServiceMock())
	postHandler.CreatePosts(writer, request)

	if expected != writer.Code {
		t.Errorf("FAIL: CreatePosts code %v, want %v", writer.Code, expected)
	}
}

func TestCreatePostErrosWhenSave(t *testing.T) {
	expected := 201

	request, _ := http.NewRequest("POST", "/resources/posts", strings.NewReader(`{"description":"test description"}`))
	writer := httptest.NewRecorder()

	postHandler := NewPostHandler(service.NewServiceMock())
	postHandler.CreatePosts(writer, request)

	if expected != writer.Code {
		t.Errorf("FAIL: CreatePosts code %v, want %v", writer.Code, expected)
	}

	var post model.Post
	json.NewDecoder(writer.Body).Decode(&post)

	if post.Id == 0 {
		t.Errorf("FAIL:ID %v, want %v", post.Id, 9988)
	}
}
