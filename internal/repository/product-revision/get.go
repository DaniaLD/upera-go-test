package productRevisionRepositoryImpl

import (
	"encoding/json"
	productRevisionSchema "github.com/DaniaLD/upera-go-test/internal/repository/schema/product-revision"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (r *productRevisionRepositoryImpl) Get(productId string) (revisions []globalModel.ProductRevision, err error) {
	prdObjId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return revisions, err
	}

	cur, err := r.getCollection().Find(nil, bson.D{
		{"productId", prdObjId},
	})
	if err != nil {
		return revisions, err
	}
	var revisionDocs []productRevisionSchema.ProductRevision
	_ = cur.All(nil, &revisionDocs)

	return revisionsMapper(revisionDocs), err
}

func revisionsMapper(revisionDocs []productRevisionSchema.ProductRevision) (revisions []globalModel.ProductRevision) {
	location, _ := time.LoadLocation("Asia/Tehran")
	for _, doc := range revisionDocs {
		revision := globalModel.ProductRevision{
			Id:                doc.Id.Hex(),
			ProductId:         doc.ProductId.Hex(),
			RevisionId:        doc.RevisionId,
			UpdatedAttributes: doc.UpdatedAttributes,
			CreatedAt:         doc.CreatedAt.In(location).Format(time.RFC3339Nano),
		}
		_ = json.Unmarshal([]byte(doc.PreviousValue), &revision.PreviousValue)
		_ = json.Unmarshal([]byte(doc.NewValue), &revision.NewValue)
		revisions = append(revisions, revision)
	}
	return revisions
}
