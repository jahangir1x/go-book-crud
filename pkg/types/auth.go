package types

import validation "github.com/go-ozzo/ozzo-validation"

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (loginRequest LoginRequest) Validate() error {
	return validation.ValidateStruct(&loginRequest,
		validation.Field(&loginRequest.UserName,
			validation.Required.Error("Username cannot be empty")),
		validation.Field(&loginRequest.Password,
			validation.Required.Error("Password cannot be empty")))
}

type RegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

func (registerRequest RegisterRequest) Validate() error {
	return validation.ValidateStruct(&registerRequest,
		validation.Field(&registerRequest.UserName,
			validation.Required.Error("Username cannot be empty"),
			validation.Length(4, 32)),
		validation.Field(&registerRequest.Password,
			validation.Required.Error("Password cannot be empty"),
			validation.Length(8, 128)),
		validation.Field(&registerRequest.Name,
			validation.Required.Error("Name cannot be empty"),
			validation.Length(2, 64)),
		validation.Field(&registerRequest.Email,
			validation.Required.Error("Email cannot be empty"),
			validation.Length(4, 128)))
}
