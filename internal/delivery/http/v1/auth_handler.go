package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/osetr/app/internal/service"
	"github.com/osetr/app/pkg/adapter"
)

type IAuthHandler interface {
	signIn(w http.ResponseWriter, r *http.Request)
	signUp(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	authService service.IAuthService
}

func NewAuthHandler(authService service.IAuthService) IAuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) signIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authService := h.authService
	signInInput := authService.GetSignInInput()

	if err := json.NewDecoder(r.Body).Decode(&signInInput); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if message, valid := signInInput.Validate(); !valid {
		result, _ := json.Marshal(map[string]interface{}{"detail": message})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if res, err := authService.SignIn(signInInput); err != nil {
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
		w.WriteHeader(http.StatusCreated)
		w.Write(result)
	}
}

func (h *AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authService := h.authService
	signUpInput := authService.GetSignUpInput()

	if err := json.NewDecoder(r.Body).Decode(&signUpInput); err != nil {
		result, _ := json.Marshal(map[string]interface{}{"detail": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if message, valid := signUpInput.Validate(); !valid {
		result, _ := json.Marshal(map[string]interface{}{"detail": message})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	if res, err := authService.SignUp(signUpInput); err != nil {
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
