package domain

import "github.com/oktaykcr/simple-go-rest-api/errs"

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//go:generate mockgen -source=./user.go -destination=../mocks/domain/mockUserRepository.go -package=domain UserRepository
type UserRepository interface {
	Create(user *User) (bool, *errs.AppError)
	Update(user *User) (bool, *errs.AppError)
	GetAll() ([]User, *errs.AppError)
	Get(id int64) (User, *errs.AppError)
	Delete(id int64) (bool, *errs.AppError)
	DeleteByName(name string) (bool, *errs.AppError)
}
