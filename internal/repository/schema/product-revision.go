package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductRevision struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	ProductId         primitive.ObjectID `json:"productId" bson:"productId"`
	RevisionId        int64              `json:"revisionId" bson:"revisionId"`
	UpdatedAttributes []string           `json:"updatedAttributes" bson:"updatedAttributes"`
	PreviousValue     string             `json:"previousValue" bson:"previousValue"`
	NewValue          string             `json:"newValue" bson:"newValue"`
	CreatedAt         time.Time          `json:"createdAt" bson:"createdAt"`
}
