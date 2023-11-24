package main

import (
	"context"
	_ "github.com/DaniaLD/upera-go-test/docs"
	"github.com/DaniaLD/upera-go-test/internal/app/consumer"
	productRevisionConsumerHandler "github.com/DaniaLD/upera-go-test/internal/app/consumer-handler/product-revision"
	"github.com/DaniaLD/upera-go-test/internal/app/router"
	productAppService "github.com/DaniaLD/upera-go-test/internal/app/service/product"
	productRevisionAppService "github.com/DaniaLD/upera-go-test/internal/app/service/product-revision"
	productRevisionDomainservice "github.com/DaniaLD/upera-go-test/internal/domain/product-revision/service"
	productDomainService "github.com/DaniaLD/upera-go-test/internal/domain/product/service"
	"github.com/DaniaLD/upera-go-test/internal/infrastructure/database"
	messagingQueue "github.com/DaniaLD/upera-go-test/internal/infrastructure/messaging-queue"
	productRepositoryImpl "github.com/DaniaLD/upera-go-test/internal/repository/product"
	productRevisionRepositoryImpl "github.com/DaniaLD/upera-go-test/internal/repository/product-revision"
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
	redisClient, err := database.NewRedisClient(viper.GetString("database.redis.uri"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	rabbitMQConnection, err := messagingQueue.NewConnection(viper.GetString("messagingQueue.rabbit.uri"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer func() {
		mongoClient.Disconnect(context.Background())
		redisClient.Close()
		rabbitMQConnection.Close()
	}()

	prdRvsRepo := productRevisionRepositoryImpl.NewProductRevisionRepository(mongoClient, dbName, redisClient)
	prdRvsUseCase := productRevisionDomainservice.NewProductRevisionService(prdRvsRepo)
	prdRvsAppService := productRevisionAppService.NewProductRevisionAppService(prdRvsUseCase)

	prdRepo := productRepositoryImpl.NewProductRepository(mongoClient, dbName, rabbitMQConnection)
	prdUseCase := productDomainService.NewProductService(prdRepo)
	prdAppService := productAppService.NewProductAppService(prdUseCase, prdRvsUseCase)

	app := fiber.New()
	rtr := router.NewRouter(app, prdAppService, prdRvsAppService)
	rtr.InitRouter()

	prdRvsCnsmrHndlr := productRevisionConsumerHandler.NewProductRevisionConsumerHandler(prdRvsAppService)
	cnsmr := consumer.NewConsumer(rabbitMQConnection, prdRvsCnsmrHndlr)
	cnsmr.InitConsumer()

	app.Listen(":" + viper.GetString("port"))
}
