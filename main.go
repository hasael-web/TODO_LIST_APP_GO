package main

import (
	"fmt"
	"net/http"
	"os"
	"todo_app/config"
	"todo_app/config/migration"
	"todo_app/handler"
	"todo_app/repository"
	"todo_app/router"
	"todo_app/service"
	"todo_app/utils"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "todo_app/docs"
)

// @Documentasi TODO LIST APP (Test Technical)
// @version 1.0
// @description API endpoint Todo List App.
// @termsOfService http://swagger.io/terms/

// @contact.name Hasael Butar Butar
// @contact.email hasaelbutarbutar80@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3005
// @BasePath /api/v1

func init() {
	ENV := "DEV"
	env := utils.NewEnv()
	env.Load(ENV)

}

func main() {
	port := os.Getenv("PORT")
	fmt.Println(port)
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "welcome to my api")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	idb := config.InitDB()
	db, err := idb.ConnectDb()

	if err != nil {
		e.Logger.Fatal(err)
	}

	migration := &migration.MigrationDB{DB: db}
	err = migration.AutoMigration()

	if err != nil {
		e.Logger.Fatal(err)
	}

	// todo init
	todoRepo := repository.NewTodoRepository(migration)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)
	routeTodo := router.NewTodoRoute(todoHandler)

	routeTodo.Route(e.Group("/api/v1"))

	// sub todo init
	subTodoRepo := repository.NewSubTodoRepository(migration)
	subTodoService := service.NewSubTodoService(subTodoRepo)
	subTodoHandler := handler.NewSubTodoHandler(subTodoService)
	routeSubTodo := router.NewSubTodoRoute(subTodoHandler)

	routeSubTodo.Route(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start(":" + port))
}
