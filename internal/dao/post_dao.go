package dao

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/pkg/database"
)

type IPostDAO interface {
	Save(*domain.Post) (*domain.Post, error)
	GetSinglePost(int) (*domain.Post, error)
	GetAllPosts() ([]domain.Post, error)
	UpdatePost(*domain.Post) (*domain.Post, error)
	DeletePost(int) error
}

type PostDAO struct {
	conn database.IСonnectionFactory
}

func NewPostDAO(conn database.IСonnectionFactory) IPostDAO {
	return &PostDAO{
		conn: conn,
	}
}

func (pdao *PostDAO) Save(p *domain.Post) (*domain.Post, error) {
	db, err := pdao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

	_, err = db.Model(p).Insert()
	if err != nil {
		return &domain.Post{}, err
	}

	return p, nil
}

func (pdao *PostDAO) GetSinglePost(id int) (*domain.Post, error) {
	db, err := pdao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

	post := &domain.Post{Id: id}
	err = db.Model(post).WherePK().Select()
	if err != nil {
		return &domain.Post{}, err
	}

	return post, nil
}

func (pdao *PostDAO) GetAllPosts() ([]domain.Post, error) {
	db, err := pdao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

	var posts []domain.Post
	err = db.Model(&posts).Select()
	if err != nil {
		return []domain.Post{}, err
	}

	return posts, nil
}

func (pdao *PostDAO) UpdatePost(p *domain.Post) (*domain.Post, error) {
	db, err := pdao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

	res, err := db.Model(p).WherePK().Update()
	if res.RowsAffected() == 0 {
		return &domain.Post{}, pg.ErrNoRows
	}
	if err != nil {
		return &domain.Post{}, err
	}

	return p, nil
}

func (pdao *PostDAO) DeletePost(id int) error {
	db, err := pdao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

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
