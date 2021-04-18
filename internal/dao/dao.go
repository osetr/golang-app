package dao

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/pkg/adapter"
	"github.com/osetr/app/pkg/database"
)

type IPostDAO interface {
	Save(*domain.Post, *pg.DB) (*domain.Post, error)
	GetSinglePost(int, *pg.DB) (*domain.Post, error)
	GetAllPosts(*pg.DB) ([]domain.Post, error)
	UpdatePost(*domain.Post, *pg.DB) (*domain.Post, error)
	DeletePost(int, *pg.DB) error
}

type IUserDAO interface {
	Save(*domain.User, *pg.DB) (*domain.User, error)
	GetSingleUser(int, *pg.DB) (*domain.User, error)
	GetAllUsers(*pg.DB) ([]domain.User, error)
	UpdateUser(*domain.User, *pg.DB) (*domain.User, error)
	DeleteUser(int, *pg.DB) error
	SignInUser(string, string, *pg.DB) (*domain.User, error)
}

type IHttpExcAdapter interface {
	Transform(error) (error, string, int)
}

type DAO struct {
	PostDAO        IPostDAO
	UserDAO        IUserDAO
	HttpExcAdapter IHttpExcAdapter
	DB             *pg.DB
}

func NewDAO() *DAO {
	return &DAO{
		PostDAO:        NewPostDAO(),
		UserDAO:        NewUserDAO(),
		HttpExcAdapter: adapter.NewHttpExcAdapter(),
		DB: func() *pg.DB {
			db, err := new(database.СonnectionFactory).GetConnection()
			if err != nil {
				log.Fatalf("Error occured while connecting to database: %v", err)
			} else {
				domain.CreatePostTable(db)
				domain.CreateUserTable(db)
			}
			return db
		}(),
	}
}
