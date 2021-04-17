package repository

import (
	"github.com/osetr/app/internal/dao"
	"github.com/osetr/app/internal/domain"
)

type PostRepository struct {
}

func (*PostRepository) Save(p *domain.Post) (*domain.Post, error) {
	return dao.NewPostDAO().Save(p)
}

func (*PostRepository) GetSinglePost(id int) (*domain.Post, error) {
	return dao.NewPostDAO().GetSinglePost(id)
}

func (*PostRepository) GetAllPosts() ([]domain.Post, error) {
	return dao.NewPostDAO().GetAllPosts()
}

func (*PostRepository) UpdatePost(p *domain.Post) (*domain.Post, error) {
	return dao.NewPostDAO().UpdatePost(p)
}

func (*PostRepository) DeletePost(id int) error {
	return dao.NewPostDAO().DeletePost(id)
}

func NewPostRepository() IPostRepository {
	return &PostRepository{}
}
