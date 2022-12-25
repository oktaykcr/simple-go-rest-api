package domain

import (
	"github.com/go-pg/pg/v10"
	"github.com/oktaykcr/simple-go-rest-api/errs"
	"github.com/oktaykcr/simple-go-rest-api/logger"
)

type UserRepositoryImpl struct {
	client *pg.DB
}

func (u UserRepositoryImpl) Create(user *User) (bool, *errs.AppError) {
	result, err := u.client.Model(user).Insert()
	if err != nil {
		logger.Error(err.Error())
	}
	return result.RowsAffected() != -1, nil
}

func (u UserRepositoryImpl) Update(user *User) (bool, *errs.AppError) {
	result, err := u.client.Model(user).WherePK().Update()
	if err != nil {
		logger.Error(err.Error())
	}
	return result.RowsAffected() != -1, nil
}

func (u UserRepositoryImpl) GetAll() ([]User, *errs.AppError) {
	var users []User
	err := u.client.Model(&users).Select()
	if err != nil {
		logger.Error(err.Error())
	}
	return users, nil
}

func (u UserRepositoryImpl) Get(id int64) (User, *errs.AppError) {
	user := User{Id: id}
	err := u.client.Model(&user).WherePK().Select()
	if err != nil {
		logger.Error(err.Error())
	}
	return user, nil
}

func (u UserRepositoryImpl) Delete(id int64) (bool, *errs.AppError) {
	user := User{Id: id}
	result, err := u.client.Model(&user).WherePK().Delete()
	if err != nil {
		logger.Error(err.Error())
	}
	return result.RowsAffected() > 0, nil
}

func (u UserRepositoryImpl) DeleteByName(name string) (bool, *errs.AppError) {
	user := User{Name: name}
	result, err := u.client.Model(&user).Where("Name = ?", user.Name).Delete()
	if err != nil {
		logger.Error(err.Error())
	}
	return result.RowsAffected() > 0, nil
}

func NewUserRepositoryImpl(client *pg.DB) UserRepositoryImpl {
	return UserRepositoryImpl{client: client}
}
