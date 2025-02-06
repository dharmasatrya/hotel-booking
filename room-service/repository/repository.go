package repository

import (
	"errors"

	"github.com/dharmasatrya/hotel-booking/room-service/dto"
	"gorm.io/gorm"
)

type RoomRepository interface {
	GetRoomById(id uint32) (*dto.Room, error)
	GetAvailableRooms() ([]dto.Room, error)
	CreateRoom(room dto.Room) (*dto.Room, error)
	EditRoomById(id uint32, room dto.Room) (*dto.Room, error)
	DeleteRoomById(id uint32) (*dto.Room, error)
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *roomRepository {
	return &roomRepository{db}
}

func (r *roomRepository) GetRoomById(id uint32) (*dto.Room, error) {
	var room dto.Room
	if err := r.db.First(&room, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) GetAvailableRooms() ([]dto.Room, error) {
	var rooms []dto.Room
	if err := r.db.Where("is_available = ?", true).Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *roomRepository) CreateRoom(room dto.Room) (*dto.Room, error) {
	room.IsAvailable = true // Set default availability to true for new rooms
	if err := r.db.Create(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) EditRoomById(id uint32, room dto.Room) (*dto.Room, error) {
	existingRoom, err := r.GetRoomById(id)
	if err != nil {
		return nil, err
	}

	existingRoom.Name = room.Name
	existingRoom.Capacity = room.Capacity
	existingRoom.Price = room.Price
	existingRoom.IsAvailable = room.IsAvailable

	if err := r.db.Save(existingRoom).Error; err != nil {
		return nil, err
	}

	return existingRoom, nil
}

func (r *roomRepository) DeleteRoomById(id uint32) (*dto.Room, error) {
	room, err := r.GetRoomById(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Delete(&room).Error; err != nil {
		return nil, err
	}

	return room, nil
}
