package productSchema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductSchema struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Color       string             `json:"color" bson:"color"`
	Price       int64              `json:"price" bson:"price"`
	ImageURL    string             `json:"imageURL" bson:"imageURL"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}
