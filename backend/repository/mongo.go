package repository

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	mgo_options "go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"pchat/repository/bson"
)

var mongoClient *qmgo.QmgoClient

type Pagination struct {
	Condition bson.M
	Page      int64
	PerPage   int64
	OrderBy   []string
}

type IndexField struct {
	Name string
	Desc bool
}

type IndexOption struct {
	Fields            []IndexField
	IsUnique          bool
	PartialExpression bson.M
}

func init() {
	ctx := context.Background()
	client, err := qmgo.Open(ctx, &qmgo.Config{
		Uri:      os.Getenv("MONGO_URI"),
		Database: os.Getenv("MONGO_DATABASE"),
	})
	if err != nil {
		panic(err)
	}
	mongoClient = client
}

func Insert(ctx context.Context, collection string, docs ...interface{}) error {
	col := mongoClient.Database.Collection(collection)
	var err error
	if len(docs) == 1 {
		_, err = col.InsertOne(ctx, docs[0])
	} else {
		_, err = col.InsertMany(ctx, docs)
	}
	return err
}

func UpdateOne(ctx context.Context, collection string, condition bson.M, updater bson.M) error {
	return mongoClient.Database.Collection(collection).UpdateOne(ctx, condition, updater)
}

func FindAll(ctx context.Context, collection string, condition bson.M, result interface{}) error {
	return mongoClient.Database.Collection(collection).Find(ctx, condition).All(result)
}

func FindOne(ctx context.Context, collection string, condition bson.M, result interface{}) error {
	return mongoClient.Database.Collection(collection).Find(ctx, condition).One(result)
}

func Count(ctx context.Context, collection string, condition bson.M) (int64, error) {
	return mongoClient.Database.Collection(collection).Find(ctx, condition).Count()
}

func FindAndApply(ctx context.Context, collection string, condition bson.M, change qmgo.Change, result interface{}) error {
	return mongoClient.Database.Collection(collection).Find(ctx, condition).Apply(change, result)
}

func CreateIndex(ctx context.Context, collection string, index IndexOption) error {
	var key []string
	for _, field := range index.Fields {
		order := ""
		if field.Desc {
			order = "-"
		}
		key = append(key, fmt.Sprintf("%s%s", order, field.Name))
	}
	opts := mgo_options.Index()
	if index.IsUnique {
		opts.SetUnique(true)
	}
	if len(index.PartialExpression) > 0 {
		opts.SetPartialFilterExpression(index.PartialExpression)
	}
	model := options.IndexModel{
		Key:          key,
		IndexOptions: opts,
	}
	return mongoClient.Database.Collection(collection).CreateOneIndex(ctx, model)
}

func FindOneWithSorter(ctx context.Context, collection string, sorter []string, condition bson.M, result interface{}) error {
	return mongoClient.Database.Collection(collection).Find(ctx, condition).Sort(sorter...).One(result)
}

func UpdateAll(ctx context.Context, collection string, condition, updater bson.M) (int64, error) {
	result, err := mongoClient.Database.Collection(collection).UpdateAll(ctx, condition, updater)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func FindAllWithSorter(ctx context.Context, collection string, sorter []string, condition bson.M, result interface{}) error {
	return mongoClient.Database.Collection(collection).Find(ctx, condition).Sort(sorter...).All(result)
}

func FindAllWithPage(ctx context.Context, collection string, pagination Pagination, result interface{}) (int64, error) {
	col := mongoClient.Database.Collection(collection)
	err := col.Find(ctx, pagination.Condition).
		Sort(pagination.OrderBy...).
		Skip((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		All(result)
	if err != nil {
		return 0, err
	}
	return col.Find(ctx, pagination.Condition).Count()
}

func Transaction(ctx context.Context, fn func(context.Context) error) error {
	session, err := mongoClient.Session()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	_, err = session.StartTransaction(ctx, func(sessCtx context.Context) (interface{}, error) {
		return nil, fn(sessCtx)
	})
	return err
}

type BatchUpdater struct {
	Condition bson.M
	Updater   bson.M
}

func BatchUpdate(ctx context.Context, collection string, updater ...BatchUpdater) error {
	bulk := mongoClient.Database.Collection(collection).Bulk()
	for _, batchUpdater := range updater {
		bulk = bulk.UpdateOne(batchUpdater.Condition, batchUpdater.Updater)
	}
	_, err := bulk.Run(ctx)
	return err
}

func BatchUpdateUnOrdered(ctx context.Context, collection string, updater ...BatchUpdater) error {
	bulk := mongoClient.Database.Collection(collection).Bulk().SetOrdered(false)
	for _, batchUpdater := range updater {
		bulk = bulk.UpdateOne(batchUpdater.Condition, batchUpdater.Updater)
	}
	_, err := bulk.Run(ctx)
	return err
}

func BatchUpsert(ctx context.Context, collection string, updater ...BatchUpdater) error {
	bulk := mongoClient.Database.Collection(collection).Bulk()
	for _, batchUpdater := range updater {
		bulk = bulk.UpsertOne(batchUpdater.Condition, batchUpdater.Updater)
	}
	_, err := bulk.Run(ctx)
	return err
}

func BatchUpsertUnOrdered(ctx context.Context, collection string, updater ...BatchUpdater) error {
	bulk := mongoClient.Database.Collection(collection).Bulk().SetOrdered(false)
	for _, batchUpdater := range updater {
		bulk = bulk.UpsertOne(batchUpdater.Condition, batchUpdater.Updater)
	}
	_, err := bulk.Run(ctx)
	return err
}

func BatchInsert(ctx context.Context, collection string, docs ...interface{}) error {
	bulk := mongoClient.Database.Collection(collection).Bulk()
	for _, doc := range docs {
		bulk = bulk.InsertOne(doc)
	}
	_, err := bulk.Run(ctx)
	return err
}

func BatchInsertUnOrdered(ctx context.Context, collection string, docs ...interface{}) error {
	bulk := mongoClient.Database.Collection(collection).Bulk().SetOrdered(false)
	for _, doc := range docs {
		bulk = bulk.InsertOne(doc)
	}
	_, err := bulk.Run(ctx)
	return err
}

func Upsert(ctx context.Context, collection string, condition, updater bson.M) error {
	_, err := mongoClient.Database.Collection(collection).Upsert(ctx, condition, updater)
	return err
}

func Aggregate(ctx context.Context, collection string, pipeline []bson.M, result interface{}) error {
	col := mongoClient.Database.Collection(collection)
	return col.Aggregate(ctx, pipeline).All(&result)
}

func Distinct(ctx context.Context, collection string, condition bson.M, key string, result interface{}) error {
	return mongoClient.Database.Collection(collection).Find(ctx, condition).Distinct(key, &result)
}

func Iterate(ctx context.Context, collection string, condition bson.M, batchSize int64, sorter []string, fields bson.M) (qmgo.CursorI, error) {
	cursor := mongoClient.
		Database.
		Collection(collection).
		Find(ctx, condition).
		BatchSize(batchSize).
		NoCursorTimeout(true).
		Sort(sorter...).
		Select(fields).Cursor()
	err := cursor.Err()
	if err != nil {
		return nil, err
	}
	return cursor, nil
}
