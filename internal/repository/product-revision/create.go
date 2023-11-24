package productRevisionRepositoryImpl

import (
	"encoding/json"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
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

	now := time.Now()
	doc := bson.D{
		{"productId", prdObjId},
		{"revisionId", 1},
		{"updatedAttributes", updatedAttributes},
		{"previousValue", string(previousDataJson)},
		{"newValue", string(newDataJson)},
		{"createdAt", now},
	}
	_, err = r.getCollection().InsertOne(nil, doc)
	if err != nil {
		return err
	}

	return nil
}
