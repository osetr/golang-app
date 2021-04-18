package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/internal/repository"
	"github.com/osetr/app/pkg/auth"
)

const (
	signatureKey = "secret"
	tokenTTL     = 12 * time.Hour
)

type SignUpService struct {
	Input struct {
		Name     string `json:"name,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}
}

func (s *SignUpService) Validate() (bool, map[string]interface{}) {
	message := make(map[string]interface{})
	i := s.Input
	valid := true

	if ul := len(i.Name); ul >= 36 || ul <= 2 {
		message["name"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}

	if ul := len(i.Email); ul >= 36 || ul <= 2 {
		message["email"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}

	if ul := len(i.Password); ul >= 36 || ul <= 2 {
		message["password"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}

	return valid, message
}

func (s *SignUpService) Execute() (map[string]interface{}, error) {
	if valid, _ := s.Validate(); !valid {
		panic("First you need validate input")
	}

	postRepo := repository.NewRepository().UserRepository
	p, err := postRepo.Save(&domain.User{
		Name:     s.Input.Name,
		Email:    s.Input.Email,
		Password: auth.GeneratePasswordHash(s.Input.Password),
	})

	return map[string]interface{}{"id": p.Id, "name": p.Name, "email": p.Email}, err
}

type SignInService struct {
	Input struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}
}

func (s *SignInService) Validate() (bool, map[string]interface{}) {
	message := make(map[string]interface{})
	i := s.Input
	valid := true

	if ul := len(i.Email); ul >= 36 || ul <= 2 {
		message["email"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}

	if ul := len(i.Password); ul >= 36 || ul <= 2 {
		message["password"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}

	return valid, message
}

func (s *SignInService) Execute() (map[string]interface{}, error) {
	userRepo := repository.NewRepository().UserRepository
	u, _ := userRepo.SignInUser(s.Input.Email, auth.GeneratePasswordHash(s.Input.Password))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		u.Id,
	})
	token, err := claims.SignedString([]byte(signatureKey))
	return map[string]interface{}{"accessToken": token}, err
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"id"`
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signatureKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("bad claims type")
	}

	return claims.UserId, nil
}
