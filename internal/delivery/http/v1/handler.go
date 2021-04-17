package v1

import (
	"github.com/gorilla/mux"
)

type Handler struct {
}

func (h *Handler) InitRoute() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/sign-in", h.signIn).Methods("POST")
	router.HandleFunc("/api/v1/sign-up", h.signUp).Methods("POST")

	router.HandleFunc("/api/v1/posts", h.createPost).Methods("POST")
	router.HandleFunc("/api/v1/posts", h.getAllPosts).Methods("GET")
	router.HandleFunc("/api/v1/posts/{id}", h.getPostById).Methods("GET")
	router.HandleFunc("/api/v1/posts/{id}", h.updatePost).Methods("PUT")
	router.HandleFunc("/api/v1/posts/{id}", h.deletePost).Methods("DELETE")

	return router
}

func NewHandler() *Handler {
	return &Handler{}
}
