package dao

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/pkg/database"
)

type IUserDAO interface {
	Save(*domain.User) (*domain.User, error)
	GetSingleUser(int) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(*domain.User) (*domain.User, error)
	DeleteUser(int) error
	SignInUser(string, string) (*domain.User, error)
}

type UserDAO struct {
	conn database.IСonnectionFactory
}

func NewUserDAO(conn database.IСonnectionFactory) IUserDAO {
	return &UserDAO{
		conn: conn,
	}
}

func (udao *UserDAO) Save(p *domain.User) (*domain.User, error) {
	db, err := udao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

	_, err = db.Model(p).Insert()
	if err != nil {
		return &domain.User{}, err
	}

	return p, nil
}

func (udao *UserDAO) GetSingleUser(id int) (*domain.User, error) {
	db, err := udao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

	user := &domain.User{Id: id}
	err = db.Model(user).WherePK().Select()
	if err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (udao *UserDAO) GetAllUsers() ([]domain.User, error) {
	db, err := udao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

	var users []domain.User
	err = db.Model(&users).Select()
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (udao *UserDAO) UpdateUser(p *domain.User) (*domain.User, error) {
	db, err := udao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

	res, err := db.Model(p).WherePK().Update()
	if res.RowsAffected() == 0 {
		return &domain.User{}, pg.ErrNoRows
	}
	if err != nil {
		return &domain.User{}, err
	}

	return p, nil
}

func (udao *UserDAO) DeleteUser(id int) error {
	db, err := udao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

	user := &domain.User{Id: id}
	res, err := db.Model(user).WherePK().Delete()
	if res.RowsAffected() == 0 {
		return pg.ErrNoRows
	}
	if err != nil {
		return err
	}

	return nil
}

func (udao *UserDAO) SignInUser(email, password string) (*domain.User, error) {
	db, err := udao.conn.GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
		defer db.Close()
	}

	user := &domain.User{}
	err = db.Model(user).
		Where("email = ?", email).
		Where("password = ?", password).
		Select()
	return user, err
}
