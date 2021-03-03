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

// addMapToBuilder adds c.m to the referenced builder's map with key as c.key.
func (c *cond) addMapToBuilder() {
	c.builder.curMap[c.key] = c.m
}

// eq adds `$eq: val` to the c.m
func (c *cond) eq(val interface{}) {
	c.m[_eq] = val
	c.addMapToBuilder()
}

// ne adds `$ne: val` to the c.m
func (c *cond) ne(val interface{}) {
	c.m[_ne] = val
	c.addMapToBuilder()
}

// regex adds `$regex: exp, $options: ""` to the c.m
func (c *cond) regex(exp string) {
	c.regexWithOpt(exp, "")
}

// regexWithOpt adds `$regex: exp, $options: opt` to the c.m
func (c *cond) regexWithOpt(exp string, opt string) {
	c.m[_regex] = primitive.Regex{Pattern: exp, Options: opt}
	c.addMapToBuilder()
}

// not adds `$not: exp, $options: ""` to the c.m
func (c *cond) not(exp string) {
	c.notWithOpt(exp, "")
}

// notWithOpt adds `$not: exp, $options: opt` to the c.m
func (c *cond) notWithOpt(exp string, opt string) {
	c.m[_not] = primitive.Regex{Pattern: exp, Options: opt}
	c.addMapToBuilder()
}

// in adds `$in: vals` to the c.m
func (c *cond) in(vals interface{}) {
	c.m[_in] = vals
	c.addMapToBuilder()
}

// nin adds `$nin: vals` to the c.m
func (c *cond) nin(vals interface{}) {
	c.m[_nin] = vals
	c.addMapToBuilder()
}
