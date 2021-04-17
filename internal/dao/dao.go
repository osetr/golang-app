package dao

import "github.com/osetr/app/internal/domain"

type IUserDAO interface {
	Save(*domain.User) (*domain.User, error)
	GetSingleUser(id int) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(*domain.User) (*domain.User, error)
	DeleteUser(id int) error
}

type IPostDAO interface {
	Save(*domain.Post) (*domain.Post, error)
	GetSinglePost(id int) (*domain.Post, error)
	GetAllPosts() ([]domain.Post, error)
	UpdatePost(*domain.Post) (*domain.Post, error)
	DeletePost(id int) error
}

type DAO struct {
	postDAO IPostDAO
	userDao IUserDAO
}

func NewDAO() *DAO {
	return &DAO{
		postDAO: NewPostDAO(),
		userDao: NewUserDAO(),
	}
}
