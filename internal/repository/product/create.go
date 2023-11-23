package productRepositoryImpl

import (
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (r *productRepositoryImpl) Create(productData globalModel.Product) (product globalModel.Product, err error) {
	doc := bson.D{
		{"name", productData.Name},
		{"description", productData.Description},
		{"color", productData.Color},
		{"price", productData.Price},
		{"imageUrl", productData.ImageURL},
		{"createdAt", time.Now()},
		{"updatedAt", time.Now()},
	}
	result, err := r.getCollection().InsertOne(nil, doc)
	if err != nil {
		return product, err
	}
	productObjId := result.InsertedID.(primitive.ObjectID).Hex()
	product = productData
	product.Id = productObjId

	return product, nil
}
