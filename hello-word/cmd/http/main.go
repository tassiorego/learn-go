package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

const baseURL = "https://jsonplaceholder.typicode.com/posts/"

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func getPosts(postId string) (*Post, error) {
	response, err := http.Get(baseURL + postId)

	if err != nil {
		return nil, errors.New("failed to fetch post")
	}

	var post Post
	jsonDecoder := json.NewDecoder(response.Body)
	err = jsonDecoder.Decode(&post)

	if err != nil {
		return nil, errors.New("failed to decode JSON response")
	}

	return &post, nil
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("GET /hello/{name}", func(response http.ResponseWriter, request *http.Request) {
		name := request.PathValue("name")
		response.Write([]byte("Hello, " + name + "!"))
	})

	server.HandleFunc("GET /", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("Page not found!"))
	})

	server.HandleFunc("GET /posts/{id}", func(w http.ResponseWriter, request *http.Request) {
		postId := request.PathValue("id")
		post, err := getPosts(postId)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(post)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to encode JSON response"))
		}

	})

	http.ListenAndServe(":8080", server)

}
