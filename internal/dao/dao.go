package dao

import "github.com/osetr/app/internal/domain"

type IAuthorizationDAO interface {
}

type IPostDAO interface {
	Save(*domain.Post) (*domain.Post, error)
	GetSinglePost(id int) (*domain.Post, error)
	GetAllPosts() ([]domain.Post, error)
	UpdatePost(*domain.Post) (*domain.Post, error)
	DeletePost(id int) error
}

type DAO struct {
	IAuthorizationDAO
	IPostDAO
}

func NewDAO() *DAO {
	return &DAO{}
}
