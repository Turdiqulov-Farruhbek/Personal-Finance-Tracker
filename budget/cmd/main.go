package main

import (
	"finance_tracker/budget/internal/app"
	"finance_tracker/budget/internal/pkg/config"
)

func main() {
	cfg := config.Load()

	app.Run(cfg)
}
