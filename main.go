package main

import (
	"fmt"

	"github.com/JerryJeager/ghost-write-backend/cmd"
	"github.com/JerryJeager/ghost-write-backend/config"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}
func main() {
	fmt.Printf("executing the api routes for ghost write...")
	cmd.ExecuteApiRoutes()
}