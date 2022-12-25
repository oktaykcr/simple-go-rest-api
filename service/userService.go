package service

import (
	"github.com/oktaykcr/simple-go-rest-api/domain"
	"github.com/oktaykcr/simple-go-rest-api/errs"
)

type UserService interface {
	CreateUser(name string, age int) (bool, *errs.AppError)
	UpdateUser(id int64, name string, age int) (bool, *errs.AppError)
	GetAllUser() ([]domain.User, *errs.AppError)
	GetUser(id int64) (domain.User, *errs.AppError)
	DeleteUser(id int64) (bool, *errs.AppError)
}

type DefaultUserService struct {
	repository domain.UserRepository
}

func (service DefaultUserService) CreateUser(name string, age int) (bool, *errs.AppError) {
	return service.repository.Create(&domain.User{
		Name: name,
		Age:  age,
	})
}

func (service DefaultUserService) UpdateUser(id int64, name string, age int) (bool, *errs.AppError) {
	return service.repository.Update(&domain.User{
		Id:   id,
		Name: name,
		Age:  age,
	})
}

func (service DefaultUserService) GetAllUser() ([]domain.User, *errs.AppError) {
	return service.repository.GetAll()
}

func (service DefaultUserService) GetUser(id int64) (domain.User, *errs.AppError) {
	return service.repository.Get(id)
}

func (service DefaultUserService) DeleteUser(id int64) (bool, *errs.AppError) {
	return service.repository.Delete(id)
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository: repository}
}
