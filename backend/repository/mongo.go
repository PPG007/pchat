package repository

import (
	"context"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"github.com/spf13/viper"
	"pchat/repository/bson"
)

var mongoClient *qmgo.QmgoClient

func init() {
	ctx := context.Background()
	client, err := qmgo.Open(ctx, &qmgo.Config{
		Uri:      viper.GetString("mongo.uri"),
		Database: viper.GetString("mongo.database"),
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

func CreateIndex(ctx context.Context, collection string, index options.IndexModel) error {
	return mongoClient.Database.Collection(collection).CreateOneIndex(ctx, index)
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

func FindAllWithPage(ctx context.Context, collection string, sorter []string, page, perPage int64, condition bson.M, result interface{}) (int64, error) {
	col := mongoClient.Database.Collection(collection)
	err := col.Find(ctx, condition).Sort(sorter...).Skip((page - 1) * perPage).Limit(perPage).All(result)
	if err != nil {
		return 0, err
	}
	return col.Find(ctx, condition).Count()
}
