package users

import (
	"context"

	"github.com/JerryJeager/ghost-write-backend/models/users"
	"github.com/JerryJeager/ghost-write-backend/utils"
	"github.com/google/uuid"
)

type UserSv interface {
	CreateUser(ctx context.Context, user *users.User) (string, error)
	CreateToken(ctx context.Context, user *users.User) (string, string, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*users.User, error)
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
}

func (o *UserServ) GetUser(ctx context.Context, userID uuid.UUID) (*users.User, error) {
	return o.repo.GetUser(ctx, userID)
}

func (o *UserServ) CreateUser(ctx context.Context, user *users.User) (string, error) {
	id := uuid.New()
	user.ID = id
	if err := user.HashPassword(); err != nil {
		return "", err
	}
	err := o.repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (o *UserServ) CreateToken(ctx context.Context, user *users.User) (string, string, error) {
	pas, err := o.repo.CreateToken(ctx, user.Email, user)
	if err != nil {
		return "", "", err
	}

	err = users.VerifyPassword(user.Password, pas)

	if err != nil {
		return "", "", err
	}

	id, err := o.repo.GetUserID(ctx, user.Email)
	
	if err != nil{
		return "", "", err
	}

	token, err := utils.GenerateToken(*user)

	if err != nil {
		return "", "", err
	}

	return id, token, nil
}
