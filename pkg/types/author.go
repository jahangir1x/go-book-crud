package types

import validation "github.com/go-ozzo/ozzo-validation"

type AuthorRequest struct {
	ID          uint   `json:"authorID"`
	AuthorName  string `json:"authorName"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

func (author AuthorRequest) Validate() error {
	return validation.ValidateStruct(&author,
		validation.Field(&author.AuthorName,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(2, 64)),
		validation.Field(&author.Address,
			validation.Length(2, 128)),
		validation.Field(&author.PhoneNumber,
			validation.Length(8, 16)))

}
