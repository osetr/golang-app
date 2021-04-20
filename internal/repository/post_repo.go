package repository

import (
	"github.com/osetr/app/internal/dao"
	"github.com/osetr/app/internal/domain"
)

type IPostRepository interface {
	Save(*domain.Post) (*domain.Post, error)
	GetSinglePost(int) (*domain.Post, error)
	GetAllPosts() ([]domain.Post, error)
	UpdatePost(*domain.Post) (*domain.Post, error)
	DeletePost(int) error
}

type PostRepository struct {
	postDAO dao.IPostDAO
}

func NewPostRepository(postDAO dao.IPostDAO) IPostRepository {
	return &PostRepository{
		postDAO: postDAO,
	}
}

func (ps *PostRepository) Save(p *domain.Post) (*domain.Post, error) {
	return ps.postDAO.Save(p)
}

func (ps *PostRepository) GetSinglePost(id int) (*domain.Post, error) {
	return ps.postDAO.GetSinglePost(id)
}

func (ps *PostRepository) GetAllPosts() ([]domain.Post, error) {
	return ps.postDAO.GetAllPosts()
}

func (ps *PostRepository) UpdatePost(p *domain.Post) (*domain.Post, error) {
	return ps.postDAO.UpdatePost(p)
}

func (ps *PostRepository) DeletePost(id int) error {
	return ps.postDAO.DeletePost(id)
}
