package service

import (
	"context"

	"github.com/dharmasatrya/hotel-booking/room-service/dto"
	pb "github.com/dharmasatrya/hotel-booking/room-service/proto"
	"github.com/dharmasatrya/hotel-booking/room-service/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RoomService struct {
	roomRepository repository.RoomRepository
	pb.UnimplementedUserServiceServer
}

func NewRoomService(roomRepository repository.RoomRepository) *RoomService {
	return &RoomService{
		roomRepository: roomRepository,
	}
}

func (s *RoomService) GetRoomById(ctx context.Context, req *pb.GetRoomByIdRequest) (*pb.Room, error) {
	room, err := s.roomRepository.GetRoomById(req.RoomId)
	if err != nil {
		return nil, err
	}

	return &pb.Room{
		Id:          room.ID,
		Name:        room.Name,
		Capacity:    room.Capacity,
		Price:       room.Price,
		IsAvailable: room.IsAvailable,
	}, nil
}

func (s *RoomService) GetAvailableRooms(ctx context.Context, _ *emptypb.Empty) (*pb.GetAvailableRoomsResponse, error) {
	rooms, err := s.roomRepository.GetAvailableRooms()
	if err != nil {
		return nil, err
	}

	var pbRooms []*pb.Room
	for _, room := range rooms {
		pbRooms = append(pbRooms, &pb.Room{
			Id:          room.ID,
			Name:        room.Name,
			Capacity:    room.Capacity,
			Price:       room.Price,
			IsAvailable: room.IsAvailable,
		})
	}

	return &pb.GetAvailableRoomsResponse{
		Rooms: pbRooms,
	}, nil
}

func (s *RoomService) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.Room, error) {
	room := dto.Room{
		Name:     req.Name,
		Capacity: req.Capacity,
		Price:    req.Price,
	}

	createdRoom, err := s.roomRepository.CreateRoom(room)
	if err != nil {
		return nil, err
	}

	return &pb.Room{
		Id:          createdRoom.ID,
		Name:        createdRoom.Name,
		Capacity:    createdRoom.Capacity,
		Price:       createdRoom.Price,
		IsAvailable: createdRoom.IsAvailable,
	}, nil
}

func (s *RoomService) EditRoomById(ctx context.Context, req *pb.EditRoomByIdRequest) (*pb.Room, error) {
	room := dto.Room{
		Name:        req.Name,
		Capacity:    req.Capacity,
		Price:       req.Price,
		IsAvailable: req.IsAvailable,
	}

	updatedRoom, err := s.roomRepository.EditRoomById(req.Id, room)
	if err != nil {
		return nil, err
	}

	return &pb.Room{
		Id:          updatedRoom.ID,
		Name:        updatedRoom.Name,
		Capacity:    updatedRoom.Capacity,
		Price:       updatedRoom.Price,
		IsAvailable: updatedRoom.IsAvailable,
	}, nil
}

func (s *RoomService) DeleteRoomById(ctx context.Context, req *pb.DeleteRoomByIdRequest) (*pb.Room, error) {
	deletedRoom, err := s.roomRepository.DeleteRoomById(req.RoomId)
	if err != nil {
		return nil, err
	}

	return &pb.Room{
		Id:          deletedRoom.ID,
		Name:        deletedRoom.Name,
		Capacity:    deletedRoom.Capacity,
		Price:       deletedRoom.Price,
		IsAvailable: deletedRoom.IsAvailable,
	}, nil
}
