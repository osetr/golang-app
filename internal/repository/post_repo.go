package repository

import (
	"github.com/osetr/app/internal/dao"
	"github.com/osetr/app/internal/domain"
)

type PostRepository struct {
}

func (*PostRepository) Save(p *domain.Post) (*domain.Post, error) {
	dao := dao.NewDAO()
	return dao.PostDAO.Save(p, dao.DB)
}

func (*PostRepository) GetSinglePost(id int) (*domain.Post, error) {
	dao := dao.NewDAO()
	return dao.PostDAO.GetSinglePost(id, dao.DB)
}

func (*PostRepository) GetAllPosts() ([]domain.Post, error) {
	dao := dao.NewDAO()
	return dao.PostDAO.GetAllPosts(dao.DB)
}

func (*PostRepository) UpdatePost(p *domain.Post) (*domain.Post, error) {
	dao := dao.NewDAO()
	return dao.PostDAO.UpdatePost(p, dao.DB)
}

func (*PostRepository) DeletePost(id int) error {
	dao := dao.NewDAO()
	return dao.PostDAO.DeletePost(id, dao.DB)
}

func NewPostRepository() IPostRepository {
	return &PostRepository{}
}
