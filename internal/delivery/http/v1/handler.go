package v1

import (
	"github.com/gorilla/mux"
)

type Handler struct {
}

func (h *Handler) InitRoute() *mux.Router {
	router := mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/sign-in", h.signIn).Methods("POST")
	v1.HandleFunc("/sign-up", h.signUp).Methods("POST")

	posts := v1.PathPrefix("/posts").Subrouter()
	posts.Use(AuthMiddlware)

	posts.HandleFunc("", h.createPost).Methods("POST")
	posts.HandleFunc("", h.getAllPosts).Methods("GET")
	posts.HandleFunc("/{id}", h.getPostById).Methods("GET")
	posts.HandleFunc("/{id}", h.updatePost).Methods("PUT")
	posts.HandleFunc("/{id}", h.deletePost).Methods("DELETE")

	return router
}

func NewHandler() *Handler {
	return &Handler{}
}
