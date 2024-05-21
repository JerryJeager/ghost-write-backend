package http

import "github.com/JerryJeager/ghost-write-backend/service/users"

type UserController struct {
	serv users.UserSv
}

func NewUserController(serv users.UserSv) *UserController {
	return &UserController{serv: serv}
}


