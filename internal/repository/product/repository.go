package productRepositoryImpl

import (
	productRepository "github.com/DaniaLD/upera-go-test/internal/domain/product"
	messagingQueue "github.com/DaniaLD/upera-go-test/internal/infrastructure/messaging-queue"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepositoryImpl struct {
	client       *mongo.Client
	dbName       string
	rabbitMQConn *messagingQueue.Connection
}

func NewProductRepository(client *mongo.Client, dbName string, rabbitMQConn *messagingQueue.Connection) productRepository.ProductRepository {
	return &productRepositoryImpl{
		client:       client,
		dbName:       dbName,
		rabbitMQConn: rabbitMQConn,
	}
}

func (r *productRepositoryImpl) getCollection() *mongo.Collection {
	return r.client.Database(r.dbName).Collection("products")
}
