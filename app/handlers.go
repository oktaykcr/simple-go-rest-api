package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oktaykcr/simple-go-rest-api/service"
	"strconv"
)

type UserHandlers struct {
	service service.DefaultUserService
}

type createUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// createUser godoc
// @tags User
// @Summary Create a new user.
// @Description create user by name and age.
// @Accept  json
// @Produce  json
// @Param request body createUserRequest true "GetContentsByListingIdsRequest"
// @Success 200
// @Router /api/v1/user  [post]
func (uh UserHandlers) createUser(ctx *fiber.Ctx) error {
	var request createUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	createUser, _ := uh.service.CreateUser(request.Name, request.Age)
	return ctx.Status(201).SendString(strconv.FormatBool(createUser))
}

type updateUserRequest struct {
	Id   *int64  `json:"id"`
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

// updateUser godoc
// @tags User
// @Summary Update a new user.
// @Description update name and age of user.
// @Accept  json
// @Produce  json
// @Param request body updateUserRequest true "GetContentsByListingIdsRequest"
// @Success 200
// @Router /api/v1/user  [put]
func (uh UserHandlers) updateUser(ctx *fiber.Ctx) error {
	var request updateUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	if request.Id == nil {
		return ctx.Status(400).SendString("id")
	}
	updateUser, _ := uh.service.UpdateUser(*request.Id, *request.Name, *request.Age)
	return ctx.Status(200).SendString(strconv.FormatBool(updateUser))
}

// getAllUser godoc
// @tags User
// @Summary Get All users.
// @Description List all users.
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/user  [get]
func (uh UserHandlers) getAllUser(ctx *fiber.Ctx) error {
	allUsers, _ := uh.service.GetAllUser()
	return ctx.Status(200).JSON(allUsers)
}

// getUser godoc
// @tags User
// @Summary Get user.
// @Description Get user according to given id.
// @Accept  json
// @Produce  json
// @Param id path int true "UserId"
// @Success 200
// @Router /api/v1/user/{id}  [get]
func (uh UserHandlers) getUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).SendString("id")
	}
	user, _ := uh.service.GetUser(int64(id))
	return ctx.Status(200).JSON(user)
}

// deleteUser godoc
// @tags User
// @Summary Delete a user.
// @Description Delete user according to given id.
// @Accept  json
// @Produce  json
// @Param id path int true "UserId"
// @Success 200
// @Router /api/v1/user/{id}  [delete]
func (uh UserHandlers) deleteUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).SendString("id")
	}
	result, _ := uh.service.DeleteUser(int64(id))
	return ctx.Status(200).SendString(strconv.FormatBool(result))
}
