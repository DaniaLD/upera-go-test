package productRevisionRepositoryImpl

import (
	"encoding/json"
	"github.com/DaniaLD/upera-go-test/internal/repository/schema"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (r *productRevisionRepositoryImpl) GetByRevisionId(productId string, revisionId int64) (product globalModel.Product, err error) {
	prdObjId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return product, err
	}

	var prdRvsDoc schema.ProductRevision
	err = r.getCollection().FindOne(nil, bson.D{
		{"productId", prdObjId},
		{"revisionId", revisionId},
	}).Decode(&prdRvsDoc)
	if err != nil {
		return product, err
	}

	var prdDoc schema.Product
	_ = json.Unmarshal([]byte(prdRvsDoc.NewValue), &prdDoc)
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
