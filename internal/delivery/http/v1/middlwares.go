package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/osetr/app/pkg/auth"
)

func AuthMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		// TODO: secret should be hided
		id, err := auth.ParseToken(tokenString, "secret")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		r.Header.Set("userId", fmt.Sprint(id))
		next.ServeHTTP(w, r)
	})
}
