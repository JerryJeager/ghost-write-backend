package manualwire

import (
	"github.com/JerryJeager/ghost-write-backend/config"
	http "github.com/JerryJeager/ghost-write-backend/http"
	"github.com/JerryJeager/ghost-write-backend/service/messages"
	"github.com/JerryJeager/ghost-write-backend/service/users"
)

func GetUserRepository() *users.UserRepo {
	repo := config.GetSession()
	return users.NewUserRepo(repo)
}

func GetUserService(repo users.UserStore) *users.UserServ {
	return users.NewUserService(repo)
}

func GetUserController() *http.UserController {
	repo := GetUserRepository()
	service := GetUserService(repo)
	return http.NewUserController(service)
}

func GetMessageRepository() *messages.MessageRepo {
	repo := config.GetSession()
	return messages.NewMessageRepo(repo)
}

func GetMessageService(repo messages.MessageStore) *messages.MessageServ {
	return messages.NewMessageService(repo)
}

func GetMessageController() *http.MessageController {
	repo := GetMessageRepository()
	service := GetMessageService(repo)
	return http.NewMessageController(service)
}
