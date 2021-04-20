package repository

import (
	"github.com/osetr/app/internal/dao"
	"github.com/osetr/app/internal/domain"
)

type IUserRepository interface {
	Save(*domain.User) (*domain.User, error)
	GetSingleUser(int) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(*domain.User) (*domain.User, error)
	DeleteUser(int) error
	SignInUser(string, string) (*domain.User, error)
}

type UserRepository struct {
	userDAO dao.IUserDAO
}

func NewUserRepository(userDAO dao.IUserDAO) IUserRepository {
	return &UserRepository{
		userDAO: userDAO,
	}
}

func (us *UserRepository) Save(p *domain.User) (*domain.User, error) {
	return us.userDAO.Save(p)
}

func (us *UserRepository) GetSingleUser(id int) (*domain.User, error) {
	return us.userDAO.GetSingleUser(id)
}

func (us *UserRepository) GetAllUsers() ([]domain.User, error) {
	return us.userDAO.GetAllUsers()
}

func (us *UserRepository) UpdateUser(p *domain.User) (*domain.User, error) {
	return us.userDAO.UpdateUser(p)
}

func (us *UserRepository) DeleteUser(id int) error {
	return us.userDAO.DeleteUser(id)
}

func (us *UserRepository) SignInUser(email, password string) (*domain.User, error) {
	return us.userDAO.SignInUser(email, password)
}
