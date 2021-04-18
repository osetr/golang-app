package repository

import "github.com/osetr/app/internal/domain"

type IPostRepository interface {
	Save(*domain.Post) (*domain.Post, error)
	GetSinglePost(int) (*domain.Post, error)
	GetAllPosts() ([]domain.Post, error)
	UpdatePost(*domain.Post) (*domain.Post, error)
	DeletePost(int) error
}

type IUserRepository interface {
	Save(*domain.User) (*domain.User, error)
	GetSingleUser(int) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(*domain.User) (*domain.User, error)
	DeleteUser(int) error
	SignInUser(string, string) (*domain.User, error)
}

type Repository struct {
	PostRepository IPostRepository
	UserRepository IUserRepository
}

func NewRepository() *Repository {
	return &Repository{
		PostRepository: NewPostRepository(),
		UserRepository: NewUserRepository(),
	}
}
