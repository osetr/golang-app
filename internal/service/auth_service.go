package service

import (
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
	GetSignUpInput() signUpInput
	SignUp(signUpInput) (map[string]interface{}, error)
	GetSignInInput() signInInput
	SignIn(signInInput) (map[string]interface{}, error)
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
func (*AuthService) GetSignUpInput() signUpInput {
	return signUpInput{}
}

type signUpInput struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (i signUpInput) Validate() (map[string]interface{}, bool) {
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

func (as *AuthService) SignUp(i signUpInput) (map[string]interface{}, error) {
	if _, valid := i.Validate(); !valid {
		panic("First you need validate input")
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
func (as *AuthService) GetSignInInput() signInInput {
	return signInInput{}
}

type signInInput struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (i signInInput) Validate() (map[string]interface{}, bool) {
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

func (as *AuthService) SignIn(i signInInput) (map[string]interface{}, error) {
	userRepo := as.userRepo
	u, _ := userRepo.SignInUser(i.Email, auth.GeneratePasswordHash(i.Password))
	token, err := auth.NewJWTToken(u.Id, secretKey, tokenTTL)
	return map[string]interface{}{"accessToken": token}, err
}
