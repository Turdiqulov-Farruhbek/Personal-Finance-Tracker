package main

import (
	"finance_tracker/gateway/internal/app"
	"finance_tracker/gateway/internal/pkg/config"
)

func main() {
	cfg := config.Load()

	app.Run(cfg)
}
