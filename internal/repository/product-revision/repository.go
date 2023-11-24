package productRevisionRepositoryImpl

import (
	productRevisionDomain "github.com/DaniaLD/upera-go-test/internal/domain/product-revision"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRevisionRepositoryImpl struct {
	mongoClient *mongo.Client
	dbName      string
	redisClient *redis.Client
}

func NewProductRevisionRepository(mongoClient *mongo.Client, dbName string, redisClient *redis.Client) productRevisionDomain.ProductRevisionRepository {
	return &productRevisionRepositoryImpl{
		mongoClient: mongoClient,
		dbName:      dbName,
		redisClient: redisClient,
	}
}

func (r *productRevisionRepositoryImpl) getCollection() *mongo.Collection {
	return r.mongoClient.Database(r.dbName).Collection("productRevisions")
}
