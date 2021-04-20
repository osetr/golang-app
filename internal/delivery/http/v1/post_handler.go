package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/osetr/app/internal/service"
	"github.com/osetr/app/pkg/adapter"
)

type IPostHandler interface {
	createPost(w http.ResponseWriter, r *http.Request)
	getAllPosts(w http.ResponseWriter, r *http.Request)
	getPostById(w http.ResponseWriter, r *http.Request)
	updatePost(w http.ResponseWriter, r *http.Request)
	deletePost(w http.ResponseWriter, r *http.Request)
}

type PostHandler struct {
	postService service.IPostService
}

func NewPostHandler(postService service.IPostService) IPostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (h *PostHandler) createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postService := h.postService
	postCreateInput := postService.GetPostCreateInput()

	if err := json.NewDecoder(r.Body).Decode(&postCreateInput); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if message, valid := postCreateInput.Validate(); !valid {
		result, _ := json.Marshal(map[string]interface{}{"detail": message})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if res, err := postService.PostCreate(postCreateInput); err != nil {
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

func (h *PostHandler) getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postService := h.postService

	if res, err := postService.PostList(); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err.Error()})
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(result)
		return
	} else {
		result, err := json.Marshal(map[string]interface{}{
			"count": len(res),
			"data":  res,
		})
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func (h *PostHandler) getPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postService := h.postService
	id_param := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": "bad id param"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}
	if res, err := postService.PostGet(int(id)); err != nil {
		httpAdapter := adapter.NewHttpExcAdapter()
		_, stMess, stCode := httpAdapter.Transform(err)
		result, _ := json.Marshal(map[string]interface{}{"detail": stMess})
		w.WriteHeader(stCode)
		w.Write(result)
		return
	} else {
		result, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func (h *PostHandler) updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postService := h.postService
	postUpdateInput := postService.GetPostUpdateInput()
	id_param := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": "bad id param"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&postUpdateInput); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if message, valid := postUpdateInput.Validate(); !valid {
		result, _ := json.Marshal(map[string]interface{}{"detail": message})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if res, err := postService.PostUpdate(int(id), postUpdateInput); err != nil {
		httpAdapter := adapter.NewHttpExcAdapter()
		_, stMess, stCode := httpAdapter.Transform(err)
		result, _ := json.Marshal(map[string]interface{}{"detail": stMess})
		w.WriteHeader(stCode)
		w.Write(result)
		return
	} else {
		result, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func (h *PostHandler) deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postService := h.postService
	id_param := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": "bad id param"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}
	if err := postService.PostDelete(int(id)); err != nil {
		httpAdapter := adapter.NewHttpExcAdapter()
		_, stMess, stCode := httpAdapter.Transform(err)
		result, _ := json.Marshal(map[string]interface{}{"detail": stMess})
		w.WriteHeader(stCode)
		w.Write(result)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
