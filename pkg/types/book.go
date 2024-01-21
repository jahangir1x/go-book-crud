package types

import validation "github.com/go-ozzo/ozzo-validation"

type BookRequest struct {
	ID          uint   `json:"bookID"`
	BookName    string `json:"bookName"`
	AuthorID    uint   `json:"authorID"`
	Publication string `json:"publication,omitempty"`
}

func (book BookRequest) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 64)),
		validation.Field(&book.AuthorID,
			validation.Required.Error("Author ID cannot be empty")),
		validation.Field(&book.Publication,
			validation.Length(2, 64)))
}
