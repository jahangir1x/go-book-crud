package services

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"book-crud/pkg/utils"
)

type AuthService struct {
	userRepo domain.IUserRepo
}

func AuthServiceInstance(userRepo domain.IUserRepo) domain.IAuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (service *AuthService) LoginUser(loginRequest *types.LoginRequest) (*types.LoginResponse, error) {
	existingUser, err := service.userRepo.GetUser(loginRequest.UserName)
	if err != nil {
		return &types.LoginResponse{}, err
	}
	if err != nil {
		return &types.LoginResponse{}, err
	}
	err = utils.CheckPassword(existingUser.PasswordHash, loginRequest.Password)
	if err != nil {
		return &types.LoginResponse{}, err
	}

	// Generate JWT token
	token, err := utils.GetJwtForUser(existingUser.Username)
	if err != nil {
		return &types.LoginResponse{}, err
	}

	return &types.LoginResponse{
		Token: token,
	}, nil

}

func (service *AuthService) SignupUser(registerRequest *types.RegisterRequest) error {
	passwordHash, err := utils.GetHashedPassword(registerRequest.Password)
	if err != nil {
		return err
	}
	user := &models.UserDetail{
		Username:     registerRequest.UserName,
		PasswordHash: passwordHash,
		Name:         registerRequest.Name,
		Email:        registerRequest.Email,
		Address:      registerRequest.Address,
	}
	if err := service.userRepo.CreateUser(user); err != nil {
		return err
	}
	return nil
}
