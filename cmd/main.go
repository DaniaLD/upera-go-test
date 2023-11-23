package main

import (
	"github.com/gofiber/fiber/v2"
	"go-upera-test/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configs: %v", err)
	}

	app := fiber.New()

	app.Listen(":" + config.Port)
}
