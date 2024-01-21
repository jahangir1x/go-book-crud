package routes

import (
	"book-crud/pkg/controllers"
	"book-crud/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type AuthorRoutes struct {
	echo             *echo.Echo
	authorController controllers.AuthorController
}

func NewAuthorRoutes(echo *echo.Echo, authorController controllers.AuthorController) *AuthorRoutes {
	return &AuthorRoutes{
		echo:             echo,
		authorController: authorController,
	}
}

func (authorRoutes *AuthorRoutes) InitAuthorRoutes() {
	e := authorRoutes.echo
	authorRoutes.initAuthorRoutes(e)
}

func (authorRoutes *AuthorRoutes) initAuthorRoutes(e *echo.Echo) {
	author := e.Group("/bookstore")
	author.GET("/authors", authorRoutes.authorController.GetAllAuthors)

	author.Use(middlewares.ValidateToken)
	author.POST("/authors", authorRoutes.authorController.CreateAuthor)
	author.GET("/authors/:authorID", authorRoutes.authorController.GetAuthor)
	author.PUT("/authors/:authorID", authorRoutes.authorController.UpdateAuthor)
	author.DELETE("/authors/:authorID", authorRoutes.authorController.DeleteAuthor)
}
