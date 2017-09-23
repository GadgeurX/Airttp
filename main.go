package main

import (
	"Airttp/server"
	"Airttp/logger"
	"Airttp/config"
	"Airttp/modules"
)

func main() {
	modules.GetManagerInstance().LoadModules()

	serverPort := config.GetConfigInstance().GetServerPort(80)
	logger.GetInstance().InfoF("Server start at %d", serverPort)
	l_server := server.New(serverPort)
	l_server.Run()
}