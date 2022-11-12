package main

import (
	"errors"
	"strconv"
	"testing"
)

type MyContext struct {
	ctx map[string]string
}

func NewMyContext() MyContext {
	return MyContext{
		ctx: make(map[string]string),
	}
}

func (c *MyContext) fromVar(key string, val interface{}) error {
	switch s := val.(type) {
	case string:
		c.ctx[key] = s
	case *string:
		c.ctx[key] = *s
	case int:
		c.ctx[key] = strconv.Itoa(s)
	case *int:
		c.ctx[key] = strconv.Itoa(*s)
	}
	return errors.New("unsupported type")
}

func (c *MyContext) fromString(key string, val string) {
	c.ctx[key] = val
}

func (c *MyContext) fromInt(key string, val int) {
	c.ctx[key] = strconv.Itoa(val)
}

func BenchmarkEmptyInterface(b *testing.B) {
	url := "/index/"
	size := 1000
	pUrl := &url
	pSize := &size

	b.Run("empty interface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ctx := NewMyContext()
			ctx.fromVar("url", url)
			ctx.fromVar("purl", pUrl)
			ctx.fromVar("size", size)
			ctx.fromVar("psize", pSize)
		}
	})
	b.Run("concrete type", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ctx := NewMyContext()
			ctx.fromString("url", url)
			ctx.fromString("purl", *pUrl)
			ctx.fromInt("size", size)
			ctx.fromInt("psize", *pSize)
		}
	})
}
