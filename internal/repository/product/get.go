package productRepositoryImpl

import (
	"github.com/DaniaLD/upera-go-test/internal/repository/schema"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (r *productRepositoryImpl) Get(productId string) (product globalModel.Product, err error) {
	prdObjId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return product, err
	}

	var prdDoc schema.Product
	err = r.getCollection().FindOne(nil, bson.D{
		{"_id", prdObjId},
	}).Decode(&prdDoc)
	if err != nil {
		return product, err
	}

	location, _ := time.LoadLocation("Asia/Tehran")
	product = globalModel.Product{
		Id:          prdDoc.Id.Hex(),
		Name:        prdDoc.Name,
		Description: prdDoc.Description,
		Color:       prdDoc.Color,
		Price:       prdDoc.Price,
		ImageURL:    prdDoc.ImageURL,
		CreatedAt:   prdDoc.CreatedAt.In(location).Format(time.RFC3339Nano),
		UpdatedAt:   prdDoc.UpdatedAt.In(location).Format(time.RFC3339Nano),
	}

	return product, err
}
