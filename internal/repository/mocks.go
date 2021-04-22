package repository

import (
	"github.com/osetr/app/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockPostRepostitory struct {
	mock.Mock
}

func (mock *MockPostRepostitory) Save(*domain.Post) (*domain.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.Post), args.Error(1)
}

func (mock *MockPostRepostitory) GetSinglePost(int) (*domain.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.Post), args.Error(1)
}

func (mock *MockPostRepostitory) GetAllPosts() ([]domain.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]domain.Post), args.Error(1)
}

func (mock *MockPostRepostitory) UpdatePost(*domain.Post) (*domain.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.Post), args.Error(1)
}

func (mock *MockPostRepostitory) DeletePost(int) error {
	args := mock.Called()
	return args.Error(1)
}

type MockUserRepostitory struct {
	mock.Mock
}

func (mock *MockUserRepostitory) Save(*domain.User) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}

func (mock *MockUserRepostitory) GetSingleUser(int) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}

func (mock *MockUserRepostitory) GetAllUsers() ([]domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]domain.User), args.Error(1)
}

func (mock *MockUserRepostitory) UpdateUser(*domain.User) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}

func (mock *MockUserRepostitory) DeleteUser(int) error {
	args := mock.Called()
	return args.Error(1)
}

func (mock *MockUserRepostitory) SignInUser(email, password string) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}
