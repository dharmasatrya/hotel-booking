package repository

import (
	"fmt"
	"os"
	"time"

	"github.com/dharmasatrya/hotel-booking/user-service/dto"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(dto.User) (*dto.User, error)
	LoginUser(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) RegisterUser(user dto.User) (*dto.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) LoginUser(req dto.LoginRequest) (*dto.LoginResponse, error) {
	var user dto.User
	if err := r.db.Table("users").Where("email = ?", req.Email).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Printf("Password comparison error: %v\n", err)
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: tokenString,
	}, nil
}
