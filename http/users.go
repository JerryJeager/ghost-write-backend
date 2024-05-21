package http

import (
	"net/http"

	"github.com/JerryJeager/ghost-write-backend/models/users"
	sv "github.com/JerryJeager/ghost-write-backend/service/users"
	"github.com/gin-gonic/gin"
)
type UserController struct {
	serv sv.UserSv
}

func NewUserController(serv sv.UserSv) *UserController {
	return &UserController{serv: serv}
}


func (o *UserController) CreateUser(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, err := o.serv.CreateUser(ctx, &user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user with this email already exists"})
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (o *UserController) CreateToken(ctx *gin.Context) {

	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, token, err := o.serv.CreateToken(ctx, &user)
	

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email or password is invaiid"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"user_id": id, "token": token})

}