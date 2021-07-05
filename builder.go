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

// Flush restes the builder to initial state
func (b *Builder) Flush() *Builder {
	b.condMaps = []bson.M{}
	b.curMap = bson.M{}
	return b
}

// WantNum indicates the builder to build a condtion for string type.
func (b *Builder) Str(key string) *strCond {
	return newStrCond(key, b)
}

// Num indicates the builder to build a condtion for number type.
func (b *Builder) Num(key string) *numCond {
	return newNumCond(key, b)
}

func (b *Builder) Date(key string, defaultFormat ...string) *dateCond {
	return newDateCond(key, b, defaultFormat...)
}

func (b *Builder) Oid() *oidCond {
	return newOidCond(b)
}

// Any constructs a condition without type restricted.
func (b *Builder) Any(key string) *cond {
	return newCond(key, b)
}

// Build builds final filter and returns it.
// Build usually should be only called once for every build since it will call b.Flush() after build up the final map.
func (b *Builder) Build() bson.M {
	if len(b.curMap) != 0 {
		b.condMaps = append(b.condMaps, b.curMap)
	}
	if len(b.condMaps) > 1 {

		return bson.M{_or: b.condMaps}
	}
	res := b.curMap
	b.Flush()
	return res
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

// AnyMap will set the given map to current condition.
func (b *Builder) AnyMap(key string, m bson.M) *Builder {
	b.curMap[key] = m
	return b
}
