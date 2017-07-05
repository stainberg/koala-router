package koala

import (
	"net/http"
	"io"
	"reflect"
)

type Controller struct {
	Ctx *Context
	methodMapping  map[string]func()
}

type ControllerInterface interface {
	Init()
	Get()
	Post()
	Put()
	Patch()
	Delete()
	Options()
	URLMapping()
	Match(ctx *Context)
}

func (c *Controller)Init() {
	c.methodMapping = make(map[string]func(), 0)
}

func (c *Controller)Get()  {
	c.Ctx.Writer.WriteHeader(http.StatusNotFound)
	io.WriteString(c.Ctx.Writer, `no Controller Get match`)
}

func (c *Controller)Post()  {
	c.Ctx.Writer.WriteHeader(http.StatusNotFound)
	io.WriteString(c.Ctx.Writer, `no Controller Post match`)
}

func (c *Controller)Put()  {
	c.Ctx.Writer.WriteHeader(http.StatusNotFound)
	io.WriteString(c.Ctx.Writer, `no Controller Put match`)
}

func (c *Controller)Patch()  {
	c.Ctx.Writer.WriteHeader(http.StatusNotFound)
	io.WriteString(c.Ctx.Writer, `no Controller Patch match`)
}

func (c *Controller)Delete()  {
	c.Ctx.Writer.WriteHeader(http.StatusNotFound)
	io.WriteString(c.Ctx.Writer, `no Controller Delete match`)
}

func (c *Controller)Options()  {
	c.Ctx.Writer.WriteHeader(http.StatusNotFound)
	io.WriteString(c.Ctx.Writer, `no Controller Options match`)
}

func (c *Controller)NotSupportMethod()  {
	c.Ctx.Writer.WriteHeader(http.StatusNotFound)
	io.WriteString(c.Ctx.Writer, `no support method`)
}

func (c *Controller)URLMapping()  {

}

func (c *Controller)Mapping(method string, fn func())  {
	c.methodMapping[method] = fn
}

func (c *Controller)Match(ctx *Context) {
	c.Ctx = ctx
	c.Ctx.ParseForm()
	if _, ok := c.methodMapping[c.Ctx.Request.Method]; ok {
		fn := c.methodMapping[c.Ctx.Request.Method]
		fv := reflect.ValueOf(fn)
		fv.Call(nil)
	} else {
		if c.Ctx.Request.Method == GET {
			c.Get()
		} else if c.Ctx.Request.Method == POST {
			c.Post()
		} else if c.Ctx.Request.Method == PATCH {
			c.Patch()
		} else if c.Ctx.Request.Method == DELETE {
			c.Delete()
		} else {
			c.NotSupportMethod()
		}
	}
}