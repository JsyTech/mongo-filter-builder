package builder

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type cond struct {
	// key is the root name for the current condtion, eq: {key: {$eq: ...}}.
	key string
	// TODO consider remove the m, populate the builder.curMap directly.
	// m is the current map stores the value of condition.
	m bson.M
	// builder refers to the current builder.
	builder *Builder
}

func newCond(key string, builder *Builder) *cond {
	return &cond{
		key:     key,
		m:       bson.M{},
		builder: builder,
	}
}

// addMapToBuilder adds baseCond.m to the referenced builder's map with key as baseCond.key.
// If the same key is set again, it will try to merge two map.
func (baseCond *cond) addMapToBuilder() {
	var v interface{}
	var ok bool

	if v, ok = baseCond.builder.curMap[baseCond.key]; !ok {
		baseCond.builder.curMap[baseCond.key] = baseCond.m
		return
	}

	preMap := v.(map[string]interface{})

	for k, v := range baseCond.m {
		preMap[k] = v
	}
}

// eq adds `$eq: val` to the baseCond.m
func (baseCond *cond) eq(val interface{}) {
	baseCond.m[_eq] = val
	baseCond.addMapToBuilder()
}

// ne adds `$ne: val` to the baseCond.m
func (baseCond *cond) ne(val interface{}) {
	baseCond.m[_ne] = val
	baseCond.addMapToBuilder()
}

// lt adds `$lt: val` to the baseCond.m
func (baseCond *cond) lt(val interface{}) {
	baseCond.m[_lt] = val
	baseCond.addMapToBuilder()
}

// lte adds `$lte: val` to the baseCond.m
func (baseCond *cond) lte(val interface{}) {
	baseCond.m[_lte] = val
	baseCond.addMapToBuilder()
}

// gt adds `$gt: val` to the baseCond.m
func (baseCond *cond) gt(val interface{}) {
	baseCond.m[_gt] = val
	baseCond.addMapToBuilder()
}

// gte adds `$gte: val` to the baseCond.m
func (baseCond *cond) gte(val interface{}) {
	baseCond.m[_gte] = val
	baseCond.addMapToBuilder()
}

// regex adds `$regex: exp, $options: ""` to the baseCond.m
func (baseCond *cond) regex(exp string) {
	baseCond.regexWithOpt(exp, "")
}

// regexWithOpt adds `$regex: exp, $options: opt` to the baseCond.m
func (baseCond *cond) regexWithOpt(exp string, opt string) {
	baseCond.m[_regex] = primitive.Regex{Pattern: exp, Options: opt}
	baseCond.addMapToBuilder()
}

// not adds `$not: exp, $options: ""` to the baseCond.m
func (baseCond *cond) not(exp string) {
	baseCond.notWithOpt(exp, "")
}

// notWithOpt adds `$not: exp, $options: opt` to the baseCond.m
func (baseCond *cond) notWithOpt(exp string, opt string) {
	baseCond.m[_not] = primitive.Regex{Pattern: exp, Options: opt}
	baseCond.addMapToBuilder()
}

// in adds `$in: vals` to the baseCond.m
func (baseCond *cond) in(vals interface{}) {
	baseCond.m[_in] = vals
	baseCond.addMapToBuilder()
}

// nin adds `$nin: vals` to the baseCond.m
func (baseCond *cond) nin(vals interface{}) {
	baseCond.m[_nin] = vals
	baseCond.addMapToBuilder()
}
