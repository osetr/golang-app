package dao

import (
	"log"

	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/pkg/database"
)

type postDAO struct {
}

func (*postDAO) Save(p *domain.Post) (*domain.Post, error) {
	db, err := new(database.СonnectionFactory).GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
	}

	_, insertError := db.Model(p).Insert()
	if insertError != nil {
		log.Fatalf("Error occured while inserting data: %v", insertError)
		return p, insertError
	}

	return p, nil
}

func (*postDAO) GetSinglePost(id int) (*domain.Post, error) {
	db, err := new(database.СonnectionFactory).GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
	}

	post := &domain.Post{Id: id}
	err = db.Model(post).WherePK().Select()
	if err != nil {
		panic(err)
	}

	return post, nil
}

func (*postDAO) GetAllPosts() ([]domain.Post, error) {
	db, err := new(database.СonnectionFactory).GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
	}

	var posts []domain.Post
	err = db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	return posts, nil
}

func (*postDAO) UpdatePost(p *domain.Post) (*domain.Post, error) {
	db, err := new(database.СonnectionFactory).GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
	}

	_, err = db.Model(p).WherePK().Update()
	if err != nil {
		panic(err)
	}

	return p, nil
}

func (*postDAO) DeletePost(id int) error {
	db, err := new(database.СonnectionFactory).GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
	}

	post := &domain.Post{Id: id}
	_, err = db.Model(post).WherePK().Delete()
	if err != nil {
		panic(err)
	}

	return nil
}

func NewPostDAO() IPostDAO {
	return &postDAO{}
}
