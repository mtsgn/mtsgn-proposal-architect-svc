package database

import (
	"boilerplate-api/internal/common/utils"
	"boilerplate-api/pkg/config"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Instance struct {
	db         *mongo.Database
	collection *mongo.Collection

	CollectionName string
}

var CollectionNotConnected = fmt.Errorf("collection not connected")

func InitMongoDB(cfg config.DatabaseConfig) (*mongo.Database, error) {
	// Use admin as the authentication database
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client.Database(cfg.DBName), nil
}

func (ins *Instance) SetInstance(db *mongo.Database) *Instance {
	ins.db = db
	ins.collection = db.Collection(ins.CollectionName)
	return ins
}

func (ins *Instance) CreateIndex(keys bson.D, options *options.IndexOptions) error {
	_, err := ins.collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    keys,
		Options: options,
	})
	return err
}

func (ins *Instance) Create(entity interface{}) (string, error) {
	if ins.collection == nil {
		return "", CollectionNotConnected
	}

	obj, err := utils.ConvertToBson(entity)
	if err != nil {
		return "", err
	}

	if obj["created_at"] == nil {
		obj["created_at"] = time.Now()
	}

	result, err := ins.collection.InsertOne(context.TODO(), obj)
	if err != nil {
		return "", err
	}

	//obj["_id"] = result.InsertedID

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (ins *Instance) Query(query interface{}, offset int64, limit int64, sortFields *bson.M) (*mongo.Cursor, error) {
	if ins.collection == nil {
		return nil, errors.New("Collection not connected")
	}

	opt := &options.FindOptions{}
	k := int64(1000)
	if limit <= 0 {
		opt.Limit = &k
	} else {
		opt.Limit = &limit
	}
	if offset > 0 {
		opt.Skip = &offset
	}
	if sortFields != nil {
		opt.Sort = sortFields
	}

	var result *mongo.Cursor
	var err error

	switch query.(type) {
	case nil, bson.M, bson.D, map[string]interface{}:
		result, err = ins.collection.Find(context.TODO(), query, opt)
	default:
		converted, err := utils.ConvertToBson(query)
		if err != nil {
			return nil, err
		}
		result, err = ins.collection.Find(context.TODO(), converted, opt)
	}

	if err != nil {
		return nil, err
	}
	if result != nil && result.Err() != nil {
		return nil, result.Err()
	}

	return result, nil
}

func (ins *Instance) QueryOne(query interface{}) (*mongo.SingleResult, error) {
	if ins.collection == nil {
		return nil, CollectionNotConnected
	}

	converted, err := utils.ConvertToBson(query)
	if err != nil {
		return nil, err
	}

	result := ins.collection.FindOne(context.TODO(), converted)

	if result == nil {
		return nil, fmt.Errorf("not found any matched %s", ins.CollectionName)
	}
	if result.Err() != nil {
		return nil, fmt.Errorf("error detail: %w", result.Err())
	}

	//if result == nil || result.Err() != nil {
	//	errMsg := "Not found any matched " + ins.CollectionName + "."
	//	if result != nil && result.Err() != nil {
	//		errMsg += " Error detail: " + result.Err().Error()
	//	}
	//	return nil, errors.New(errMsg)
	//}
	return result, nil
}

func (ins *Instance) UpdateOne(query interface{}, updater interface{}, opts ...*options.FindOneAndUpdateOptions) (*mongo.SingleResult, error) {
	// check col
	if ins.collection == nil {
		return nil, CollectionNotConnected
	}

	// convert
	bUpdater, err := utils.ConvertToBson(updater)
	if err != nil {
		return nil, err
	}
	// bUpdater["updated_at"] = time.Now()

	// transform to bson
	converted, err := utils.ConvertToBson(query)
	if err != nil {
		return nil, err
	}

	// do update
	if opts == nil {
		after := options.After
		opts = []*options.FindOneAndUpdateOptions{
			{
				ReturnDocument: &after,
			},
		}
	}

	result := ins.collection.FindOneAndUpdate(context.TODO(), converted, bUpdater, opts...)
	if result.Err() != nil {
		//detail := result.Err().Error()
		//return nil, errors.New("Not found any matched " + ins.CollectionName + ". Error detail: " + detail)
		return nil, fmt.Errorf("not found any matched %s. Error detail: %w", ins.CollectionName, result.Err())
	}

	return result, nil
}

func (ins *Instance) Aggregate(pipeline interface{}) (*mongo.Cursor, error) {
	// check col
	if ins.collection == nil {
		return nil, CollectionNotConnected
	}

	cur, err := ins.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	if cur.Err() != nil {
		return nil, cur.Err()
	}

	return cur, nil
}

func (ins *Instance) Count(query interface{}) (int64, error) {
	// check col
	if ins.collection == nil {
		return 0, CollectionNotConnected
	}

	// convert query
	converted, err := utils.ConvertToBson(query)
	if err != nil {
		return 0, err
	}

	// if query is empty -> count by EstimatedDocumentCount else count by CountDocuments
	count := int64(0)
	if len(converted) == 0 {
		count, err = ins.collection.EstimatedDocumentCount(context.TODO(), nil)
	} else {
		count, err = ins.collection.CountDocuments(context.TODO(), converted)
	}
	if err != nil {
		return 0, err
	}

	return count, nil

}

func (ins *Instance) QuerySpecificFields(ctx context.Context, filter bson.M, fields bson.M, result interface{}, sort ...bson.M) error {
	if ins.db == nil {
		return CollectionNotConnected
	}

	opts := options.FindOne().SetProjection(fields)

	if len(sort) > 0 {
		opts.SetSort(sort[0])
	}

	err := ins.collection.FindOne(ctx, filter, opts).Decode(result)
	return err
}
