package builder

// numCond represents a numeric-type condition builder.
// For convince, the val passed in MUST to be a numeric type,
// APIs WILL NOT check its type.
type numCond struct {
	*cond
}

func newNumCond(key string, builder *Builder) *numCond {
	return &numCond{
		cond: newCond(key, builder),
	}
}

// Eq adds `$eq: val` to the c.m
func (c *numCond) Eq(val interface{}) *Builder {
	c.cond.eq(val)
	return c.builder
}

// Ne adds `$ne: val` to the c.m
func (c *numCond) Ne(val interface{}) *Builder {
	c.cond.ne(val)
	return c.builder
}

// Lt adds `$lt: val` to the c.m
func (c *numCond) Lt(val interface{}) *Builder {
	c.cond.lt(val)
	return c.builder
}

// Lte adds `$lte: val` to the c.m
func (c *numCond) Lte(val interface{}) *Builder {
	c.cond.lte(val)
	return c.builder
}

// Gt adds `$gt: val` to the c.m
func (c *numCond) Gt(val interface{}) *Builder {
	c.cond.gt(val)
	return c.builder
}

// Gte adds `$gte: val` to the c.m
func (c *numCond) Gte(val interface{}) *Builder {
	c.cond.gte(val)
	return c.builder
}

// Between => [min, max]
func (c *numCond) Between(min interface{}, max interface{}) *Builder {
	c.cond.gte(min)
	c.cond.lte(max)
	return c.builder
}

// In adds `$nin: vals` to the c.m
func (c *numCond) In(nums interface{}) *Builder {
	c.cond.in(nums)
	return c.builder
}
