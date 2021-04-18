package repository

import (
	"github.com/osetr/app/internal/dao"
	"github.com/osetr/app/internal/domain"
)

type UserRepository struct {
}

func (*UserRepository) Save(p *domain.User) (*domain.User, error) {
	dao := dao.NewDAO()
	return dao.UserDAO.Save(p, dao.DB)
}

func (*UserRepository) GetSingleUser(id int) (*domain.User, error) {
	dao := dao.NewDAO()
	return dao.UserDAO.GetSingleUser(id, dao.DB)
}

func (*UserRepository) GetAllUsers() ([]domain.User, error) {
	dao := dao.NewDAO()
	return dao.UserDAO.GetAllUsers(dao.DB)
}

func (*UserRepository) UpdateUser(p *domain.User) (*domain.User, error) {
	dao := dao.NewDAO()
	return dao.UserDAO.UpdateUser(p, dao.DB)
}

func (*UserRepository) DeleteUser(id int) error {
	dao := dao.NewDAO()
	return dao.UserDAO.DeleteUser(id, dao.DB)
}

func (*UserRepository) SignInUser(email, password string) (*domain.User, error) {
	dao := dao.NewDAO()
	return dao.UserDAO.SignInUser(email, password, dao.DB)
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}
