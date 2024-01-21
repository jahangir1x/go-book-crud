package controllers

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type IAuthorController interface {
	CreateAuthor(e echo.Context) error
	GetAuthor(e echo.Context) error
	GetAllAuthors(e echo.Context) error
	UpdateAuthor(e echo.Context) error
	DeleteAuthor(e echo.Context) error
}

type AuthorController struct {
	authorSvc domain.IAuthorService
}

func NewAuthorController(authorSvc domain.IAuthorService) AuthorController {
	return AuthorController{
		authorSvc: authorSvc,
	}
}

func (authorService *AuthorController) CreateAuthor(e echo.Context) error {
	authorRequest := &types.AuthorRequest{}
	if err := e.Bind(authorRequest); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}

	// validate the request body
	if err := authorRequest.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	author := &models.AuthorDetail{
		AuthorName:  authorRequest.AuthorName,
		Address:     authorRequest.Address,
		PhoneNumber: authorRequest.PhoneNumber,
	}
	if err := authorService.authorSvc.CreateAuthor(author); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "Author was created successfully")
}

func (authorService *AuthorController) GetAuthor(e echo.Context) error {
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil && tempAuthorID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid author ID")
	}
	author, err := authorService.authorSvc.GetAuthor(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, author)
}

func (authorService *AuthorController) GetAllAuthors(e echo.Context) error {
	authors, err := authorService.authorSvc.GetAllAuthors()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, authors)
}

func (authorService *AuthorController) UpdateAuthor(e echo.Context) error {
	authorRequest := &types.AuthorRequest{}
	if err := e.Bind(authorRequest); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid author ID")
	}
	updatedAuthor := &models.AuthorDetail{
		ID:          uint(authorID),
		AuthorName:  authorRequest.AuthorName,
		Address:     authorRequest.Address,
		PhoneNumber: authorRequest.PhoneNumber,
	}
	if err := authorService.authorSvc.UpdateAuthor(updatedAuthor); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Author was updated successfully")
}

func (authorService *AuthorController) DeleteAuthor(e echo.Context) error {
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	_, err = authorService.authorSvc.GetAuthor(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := authorService.authorSvc.DeleteAuthor(uint(authorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Author was deleted successfully")
}
