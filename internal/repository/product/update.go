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

func (r *productRepositoryImpl) Update(productData globalModel.Product) (product globalModel.Product, err error) {
	prdObjId, err := primitive.ObjectIDFromHex(productData.Id)
	if err != nil {
		return product, err
	}

	var productDoc productSchema.ProductSchema
	updatedAttributes := bson.D{
		{"name", productData.Name},
		{"description", productData.Description},
		{"color", productData.Color},
		{"price", productData.Price},
		{"imageUrl", productData.ImageURL},
		{"updatedAt", time.Now()},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.getCollection().FindOneAndUpdate(nil,
		bson.D{{"_id", prdObjId}},
		bson.D{{"$set", updatedAttributes}},
		opts).Decode(&productDoc)
	if err == mongo.ErrNoDocuments {
		return product, nil
	} else if err != nil {
		return product, err
	}
	product = globalModel.Product{
		Id:          productDoc.Id.Hex(),
		Name:        productDoc.Name,
		Description: productDoc.Description,
		Color:       productDoc.Color,
		Price:       productDoc.Price,
		ImageURL:    productDoc.ImageURL,
		CreatedAt:   productDoc.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt:   productDoc.UpdatedAt.Format(time.RFC3339Nano),
	}

	return product, nil
}
