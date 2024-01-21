package routes

import (
	"book-crud/pkg/controllers"
	"book-crud/pkg/middlewares"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookRoutes struct {
	echo    *echo.Echo
	bookCtr controllers.BookController
}

func NewBookRoutes(echo *echo.Echo, bookCtr controllers.BookController) *BookRoutes {
	return &BookRoutes{
		echo:    echo,
		bookCtr: bookCtr,
	}
}

func (bc *BookRoutes) InitBookRoute() {
	e := bc.echo
	bc.initBookRoutes(e)
}

func (bc *BookRoutes) initBookRoutes(e *echo.Echo) {
	//grouping route endpoints
	book := e.Group("/bookstore")

	book.GET("/ping", Pong)

	//initializing http methods - routing endpoints and their handlers
	book.GET("/books", bc.bookCtr.GetAllBooks)

	book.Use(middlewares.ValidateToken)
	book.POST("/books", bc.bookCtr.CreateBook)
	book.GET("/books/:bookID", bc.bookCtr.GetBook)
	book.PUT("/books/:bookID", bc.bookCtr.UpdateBook)
	book.DELETE("/books/:bookID", bc.bookCtr.DeleteBook)
}

func Pong(ctx echo.Context) error {
	fmt.Println("Pong")
	return ctx.JSON(http.StatusOK, "Pong")
}
