package main

import (
	_ "github.com/DaniaLD/upera-go-test/docs"
	"github.com/DaniaLD/upera-go-test/internal/app/router"
	"github.com/DaniaLD/upera-go-test/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

// @title           Upera - Product Versioning
// @version         1.0
// @description     This is an assessment for Upera company
// @contact.name   	API Support
// @contact.email  	danial.6274@gmail.com
// @BasePath  		/
func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configs: %v", err)
	}

	app := fiber.New()
	rtr := router.NewRouter(app)
	rtr.InitRouter()

	app.Listen(":" + config.Port)
}
