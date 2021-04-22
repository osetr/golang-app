package service

import (
	"github.com/osetr/app/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockPostService struct {
	mock.Mock
}

func (mock *MockPostService) PostCreate(PostCreateInput) (*domain.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.Post), args.Error(1)
}

func (mock *MockPostService) GetPostCreateInput() PostCreateInput {
	args := mock.Called()
	result := args.Get(0)
	return result.(PostCreateInput)
}

func (mock *MockPostService) PostList() ([]domain.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]domain.Post), args.Error(1)
}

func (mock *MockPostService) PostGet(int) (*domain.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.Post), args.Error(1)
}

func (mock *MockPostService) GetPostUpdateInput() PostUpdateInput {
	args := mock.Called()
	result := args.Get(0)
	return result.(PostUpdateInput)
}

func (mock *MockPostService) PostUpdate(int, PostUpdateInput) (*domain.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.Post), args.Error(1)
}

func (mock *MockPostService) PostDelete(int) error {
	args := mock.Called()
	return args.Error(1)
}

type MockAuthService struct {
	mock.Mock
}

func (mock *MockAuthService) GetSignUpInput() PostUpdateInput {
	args := mock.Called()
	result := args.Get(0)
	return result.(PostUpdateInput)
}

func (mock *MockAuthService) SignUp(SignUpInput) (map[string]interface{}, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(map[string]interface{}), args.Error(1)
}

func (mock *MockAuthService) GetSignInInput() PostUpdateInput {
	args := mock.Called()
	result := args.Get(0)
	return result.(PostUpdateInput)
}

func (mock *MockAuthService) SignIn(SignInInput) (map[string]interface{}, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(map[string]interface{}), args.Error(1)
}
