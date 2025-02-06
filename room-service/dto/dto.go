package dto

type Room struct {
	ID          uint32 `gorm:"primaryKey"`
	Name        string
	Capacity    uint32
	Price       uint32
	IsAvailable bool
}
