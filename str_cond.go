package builder

import (
	"go.mongodb.org/mongo-driver/bson"
)

// strCond represents a string-type condition builder.
type strCond struct {
	*cond
}

// newStrCond constructs a new strCond.
func newStrCond(key string, builderRef *Builder) *strCond {
	return &strCond{
		cond: &cond{
			key:     key,
			m:       bson.M{},
			builder: builderRef,
		},
	}
}

// Eq adds `$eq: val` to the strc.m
func (strc *strCond) Eq(val string) *Builder {
	strc.cond.eq(val)
	return strc.builder
}

// Ne adds `$ne: val` to the strc.m
func (strc *strCond) Ne(val string) *Builder {
	strc.cond.ne(val)
	return strc.builder
}

// Regex adds `$regex: exp, $options: ""` to the strc.m
func (strc *strCond) Regex(exp string) *Builder {
	strc.cond.regex(exp)
	return strc.builder
}

// RegexWithOpt adds `$regex: exp, $options: opt` to the strc.m
func (strc *strCond) RegexWithOpt(exp string, opt string) *Builder {
	strc.cond.regexWithOpt(exp, opt)
	return strc.builder
}

// Like calls strc.Regex(val) under the wood
func (strc *strCond) Like(val string) *Builder {
	return strc.Regex(val)
}

// NotLike calls strc.Not(val) under the wood
func (strc *strCond) NotLike(val string) *Builder {
	return strc.Not(val)
}

// Not adds `$not: exp, $options: ""` to the strc.m
func (strc *strCond) Not(exp string) *Builder {
	strc.not(exp)
	return strc.builder
}

// NotWithOpt adds `$not: exp, $options: opt` to the strc.m
func (strc *strCond) NotWithOpt(exp string, opt string) *Builder {
	strc.cond.notWithOpt(exp, opt)
	return strc.builder
}

// In adds `$in: vals` to the strc.m
func (strc *strCond) In(vals ...string) *Builder {
	strc.cond.in(vals)
	return strc.builder
}

// In adds `$nin: vals` to the strc.m
func (strc *strCond) Nin(vals ...string) *Builder {
	strc.nin(vals)
	return strc.builder
}
