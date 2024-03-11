package main

import (
	"fmt"
	"portfolio-be/config"
	"portfolio-be/internal/app"
)

func main() {
	cfg, err := config.NewConfig(".")
	if err != nil {
		fmt.Printf("Config error: %s", err)
	}
	app.Run(cfg)
}
