package http

import (
	"errors"
	"net/http"

	"github.com/JerryJeager/ghost-write-backend/models/messages"
	sv "github.com/JerryJeager/ghost-write-backend/service/messages"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageController struct {
	serv sv.MessageSv
}

func NewMessageController(serv sv.MessageSv) *MessageController {
	return &MessageController{serv: serv}
}

func (o *MessageController) CreateMessage(ctx *gin.Context) {
	var pp UserIDPathParm
	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user id is of invalid format"})
		return
	}

	var message messages.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request object"})
		return
	}

	err := o.serv.CreateMessage(ctx, uuid.MustParse(pp.UserID), &message)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (o *MessageController) GetMessages(ctx *gin.Context) {
	var pp UserIDPathParm
	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user id is of invalid format"})
		return
	}

	messages, err := o.serv.GetMessages(ctx, uuid.MustParse(pp.UserID))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		ctx.Status(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, *messages)
}
