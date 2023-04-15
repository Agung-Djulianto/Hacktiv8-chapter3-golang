package service

import (
	"chap3-challenge2/helper"
	"chap3-challenge2/model"
	"chap3-challenge2/repository"
)

const (
	ROLE_USER  = "user"
	ROLE_ADMIN = "admin"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) AddUser(user model.UserRegisterRequest) (model.UserRegisterResponse, error) {
	id := helper.GenerateID()
	hashed_password, err := helper.HashPassword(user.Password)

	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	newUser := model.User{
		ID:       id,
		Email:    user.Email,
		Password: hashed_password,
		Role:     ROLE_USER,
	}

	res, err := us.UserRepository.Save(newUser)
	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	return model.UserRegisterResponse{
		ID:    res.ID,
		Email: res.Email,
		Role:  res.Role,
	}, nil
}

func (us *UserService) AddAdmin(user model.UserRegisterRequest) (model.UserRegisterResponse, error) {
	id := helper.GenerateID()
	hashed_password, err := helper.HashPassword(user.Password)

	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	newUser := model.User{
		ID:       id,
		Email:    user.Email,
		Password: hashed_password,
		Role:     ROLE_ADMIN,
	}

	res, err := us.UserRepository.Save(newUser)
	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	return model.UserRegisterResponse{
		ID:    res.ID,
		Email: res.Email,
		Role:  res.Role,
	}, nil
}

func (us *UserService) Login(userLoginRequest model.UserLoginRequest) (model.UserLoginResponse, error) {
	email := userLoginRequest.Email
	password := userLoginRequest.Password

	res, err := us.UserRepository.GetByEmail(email)
	if err != nil {
		return model.UserLoginResponse{}, err
	}

	if !helper.CheckPasswordHash(password, res.Password) {
		return model.UserLoginResponse{}, model.ErrorInvalidEmailOrPassword
	}

	token, err := helper.GenerateToken(res.ID, res.Role)
	if err != nil {
		return model.UserLoginResponse{}, err
	}

	return model.UserLoginResponse{
		Token: token,
	}, nil
}
