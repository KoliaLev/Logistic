package services

import (
	"errors"
	"fmt"
	"logistics/models"
	"logistics/repositories"

	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Birthday string `json:"birthDay" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
}

type AuthService interface {
	Login(LoginReq *LoginRequest) (*models.User, string, error)
	Register(registUser *RegisterRequest) (*models.User, error)
}

type authServiseStruct struct {
	authRepo repositories.AuthRepo
}

func NewAuthService(repo repositories.AuthRepo) AuthService {
	return &authServiseStruct{
		authRepo: repo,
	}
}

func (s *authServiseStruct) Login(LoginReq *LoginRequest) (*models.User, string, error) {
	fmt.Println("call login inthe authService")
	s.authRepo.Login()
	user, err := s.authRepo.GetByEmail(LoginReq.Email)
	if err != nil {
		logrus.Info("err getting user in the login service")
		return nil, "", err
	}

	if user.Password != LoginReq.Password {
		return nil, "", errors.New("incorect email or password")
	}
	token := CreateToken(int(user.ID), user.Email)
	return user, token, nil
}

func (s *authServiseStruct) Register(registUser *RegisterRequest) (*models.User, error) {
	fmt.Println("call register in the authService ", registUser)

	// check whether user already registered
	_, err := s.authRepo.GetByEmail(registUser.Email)
	if err == nil {
		return nil, errors.New("such a user already exists")
	}

	user, err := s.authRepo.Register(&models.User{
		Email:    registUser.Email,
		Password: registUser.Password,
		Birthday: registUser.Birthday,
		Sex:      registUser.Sex,
	})
	if err != nil {
		logrus.Info("in the register service ", err)
		return nil, err
	}
	return user, nil
}
