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
	c.cond.Eq(val)
	return c.builder
}

// Ne adds `$ne: val` to the c.m
func (c *numCond) Ne(val interface{}) *Builder {
	c.cond.Ne(val)
	return c.builder
}

// Lt adds `$lt: val` to the c.m
func (c *numCond) Lt(val interface{}) *Builder {
	c.cond.Lt(val)
	return c.builder
}

// Lte adds `$lte: val` to the c.m
func (c *numCond) Lte(val interface{}) *Builder {
	c.cond.Lte(val)
	return c.builder
}

// Gt adds `$gt: val` to the c.m
func (c *numCond) Gt(val interface{}) *Builder {
	c.cond.gt(val)
	return c.builder
}

// Gte adds `$gte: val` to the c.m
func (c *numCond) Gte(val interface{}) *Builder {
	c.cond.Gte(val)
	return c.builder
}

// Between => [min, max]
func (c *numCond) Between(min interface{}, max interface{}) *Builder {
	c.cond.Gte(min)
	c.cond.Lte(max)
	return c.builder
}

// In adds `$nin: vals` to the c.m
func (c *numCond) In(nums interface{}) *Builder {
	c.cond.In(nums)
	return c.builder
}
