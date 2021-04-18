package dao

import (
	"github.com/go-pg/pg/v10"
	"github.com/osetr/app/internal/domain"
)

type UserDAO struct {
}

func (*UserDAO) Save(p *domain.User, db *pg.DB) (*domain.User, error) {
	_, err := db.Model(p).Insert()
	if err != nil {
		return &domain.User{}, err
	}

	return p, nil
}

func (*UserDAO) GetSingleUser(id int, db *pg.DB) (*domain.User, error) {
	user := &domain.User{Id: id}
	err := db.Model(user).WherePK().Select()
	if err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (*UserDAO) GetAllUsers(db *pg.DB) ([]domain.User, error) {
	var users []domain.User
	err := db.Model(&users).Select()
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (*UserDAO) UpdateUser(p *domain.User, db *pg.DB) (*domain.User, error) {
	res, err := db.Model(p).WherePK().Update()
	if res.RowsAffected() == 0 {
		return &domain.User{}, pg.ErrNoRows
	}
	if err != nil {
		return &domain.User{}, err
	}

	return p, nil
}

func (*UserDAO) DeleteUser(id int, db *pg.DB) error {
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

func (*UserDAO) SignInUser(email, password string, db *pg.DB) (*domain.User, error) {
	user := &domain.User{}
	err := db.Model(user).
		Where("email = ?", email).
		Where("password = ?", password).
		Select()
	return user, err
}

func NewUserDAO() IUserDAO {
	return &UserDAO{}
}
