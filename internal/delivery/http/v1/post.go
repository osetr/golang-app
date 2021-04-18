package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/osetr/app/internal/dao"
	"github.com/osetr/app/internal/service"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postCreateService := service.PostCreateService{}

	if err := json.NewDecoder(r.Body).Decode(&postCreateService.Input); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err.Error()})
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
	w.Header().Set("Content-Type", "application/json")
	postListService := service.PostListService{}

	if res, err := postListService.Execute(); err != nil {
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

func (h *Handler) getPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postGetService := service.PostGetService{}
	id_param := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": "bad id param"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}
	if res, err := postGetService.Execute(int(id)); err != nil {
		httpAdapter := dao.NewDAO().HttpExcAdapter
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

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postUpdateService := service.PostUpdateService{}
	id_param := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": "bad id param"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&postUpdateService.Input); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if valid, message := postUpdateService.Validate(); !valid {
		result, _ := json.Marshal(map[string]interface{}{"detail": message})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if res, err := postUpdateService.Execute(int(id)); err != nil {
		httpAdapter := dao.NewDAO().HttpExcAdapter
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

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postDeleteService := service.PostDeleteService{}
	id_param := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": "bad id param"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}
	if err := postDeleteService.Execute(int(id)); err != nil {
		httpAdapter := dao.NewDAO().HttpExcAdapter
		_, stMess, stCode := httpAdapter.Transform(err)
		result, _ := json.Marshal(map[string]interface{}{"detail": stMess})
		w.WriteHeader(stCode)
		w.Write(result)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
