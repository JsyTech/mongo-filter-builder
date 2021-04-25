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
	preMap := v.(bson.M)

	for k, v := range baseCond.m {
		preMap[k] = v
	}
}

// Eq adds `$Eq: val` to the baseCond.m
func (baseCond *cond) Eq(val interface{}) {
	baseCond.m[_eq] = val
	baseCond.addMapToBuilder()
}

// Ne adds `$Ne: val` to the baseCond.m
func (baseCond *cond) Ne(val interface{}) {
	baseCond.m[_ne] = val
	baseCond.addMapToBuilder()
}

// Lt adds `$Lt: val` to the baseCond.m
func (baseCond *cond) Lt(val interface{}) {
	baseCond.m[_lt] = val
	baseCond.addMapToBuilder()
}

// Lte adds `$Lte: val` to the baseCond.m
func (baseCond *cond) Lte(val interface{}) {
	baseCond.m[_lte] = val
	baseCond.addMapToBuilder()
}

// gt adds `$gt: val` to the baseCond.m
func (baseCond *cond) gt(val interface{}) {
	baseCond.m[_gt] = val
	baseCond.addMapToBuilder()
}

// Gte adds `$Gte: val` to the baseCond.m
func (baseCond *cond) Gte(val interface{}) {
	baseCond.m[_gte] = val
	baseCond.addMapToBuilder()
}

// Regex adds `$Regex: exp, $options: ""` to the baseCond.m
func (baseCond *cond) Regex(exp string) {
	baseCond.RegexWithOpt(exp, "")
}

// RegexWithOpt adds `$regex: exp, $options: opt` to the baseCond.m
func (baseCond *cond) RegexWithOpt(exp string, opt string) {
	baseCond.m[_regex] = primitive.Regex{Pattern: exp, Options: opt}
	baseCond.addMapToBuilder()
}

// Not adds `$not: exp, $options: ""` to the baseCond.m
func (baseCond *cond) Not(exp string) {
	baseCond.NotWithOpt(exp, "")
}

// NotWithOpt adds `$not: exp, $options: opt` to the baseCond.m
func (baseCond *cond) NotWithOpt(exp string, opt string) {
	baseCond.m[_not] = primitive.Regex{Pattern: exp, Options: opt}
	baseCond.addMapToBuilder()
}

// In adds `$In: vals` to the baseCond.m
func (baseCond *cond) In(vals interface{}) {
	baseCond.m[_in] = vals
	baseCond.addMapToBuilder()
}

// Nin adds `$Nin: vals` to the baseCond.m
func (baseCond *cond) Nin(vals interface{}) {
	baseCond.m[_nin] = vals
	baseCond.addMapToBuilder()
}
