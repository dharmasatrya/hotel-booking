package service

import (
	"context"

	"github.com/dharmasatrya/hotel-booking/user-service/dto"
	pb "github.com/dharmasatrya/hotel-booking/user-service/proto"
)

type UserService interface {
	RegisterUser(dto.RegisterRequest) (*pb.RegisterUserResponse, error)
}

type userService struct {
	Client pb.UserServiceClient
}

func NewUserService(userClient pb.UserServiceClient) *userService {
	return &userService{userClient}
}

func (us *userService) RegisterUser(payload dto.RegisterRequest) (*pb.RegisterUserResponse, error) {
	res, err := us.Client.RegisterUser(context.Background(), &pb.RegisterUserRequest{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Phone:    payload.Phone,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
