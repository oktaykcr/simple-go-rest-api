package service

import (
	"github.com/golang/mock/gomock"
	domain2 "github.com/oktaykcr/simple-go-rest-api/domain"
	"github.com/oktaykcr/simple-go-rest-api/mocks/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserShouldReturnTrue(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := domain.NewMockUserRepository(controller)
	mockUserRepository.EXPECT().Create(gomock.Any()).Return(true, nil)

	service := NewUserService(mockUserRepository)
	//Act
	result, appError := service.CreateUser("Bob", 20)
	//Assert
	assert.True(t, result)
	assert.Nil(t, appError)
}

func TestUpdateUserShouldReturnTrue(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := domain.NewMockUserRepository(controller)
	mockUserRepository.EXPECT().Update(gomock.Any()).Return(true, nil)

	service := NewUserService(mockUserRepository)
	//Act
	result, appError := service.UpdateUser(1, "Bob", 20)
	//Assert
	assert.True(t, result)
	assert.Nil(t, appError)
}

func TestDeleteUserShouldReturnTrue(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := domain.NewMockUserRepository(controller)
	mockUserRepository.EXPECT().Delete(gomock.Any()).Return(true, nil)

	service := NewUserService(mockUserRepository)
	//Act
	result, appError := service.DeleteUser(1)
	//Assert
	assert.True(t, result)
	assert.Nil(t, appError)
}

func TestGetUserShouldReturnUser(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)
	defer controller.Finish()

	expectedUser := domain2.User{
		Id:   1,
		Name: "Bob",
		Age:  20,
	}

	mockUserRepository := domain.NewMockUserRepository(controller)
	mockUserRepository.EXPECT().Get(gomock.Any()).Return(expectedUser, nil)

	service := NewUserService(mockUserRepository)
	//Act
	user, appError := service.GetUser(1)
	//Assert
	assert.Equal(t, expectedUser.Id, user.Id)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Age, user.Age)
	assert.Nil(t, appError)
}

func TestGetAllUserShouldReturnUsers(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)
	defer controller.Finish()

	expectedUsers := []domain2.User{
		{
			Id:   1,
			Name: "Bob",
			Age:  20,
		},
		{
			Id:   2,
			Name: "Alice",
			Age:  21,
		},
	}

	mockUserRepository := domain.NewMockUserRepository(controller)
	mockUserRepository.EXPECT().GetAll().Return(expectedUsers, nil)

	service := NewUserService(mockUserRepository)
	//Act
	users, appError := service.GetAllUser()
	//Assert
	assert.True(t, len(expectedUsers) == len(users))
	assert.Nil(t, appError)
}
