package users

import (
	"context"
	"fmt"
	"log"

	"github.com/JerryJeager/ghost-write-backend/config"
	"github.com/JerryJeager/ghost-write-backend/models/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type UserStore interface {
	CreateUser(ctx context.Context, user *users.User) error
	CreateToken(ctx context.Context, userEmail string, user *users.User) (string, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*users.User, error)
	GetUserID(ctx context.Context, email string) (string, error)
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {	
	return &UserRepo{client: client}
}

func (o *UserRepo) GetUser(ctx context.Context, userID uuid.UUID) (*users.User, error) {
	var user users.User
	query := config.Session.First(&user, "id = ?", userID).WithContext(ctx)
	if query.Error != nil {
		return &users.User{}, query.Error
	}
	return &user, nil
}

func (o *UserRepo) GetUserID(ctx context.Context, email string) (string, error) {
	var user users.User
	query := config.Session.First(&user, "email = ?", email).WithContext(ctx)

	if query.Error != nil{
		return "", query.Error
	}
	return user.ID.String(), nil
}

func (o *UserRepo) CreateUser(ctx context.Context, user *users.User) error {
	if o.client != nil {
		fmt.Println("gorm client is available in the store...")
	}
	query := config.Session.Create(&user).WithContext(ctx)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (o *UserRepo) CreateToken(ctx context.Context, userEmail string, user *users.User) (string, error) {
	var userModel users.User

	if err := config.Session.First(&userModel, "email = ?", userEmail).WithContext(ctx).Error; err != nil {
		return "", err
	}

	log.Print(userModel.Email)

	return userModel.Password, nil
}