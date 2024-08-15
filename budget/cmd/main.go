package main

import (
	"gitlab.com/saladin2098/finance_tracker1/budget/internal/app"
	"gitlab.com/saladin2098/finance_tracker1/budget/internal/pkg/config"
)

func main() {
	cfg := config.Load()

	app.Run(cfg)
}
