package main

import (
	"gitlab.com/saladin2098/finance_tracker1/notification_service/config"
	"gitlab.com/saladin2098/finance_tracker1/notification_service/server"
)

func main() {
	//configurations
	cfg := config.Load()

	// Run
	server.Run(&cfg)
}
