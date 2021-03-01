package builder

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// strCond represents a string-type condition builder.
type strCond struct {
	// key is the root name for the current condtion, eq: {key: {$eq: ...}}.
	key string
	// TODO consider remove the m, populate the builder.curMap directly.
	// m is the current map stores the value of condition.
	m bson.M
	// builder refers to the current builder.
	builder *Builder
}

// newStrCond constructs a new strCond.
func newStrCond(key string, builderRef *Builder) *strCond {
	return &strCond{
		key:     key,
		m:       bson.M{},
		builder: builderRef,
	}
}

// addMapToBuilder adds strc.m to the referenced builder's map with key as strc.key.
func (strc *strCond) addMapToBuilder() {
	strc.builder.curMap[strc.key] = strc.m
}

// Eq adds `$eq: val` to the strc.m
func (strc *strCond) Eq(val string) *Builder {
	strc.m[_eq] = val
	strc.addMapToBuilder()
	return strc.builder
}

// Ne adds `$ne: val` to the strc.m
func (strc *strCond) Ne(val string) *Builder {
	strc.m[_ne] = val
	strc.addMapToBuilder()
	return strc.builder
}

// Regx adds `$regex: exp, $options: ""` to the strc.m
func (strc *strCond) Regx(exp string) *Builder {
	return strc.RegxWithOpt(exp, "")
}

// RegxWithOpt adds `$regex: exp, $options: opt` to the strc.m
func (strc *strCond) RegxWithOpt(exp string, opt string) *Builder {
	strc.m[_regex] = primitive.Regex{Pattern: exp, Options: opt}
	strc.addMapToBuilder()
	return strc.builder
}

// Like calls strc.Regex(val) under the wood
func (strc *strCond) Like(val string) *Builder {
	return strc.Regx(val)
}

// NotLike calls strc.Not(val) under the wood
func (strc *strCond) NotLike(val string) *Builder {
	return strc.Not(val)
}

// Not adds `$not: exp, $options: ""` to the strc.m
func (strc *strCond) Not(exp string) *Builder {
	return strc.NotWithOpt(exp, "")
}

// NotWithOpt adds `$not: exp, $options: opt` to the strc.m
func (strc *strCond) NotWithOpt(exp string, opt string) *Builder {
	strc.m[_not] = primitive.Regex{Pattern: exp, Options: opt}
	strc.addMapToBuilder()
	return strc.builder
}

// In adds `$in: vals` to the strc.m
func (strc *strCond) In(vals ...string) *Builder {
	strc.m[_in] = vals
	strc.addMapToBuilder()
	return strc.builder
}

// In adds `$nin: vals` to the strc.m
func (strc *strCond) Nin(vals ...string) *Builder {
	strc.m[_nin] = vals
	strc.addMapToBuilder()
	return strc.builder
}
