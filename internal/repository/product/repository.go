package productRepositoryImpl

import (
	productRepository "github.com/DaniaLD/upera-go-test/internal/domain/product"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepositoryImpl struct {
	name   string
	client *mongo.Client
	dbName string
}

func NewProductRepository(client *mongo.Client, dbName string) productRepository.ProductRepository {
	return &productRepositoryImpl{
		client: client,
		dbName: dbName,
	}
}

func (r *productRepositoryImpl) getCollection() *mongo.Collection {
	return r.client.Database(r.dbName).Collection("products")
}
