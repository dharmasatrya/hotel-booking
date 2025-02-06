package service

import (
	"context"
	"log"

	"github.com/dharmasatrya/hotel-booking/user-service/dto"
	"github.com/dharmasatrya/hotel-booking/user-service/repository"

	pb "github.com/dharmasatrya/hotel-booking/user-service/proto"
)

type UserService struct {
	userRepository repository.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserService(
	userRepository repository.UserRepository,
) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := dto.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	}

	res, err := s.userRepository.RegisterUser(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.RegisterUserResponse{
		Id:    res.ID,
		Name:  res.Name,
		Email: res.Email,
		Phone: res.Phone,
	}, nil
}

func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	payload := dto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := s.userRepository.LoginUser(payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.LoginUserResponse{
		Token: res.Token,
	}, nil
}
