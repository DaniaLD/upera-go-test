package productRevisionRepositoryImpl

import (
	"context"
	"encoding/json"
	"fmt"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (r *productRevisionRepositoryImpl) Create(
	productId string,
	previousData globalModel.Product,
	newData globalModel.Product,
	updatedAttributes []string,
) error {
	prdObjId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}

	previousDataJson, _ := json.Marshal(previousData)
	newDataJson, _ := json.Marshal(newData)

	revisionId := r.getProductLatestRevisionNum(productId)
	now := time.Now()
	doc := bson.D{
		{"productId", prdObjId},
		{"revisionId", revisionId},
		{"updatedAttributes", updatedAttributes},
		{"previousValue", string(previousDataJson)},
		{"newValue", string(newDataJson)},
		{"createdAt", now},
	}
	_, err = r.getCollection().InsertOne(nil, doc)
	if err != nil {
		return err
	}
	r.setProductLatestRevisionNum(productId, revisionId)

	return nil
}

func (r *productRevisionRepositoryImpl) getProductLatestRevisionNum(productId string) int64 {
	ctx := context.Background()
	key := fmt.Sprintf("product:%s:revision", productId)
	rvs, err := r.redisClient.Get(ctx, key).Int64()
	if err == redis.Nil {
		r.setProductLatestRevisionNum(productId, 0)
		return 1
	}
	return rvs
}

func (r *productRevisionRepositoryImpl) setProductLatestRevisionNum(productId string, revisionId int64) {
	ctx := context.Background()
	key := fmt.Sprintf("product:%s:revision", productId)
	err := r.redisClient.Incr(ctx, key).Err()
	// fallback for handling redis error
	if err != nil {
		revisionId++
		r.redisClient.Set(ctx, key, revisionId, 0)
	}
}
