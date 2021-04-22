package service

import (
	"errors"
	"time"

	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/internal/repository"
	"github.com/osetr/app/pkg/auth"
)

const (
	secretKey = "secret"
	tokenTTL  = 12 * time.Hour
)

// Auth service implementation
type IAuthService interface {
	GetSignUpInput() SignUpInput
	SignUp(SignUpInput) (map[string]interface{}, error)
	GetSignInInput() SignInInput
	SignIn(SignInInput) (map[string]interface{}, error)
}

type AuthService struct {
	userRepo repository.IUserRepository
}

func NewAuthService(userRepo repository.IUserRepository) IAuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// Sign-up functionality
func (*AuthService) GetSignUpInput() SignUpInput {
	return SignUpInput{}
}

type SignUpInput struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (i SignUpInput) Validate() (map[string]interface{}, bool) {
	message := make(map[string]interface{})
	valid := true

	// TODO: make valid validation
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

	return message, valid
}

func (as *AuthService) SignUp(i SignUpInput) (map[string]interface{}, error) {
	if _, valid := i.Validate(); !valid {
		return nil, errors.New("first you need validate input")
	}

	postRepo := as.userRepo
	p, err := postRepo.Save(&domain.User{
		Name:     i.Name,
		Email:    i.Email,
		Password: auth.GeneratePasswordHash(i.Password),
	})

	return map[string]interface{}{"id": p.Id, "name": p.Name, "email": p.Email}, err
}

// Sign-in functionality
func (as *AuthService) GetSignInInput() SignInInput {
	return SignInInput{}
}

type SignInInput struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (i SignInInput) Validate() (map[string]interface{}, bool) {
	message := make(map[string]interface{})
	valid := true

	// TODO: make valid validation
	if ul := len(i.Email); ul >= 36 || ul <= 2 {
		message["email"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}

	if ul := len(i.Password); ul >= 36 || ul <= 2 {
		message["password"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}

	return message, valid
}

func (as *AuthService) SignIn(i SignInInput) (map[string]interface{}, error) {
	if _, valid := i.Validate(); !valid {
		return nil, errors.New("first you need validate input")
	}

	userRepo := as.userRepo
	u, _ := userRepo.SignInUser(i.Email, auth.GeneratePasswordHash(i.Password))
	token, err := auth.NewJWTToken(u.Id, secretKey, tokenTTL)
	return map[string]interface{}{"accessToken": token}, err
}
