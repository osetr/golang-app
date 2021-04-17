package dao

import (
	"log"

	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/pkg/database"
)

type userDAO struct {
}

func (*userDAO) Save(p *domain.User) (*domain.User, error) {
	db, err := new(database.СonnectionFactory).GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreateUserTable(db)
	}

	_, insertError := db.Model(p).Insert()
	if insertError != nil {
		log.Fatalf("Error occured while inserting data: %v", insertError)
		return p, insertError
	}

	return p, nil
}

func (*userDAO) GetSingleUser(id int) (*domain.User, error) {
	db, err := new(database.СonnectionFactory).GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreateUserTable(db)
	}

	user := &domain.User{Id: id}
	err = db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}

	return user, nil
}

func (*userDAO) GetAllUsers() ([]domain.User, error) {
	db, err := new(database.СonnectionFactory).GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreateUserTable(db)
	}

	var users []domain.User
	err = db.Model(&users).Select()
	if err != nil {
		panic(err)
	}

	return users, nil
}

func (*userDAO) UpdateUser(p *domain.User) (*domain.User, error) {
	db, err := new(database.СonnectionFactory).GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreateUserTable(db)
	}

	_, err = db.Model(p).WherePK().Update()
	if err != nil {
		panic(err)
	}

	return p, nil
}

func (*userDAO) DeleteUser(id int) error {
	db, err := new(database.СonnectionFactory).GetConnection()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreateUserTable(db)
	}

	user := &domain.User{Id: id}
	_, err = db.Model(user).WherePK().Delete()
	if err != nil {
		panic(err)
	}

	return nil
}

func NewUserDAO() IUserDAO {
	return &userDAO{}
}
