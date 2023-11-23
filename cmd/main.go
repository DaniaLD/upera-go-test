package main

import (
	_ "github.com/DaniaLD/upera-go-test/docs"
	"github.com/DaniaLD/upera-go-test/internal/app/router"
	productAppService "github.com/DaniaLD/upera-go-test/internal/app/service/product"
	productDomainService "github.com/DaniaLD/upera-go-test/internal/domain/product/service"
	"github.com/DaniaLD/upera-go-test/internal/infrastructure/database"
	productRepositoryImpl "github.com/DaniaLD/upera-go-test/internal/repository/product"
	"github.com/DaniaLD/upera-go-test/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
)

// @title           Upera - Product Versioning
// @version         1.0
// @description     This is an assessment for Upera company
// @contact.name   	API Support
// @contact.email  	danial.6274@gmail.com
// @BasePath  		/
func main() {
	utils.LoadConfig()
	dbName := viper.GetString("database.mongo.dbName")
	mongoClient, err := database.NewMongoClient(viper.GetString("database.mongo.uri"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	prdRepo := productRepositoryImpl.NewProductRepository(mongoClient, dbName)
	prdUseCase := productDomainService.NewProductService(prdRepo)
	prdAppService := productAppService.NewProductAppService(prdUseCase)

	app := fiber.New()
	rtr := router.NewRouter(app, prdAppService)
	rtr.InitRouter()

	app.Listen(":" + viper.GetString("port"))
}
