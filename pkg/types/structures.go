package types

import validation "github.com/go-ozzo/ozzo-validation"

// response struct | marshalled into json fromat from struct
type BookRequest struct {
	ID          uint   `json:"bookID"`
	BookName    string `json:"bookName"`
	AuthorID    uint   `json:"authorID"`
	Publication string `json:"publication,omitempty"`
}

type AuthorRequest struct {
	ID          uint   `json:"authorID"`
	AuthorName  string `json:"authorName"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

func (book BookRequest) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&book.AuthorID,
			validation.Required.Error("Author ID cannot be empty")))
}

func (author AuthorRequest) Validate() error {
	return validation.ValidateStruct(&author,
		validation.Field(&author.AuthorName,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(1, 50)))
}
