package productRepositoryImpl

import (
	productSchema "github.com/DaniaLD/upera-go-test/internal/repository/schema/product"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (r *productRepositoryImpl) Update(productData globalModel.Product) (orgProduct globalModel.Product, updateTime string, err error) {
	prdObjId, err := primitive.ObjectIDFromHex(productData.Id)
	if err != nil {
		return orgProduct, "", err
	}

	now := time.Now()
	var orgProductDoc productSchema.ProductSchema
	updatedAttributes := bson.D{
		{"name", productData.Name},
		{"description", productData.Description},
		{"color", productData.Color},
		{"price", productData.Price},
		{"imageUrl", productData.ImageURL},
		{"updatedAt", now},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.Before)
	err = r.getCollection().FindOneAndUpdate(nil,
		bson.D{{"_id", prdObjId}},
		bson.D{{"$set", updatedAttributes}},
		opts).Decode(&orgProductDoc)
	if err == mongo.ErrNoDocuments {
		return orgProduct, "", nil
	} else if err != nil {
		return orgProduct, "", err
	}
	location, _ := time.LoadLocation("Asia/Tehran")
	orgProduct = globalModel.Product{
		Id:          productData.Id,
		Name:        orgProductDoc.Name,
		Description: orgProductDoc.Description,
		Color:       orgProductDoc.Color,
		Price:       orgProductDoc.Price,
		ImageURL:    orgProductDoc.ImageURL,
		CreatedAt:   orgProductDoc.CreatedAt.In(location).Format(time.RFC3339Nano),
		UpdatedAt:   orgProductDoc.UpdatedAt.In(location).Format(time.RFC3339Nano),
	}

	return orgProduct, now.Format(time.RFC3339Nano), nil
}
