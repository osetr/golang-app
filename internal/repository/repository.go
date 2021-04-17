package repository

import "github.com/osetr/app/internal/domain"

type IPostRepository interface {
	Save(*domain.Post) (*domain.Post, error)
	GetSinglePost(id int) (*domain.Post, error)
	GetAllPosts() ([]domain.Post, error)
	UpdatePost(*domain.Post) (*domain.Post, error)
	DeletePost(id int) error
}

type Repository struct {
	PostRepository IPostRepository
}

func NewRepository() *Repository {
	return &Repository{
		PostRepository: NewPostRepository(),
	}
}
