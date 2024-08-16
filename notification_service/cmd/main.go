package main

import (
	"finance_tracker/notification_service/config"
	"finance_tracker/notification_service/server"
)

func main() {
	cfg := config.Load()
	server.Run(&cfg)
}
