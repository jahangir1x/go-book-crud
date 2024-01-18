package containers

import (
	"book-crud/pkg/config"
	"book-crud/pkg/connection"
	"book-crud/pkg/controllers"
	"book-crud/pkg/repositories"
	"book-crud/pkg/routes"
	"book-crud/pkg/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func Serve(e *echo.Echo) {
	//config initialization
	config.SetConfig()

	//database initializations
	db := connection.GetDB()

	// repository initialization
	bookRepo := repositories.BookDBInstance(db)
	authorRepo := repositories.AuthorDBInstance(db)

	//service initialization
	bookService := services.BookServiceInstance(bookRepo, authorRepo)
	authorService := services.AuthorServiceInstance(authorRepo, bookRepo)

	//controller initialization
	bookCtr := controllers.NewBookController(bookService)
	authorCtr := controllers.NewAuthorController(authorService)

	//route initialization
	b := routes.BookRoutes(e, bookCtr)
	authorRoutes := routes.AuthorRoutes(e, authorCtr)

	b.InitBookRoute()
	authorRoutes.InitAuthorRoutes()

	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
