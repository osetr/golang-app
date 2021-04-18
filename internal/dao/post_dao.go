package dao

import (
	"github.com/go-pg/pg/v10"
	"github.com/osetr/app/internal/domain"
)

type PostDAO struct {
}

func (*PostDAO) Save(p *domain.Post, db *pg.DB) (*domain.Post, error) {
	_, err := db.Model(p).Insert()
	if err != nil {
		return &domain.Post{}, err
	}

	return p, nil
}

func (*PostDAO) GetSinglePost(id int, db *pg.DB) (*domain.Post, error) {
	post := &domain.Post{Id: id}
	err := db.Model(post).WherePK().Select()
	if err != nil {
		return &domain.Post{}, err
	}

	return post, nil
}

func (*PostDAO) GetAllPosts(db *pg.DB) ([]domain.Post, error) {
	var posts []domain.Post
	err := db.Model(&posts).Select()
	if err != nil {
		return []domain.Post{}, err
	}

	return posts, nil
}

func (*PostDAO) UpdatePost(p *domain.Post, db *pg.DB) (*domain.Post, error) {
	res, err := db.Model(p).WherePK().Update()
	if res.RowsAffected() == 0 {
		return &domain.Post{}, pg.ErrNoRows
	}
	if err != nil {
		return &domain.Post{}, err
	}

	return p, nil
}

func (*PostDAO) DeletePost(id int, db *pg.DB) error {
	post := &domain.Post{Id: id}
	res, err := db.Model(post).WherePK().Delete()
	if res.RowsAffected() == 0 {
		return pg.ErrNoRows
	}
	if err != nil {
		return err
	}

	return nil
}

func NewPostDAO() IPostDAO {
	return &PostDAO{}
}
