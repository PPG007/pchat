package repository

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
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
	model := options.IndexModel{}
	if index.IsUnique {
		model.SetUnique(true)
	}
	if len(index.PartialExpression) > 0 {
		model.SetPartialFilterExpression(index.PartialExpression)
	}
	var key []string
	for _, field := range index.Fields {
		order := ""
		if field.Desc {
			order = "-"
		}
		key = append(key, fmt.Sprintf("%s%s", order, field.Name))
	}
	model.Key = key
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

func Upsert(ctx context.Context, collection string, condition, updater bson.M) error {
	_, err := mongoClient.Database.Collection(collection).Upsert(ctx, condition, updater)
	return err
}
