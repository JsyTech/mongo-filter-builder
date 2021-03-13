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

func NewCond(key string, builder *Builder) *cond {
	return &cond{
		key:     key,
		m:       bson.M{},
		builder: builder,
	}
}

// addMapToBuilder adds baseCond.m to the referenced builder's map with key as baseCond.key.
func (baseCond *cond) addMapToBuilder() {
	baseCond.builder.curMap[baseCond.key] = baseCond.m
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
