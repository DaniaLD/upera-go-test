package productRevisionRepositoryImpl

import (
	productRevisionDomain "github.com/DaniaLD/upera-go-test/internal/domain/product-revision"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRevisionRepositoryImpl struct {
	client *mongo.Client
	dbName string
}

func NewProductRevisionRepository(client *mongo.Client, dbName string) productRevisionDomain.ProductRevisionRepository {
	return &productRevisionRepositoryImpl{
		client: client,
		dbName: dbName,
	}
}

func (r *productRevisionRepositoryImpl) getCollection() *mongo.Collection {
	return r.client.Database(r.dbName).Collection("productRevisions")
}
