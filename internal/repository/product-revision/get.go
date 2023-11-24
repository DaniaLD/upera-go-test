package productRevisionRepositoryImpl

import (
	"encoding/json"
	productRevisionSchema "github.com/DaniaLD/upera-go-test/internal/repository/schema"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

func (r *productRevisionRepositoryImpl) Get(productId string, limit int64, skip int64) (revisions []globalModel.ProductRevision, total int64, err error) {
	prdObjId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return revisions, total, err
	}

	rvsCh := make(chan []productRevisionSchema.ProductRevision, limit)
	rvsErrCh := make(chan error, 1)
	totalCh := make(chan int64, 1)
	totalErrCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		opts := options.Find().SetLimit(limit).SetSkip(skip)
		cur, err := r.getCollection().Find(nil, bson.D{
			{"productId", prdObjId},
		}, opts)
		if err != nil {
			rvsErrCh <- err
		}
		var revisionDocs []productRevisionSchema.ProductRevision
		_ = cur.All(nil, &revisionDocs)
		rvsCh <- revisionDocs
	}()
	go func() {
		defer wg.Done()
		total, err := r.getCollection().CountDocuments(nil, bson.D{{"productId", prdObjId}})
		if err != nil {
			totalErrCh <- err
		}
		totalCh <- total
	}()
	select {
	case rvsDocs := <-rvsCh:
		revisions = revisionsMapper(rvsDocs)
	case err = <-rvsErrCh:
		return revisions, total, err
	}
	select {
	case total = <-totalCh:

	case err = <-totalErrCh:
		return revisions, total, err
	}
	wg.Wait()
	return revisions, total, err
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
