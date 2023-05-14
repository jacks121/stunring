package services

import (
	"errors"
	"swetelove/models"
	"swetelove/repositories"
	"swetelove/utils"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(db *gorm.DB) (*UserService, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}

	return &UserService{
		userRepo: repositories.NewUserRepository(db),
	}, nil
}

func (s *UserService) Register(username string, password string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password cannot be empty")
	}
	existingUser, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}
	passwordHash := utils.HashPassword(password)
	newUser := &models.User{Username: username, Password: passwordHash}
	if err := s.userRepo.CreateUser(newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *UserService) Login(username string, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsernameAndPassword(username, utils.HashPassword(password))
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid username or password")
	}
	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserService) GetUserByToken(tokenString string) (*models.User, error) {
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	userId := uint(claims["id"].(float64))
	user, err := s.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
