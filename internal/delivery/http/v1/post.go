package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/osetr/app/internal/service"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postCreateService := service.PostCreateService{}

	if err := json.NewDecoder(r.Body).Decode(&postCreateService.Input); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if valid, message := postCreateService.Validate(); !valid {
		result, _ := json.Marshal(map[string]interface{}{"detail": message})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if res, err := postCreateService.Execute(); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err.Error()})
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(result)
		return
	} else {
		result, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(result)
	}
}

func (h *Handler) getAllPosts(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) getPostById(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
}
