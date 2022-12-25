package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/oktaykcr/simple-go-rest-api/domain"
	"github.com/oktaykcr/simple-go-rest-api/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app *fiber.App
var client *pg.DB
var userHandler *UserHandlers
var userRepository domain.UserRepository

const userBasePath = "/api/v1/user"

var testUser = &domain.User{
	Id:   999,
	Name: "testUser",
	Age:  20,
}

func TestMain(m *testing.M) {
	fmt.Println("Before")
	app = fiber.New()
	client = initDB()
	userRepository = domain.NewUserRepositoryImpl(client)
	userHandler = &UserHandlers{service: service.NewUserService(userRepository)}
	m.Run()
	fmt.Println("After")
	userRepository.DeleteByName(testUser.Name)
	defer client.Close()
}

func initDB() *pg.DB {
	client := pg.Connect(&pg.Options{
		User:     "root",
		Password: "password",
	})
	err := createSchema(client)
	if err != nil {
		panic(err)
	}
	return client
}

func TestArticleHandler_HandleCreateArticle(t *testing.T) {
	// Setup
	app.Post(userBasePath, userHandler.createUser)

	bodyMarshal, _ := json.Marshal(createUserRequest{
		Name: testUser.Name,
		Age:  testUser.Age,
	})
	req := httptest.NewRequest(http.MethodPost, userBasePath, bytes.NewReader(bodyMarshal))
	req.Header.Set("Content-Type", "application/json")
	req.Body.Close()

	resp, err := app.Test(req, 10)
	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
