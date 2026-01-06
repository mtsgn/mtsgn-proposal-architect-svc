package utils

import (
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertToBson(entity interface{}) (bson.M, error) {
	if entity == nil {
		return bson.M{}, nil
	}

	sel, err := bson.Marshal(entity)
	if err != nil {
		return nil, err
	}

	obj := bson.M{}
	bson.Unmarshal(sel, &obj)

	return obj, nil
}
