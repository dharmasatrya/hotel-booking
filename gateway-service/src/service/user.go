package service

import (
	"context"

	"github.com/dharmasatrya/hotel-booking/user-service/dto"
	pb "github.com/dharmasatrya/hotel-booking/user-service/proto"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(dto.RegisterRequest) (*pb.RegisterUserResponse, error)
	LoginUser(payload dto.LoginRequest) (*pb.LoginUserResponse, error)
}

type userService struct {
	Client pb.UserServiceClient
}

func NewUserService(userClient pb.UserServiceClient) *userService {
	return &userService{userClient}
}

func (us *userService) RegisterUser(payload dto.RegisterRequest) (*pb.RegisterUserResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	res, err := us.Client.RegisterUser(context.Background(), &pb.RegisterUserRequest{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
		Phone:    payload.Phone,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *userService) LoginUser(payload dto.LoginRequest) (*pb.LoginUserResponse, error) {

	res, err := us.Client.LoginUser(context.Background(), &pb.LoginUserRequest{
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
