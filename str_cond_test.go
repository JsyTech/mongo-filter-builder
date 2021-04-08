package builder_test

import (
	"testing"

	builder "github.com/LegeTech/mongo-filter-builder"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBuilderStr(t *testing.T) {
	setup()
	defer tearDown()
	b := builder.New()
	singleCondFilter := b.Str("name").Eq("a").Build()
	caze := bson.M{
		"name": bson.M{
			"$eq": "a",
		},
	}
	mustEqual(t, singleCondFilter, caze)

	b2 := builder.New()

	MultiCondFilter := b2.
		Str("name").Eq("a").
		Str("capName").Eq("B").Build()

	caze2 := bson.M{
		"name":    bson.M{"$eq": "a"},
		"capName": bson.M{"$eq": "B"},
	}

	mustEqual(t, MultiCondFilter, caze2)
}

func TestStrOr(t *testing.T) {
	setup()
	defer tearDown()

	b := builder.New()

	f := b.Str("name").Eq("a").
		Or().
		Str("name").Eq("b").Build()

	caze := bson.M{
		"$or": []bson.M{
			{"name": "a"},
			{"name": "b"},
		},
	}

	mustEqual(t, f, caze)
}

func TestStrRegex(t *testing.T) {
	setup()
	defer tearDown()
	f := builder.New().Str("name").Like("a").Build()

	caze := bson.M{
		"name": bson.M{"$regex": "a"},
	}
	mustEqual(t, f, caze)

	f = builder.New().Str("name").NotLike("a").Build()
	caze = bson.M{
		"name": bson.M{"$not": primitive.Regex{Pattern: "a"}},
	}
	mustEqual(t, f, caze)
}
