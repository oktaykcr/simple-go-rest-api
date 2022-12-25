package app

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/oktaykcr/simple-go-rest-api/config"
	_ "github.com/oktaykcr/simple-go-rest-api/docs"
	"github.com/oktaykcr/simple-go-rest-api/domain"
	"github.com/oktaykcr/simple-go-rest-api/logger"
	"github.com/oktaykcr/simple-go-rest-api/service"
	"net/http"
)

func Start() {
	logger.Info("Starting the application...")
	config.Initialize()
	app := fiber.New()
	app.Use(cors.New())

	client := initDb()

	userHandler := UserHandlers{service: service.NewUserService(domain.NewUserRepositoryImpl(client))}

	setupRoutes(app, userHandler)

	err := app.Listen(fmt.Sprintf("%s:%s", config.Conf.Server.Host, config.Conf.Server.Port))
	if err != nil {
		logger.Error(err.Error())
	}
}

func setupRoutes(app *fiber.App, userHandlers UserHandlers) {
	app.Post("/api/v1/user", userHandlers.createUser)
	app.Put("/api/v1/user", userHandlers.updateUser)
	app.Get("/api/v1/user", userHandlers.getAllUser)
	app.Get("/api/v1/user/:id", userHandlers.getUser)
	app.Delete("/api/v1/user/:id", userHandlers.deleteUser)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", swaggerRedirect)
}

func initDb() *pg.DB {
	opt, err := pg.ParseURL(fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/postgres?sslmode=disable", //hostname=db for docker-compose
		config.Conf.Database.Username,
		config.Conf.Database.Password))
	if err != nil {
		panic(err)
	}
	client := pg.Connect(opt)

	err = createSchema(client)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	return client
}

func createSchema(client *pg.DB) error {
	models := []interface{}{
		(*domain.User)(nil),
	}
	for _, model := range models {
		err := client.Model(model).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}
	//seedUser(client)
	return nil
}

func seedUser(client *pg.DB) {
	customer := &domain.User{
		Name: "Bob", Age: 20,
	}
	_, err := client.Model(customer).Insert()
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
}

func swaggerRedirect(ctx *fiber.Ctx) error {
	return ctx.Redirect("/swagger/index.html", http.StatusSeeOther)
}
