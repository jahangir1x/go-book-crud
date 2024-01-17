package routes

import (
	"book-crud/pkg/controllers"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type bookRoutes struct {
	echo    *echo.Echo
	bookCtr controllers.BookController
}

func BookRoutes(echo *echo.Echo, bookCtr controllers.BookController) *bookRoutes {
	return &bookRoutes{
		echo:    echo,
		bookCtr: bookCtr,
	}
}

func (bc *bookRoutes) InitBookRoute() {
	e := bc.echo
	bc.initBookRoutes(e)
}

func (bc *bookRoutes) initBookRoutes(e *echo.Echo) {
	//grouping route endpoints
	book := e.Group("/bookstore")

	book.GET("/ping", Pong)

	//initializing http methods - routing endpoints and their handlers
	book.POST("/books", bc.bookCtr.CreateBook)
	book.GET("/books", bc.bookCtr.GetBook)
	book.PUT("/books/:bookID", bc.bookCtr.UpdateBook)
	book.DELETE("/books/:bookID", bc.bookCtr.DeleteBook)
}

func Pong(ctx echo.Context) error {
	fmt.Println("Pong")
	return ctx.JSON(http.StatusOK, "Pong")
}
