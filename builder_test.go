package builder_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/assert.v1"
)

var db *mongo.Database

const dbCollName = "_filter_builder_test"

func setup() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	d := client.Database(dbCollName)
	db = d
	err = d.Drop(ctx)
	if err != nil {
		panic(err)
	}

	coll := d.Collection(dbCollName)

	_, err = coll.InsertMany(ctx, []interface{}{
		bson.M{"name": "a", "capName": "A", "age": 8, "birthdate": time.Now()},
		bson.M{"name": "aa", "capName": "AA", "age": 58, "birthdate": time.Now()},
		bson.M{"name": "b", "capName": "B", "age": 18, "birthdate": time.Now().Add(24 * time.Hour)},
		bson.M{"name": "bb", "capName": "BB", "age": 18, "birthdate": time.Now().Add(24 * time.Hour)},
		bson.M{"name": "c", "capName": "C", "age": 38, "birthdate": time.Now().Add(3 * 24 * time.Hour)},
		bson.M{"name": "d", "capName": "D", "age": 28, "birthdate": time.Now().Add(6 * 24 * time.Hour)},
		bson.M{"name": "ab", "capName": "AB", "age": 5, "birthdate": time.Now()},
	})
	if err != nil {
		panic(err)
	}
}

func tearDown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.Drop(ctx)
	if err != nil {
		panic(err)
	}
}

func fetchData(filter interface{}) []interface{} {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := db.Collection(dbCollName).Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	defer c.Close(ctx)
	var res []interface{}
	for c.Next(ctx) {
		var result bson.D
		err := c.Decode(&result)
		if err != nil {
			panic(err)
		}
		res = append(res, result)
	}
	if err := c.Err(); err != nil {
		panic(err)
	}
	return res
}

func mustEqual(t *testing.T, val1, val2 interface{}) {
	res1, res2 := fetchData(val1), fetchData(val2)
	// t.Logf("\n%v\n%v\n", res1, res2)
	assert.Equal(t, reflect.DeepEqual(res1, res2), true)
}
