package builder

import (
	"fmt"
	"time"
)

// dateCond represents a date-type condition builder.
type dateCond struct {
	*cond
	defaultFormat string
}

// newDateCond retunrs a new dateCond.
// RFC3339Nano format is used as default time format.
func newDateCond(key string, builder *Builder, format ...string) *dateCond {
	defaultFormat := time.RFC3339Nano
	if len(format) != 0 {
		defaultFormat = format[0]
	}
	return &dateCond{
		cond:          NewCond(key, builder),
		defaultFormat: defaultFormat,
	}
}

func (c *dateCond) Eq(val time.Time) *Builder {
	c.cond.eq(val)
	return c.builder
}

func (c *dateCond) EqStr(val string, format ...string) *Builder {
	t := c.mustParse(val, format...)
	return c.Eq(t)
}

func (c *dateCond) Ne(val time.Time) *Builder {
	c.cond.ne(val)
	return c.builder
}

func (c *dateCond) NeStr(val string, format ...string) *Builder {
	t := c.mustParse(val, format...)
	return c.Ne(t)
}

func (c *dateCond) Lt(val time.Time) *Builder {
	c.cond.lt(val)
	return c.builder
}

func (c *dateCond) LtStr(val string, format ...string) *Builder {
	t := c.mustParse(val, format...)
	return c.Lt(t)
}

func (c *dateCond) Lte(val time.Time) *Builder {
	c.cond.lte(val)
	return c.builder
}

func (c *dateCond) LteStr(val string, format ...string) *Builder {
	t := c.mustParse(val, format...)
	return c.Lte(t)
}

func (c *dateCond) Gt(val time.Time) *Builder {
	c.cond.gt(val)
	return c.builder
}

func (c *dateCond) GtStr(val string, format ...string) *Builder {
	t := c.mustParse(val, format...)
	return c.Gt(t)
}

func (c *dateCond) Gte(val time.Time) *Builder {
	c.cond.gte(val)
	return c.builder
}

func (c *dateCond) GteStr(val string, format ...string) *Builder {
	t := c.mustParse(val, format...)
	return c.Gte(t)
}

func (c *dateCond) Between(min, max time.Time) *Builder {
	c.cond.gte(min)
	c.cond.lte(max)
	return c.builder
}

func (c *dateCond) BetweenStr(min, max string, format ...string) *Builder {
	minT := c.mustParse(min, format...)
	maxT := c.mustParse(max, format...)
	return c.Between(minT, maxT)
}

// mustParse parses time from string with format.
func (c *dateCond) mustParse(timeStr string, format ...string) time.Time {
	f := c.defaultFormat
	if len(format) != 0 {
		f = format[0]
	}

	t, err := time.Parse(f, timeStr)
	if err != nil {
		panic(fmt.Errorf("filterBuilder: failed to parse time from string: %s with format: %s, err: %v", timeStr, f, err))
	}
	return t
}
