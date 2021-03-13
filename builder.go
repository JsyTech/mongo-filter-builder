package builder

import (
	"go.mongodb.org/mongo-driver/bson"
)

const (
	_eq = "$eq"
	_ne = "$ne"

	_gt  = "$gt"
	_gte = "$gte"
	_lt  = "$lt"
	_lte = "$lte"

	_in  = "$in"
	_nin = "nin"

	_regex = "$regex"
	_not   = "$not"

	_or = "$or"
)

// Builder represents a filter builder.
type Builder struct {
	// condMaps stores all condtion maps.
	condMaps []bson.M
	// curMap represents the currently operated condtion map.
	// A condtion map can be a single element map also can be a multiple elements map.
	curMap bson.M
}

// New constructs a new Builder.
func New() *Builder {
	maps := []bson.M{}
	return &Builder{
		condMaps: maps,
		curMap:   bson.M{},
	}
}

// WantNum indicates the builder to build a condtion for string type.
func (b *Builder) Str(key string) *strCond {
	return newStrCond(key, b)
}

// Num indicates the builder to build a condtion for number type.
func (b *Builder) Num(key string) interface{} {

	return nil
}

// Build builds final filter and returns it.
func (b *Builder) Build() bson.M {
	if len(b.curMap) != 0 {
		b.condMaps = append(b.condMaps, b.curMap)
	}
	if len(b.condMaps) > 1 {

		return bson.M{_or: b.condMaps}
	}
	return b.curMap
}

// Or appends b.curMap to b.condMaps, and b.curMap will be assigned to a new empty map.
// Thus if finally b.condMaps's len is bigger than 1, then the final filter will wraps all maps into a $or condtion.
func (b *Builder) Or() *Builder {
	if len(b.curMap) == 0 {
		return b
	}
	b.condMaps = append(b.condMaps, b.curMap)
	b.curMap = bson.M{}
	return b
}
