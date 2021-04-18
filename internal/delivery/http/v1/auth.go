package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/osetr/app/internal/dao"
	"github.com/osetr/app/internal/service"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	signInService := service.SignInService{}

	if err := json.NewDecoder(r.Body).Decode(&signInService.Input); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if valid, message := signInService.Validate(); !valid {
		result, _ := json.Marshal(map[string]interface{}{"detail": message})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if res, err := signInService.Execute(); err != nil {
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
		w.WriteHeader(http.StatusCreated)
		w.Write(result)
	}
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	signUpService := service.SignUpService{}

	if err := json.NewDecoder(r.Body).Decode(&signUpService.Input); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if valid, message := signUpService.Validate(); !valid {
		result, _ := json.Marshal(map[string]interface{}{"detail": message})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if res, err := signUpService.Execute(); err != nil {
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
