package models

import (
	"koala"
	"net/http"
	"io"
	"fmt"
	"io/ioutil"
)

type V1Controller struct {
	koala.Controller
}


func (c *V1Controller) URLMapping() {
	c.Mapping(koala.GET, c.Get)
	c.Mapping(koala.POST, c.Post)
	c.Mapping(koala.PUT, c.Put)
	c.Mapping(koala.PATCH, c.Patch)
	c.Mapping(koala.DELETE, c.Delete)
	c.Mapping(koala.OPTIONS, c.Options)
}

func (c *V1Controller) Get()  {
	c.Ctx.Writer.WriteHeader(http.StatusOK)
	io.WriteString(c.Ctx.Writer, `Get V1Controller` + fmt.Sprint(c.Ctx.Form))

}

func (c *V1Controller) Post()  {
	c.Ctx.Writer.WriteHeader(http.StatusOK)
	io.WriteString(c.Ctx.Writer, `Post V1Controller` + fmt.Sprint(c.Ctx.Form))
}

func (c *V1Controller)Put()  {
	c.Ctx.Writer.WriteHeader(http.StatusOK)
	io.WriteString(c.Ctx.Writer, `Put V1Controller` + fmt.Sprint(c.Ctx.Form))
}

func (c *V1Controller)Patch()  {
	c.Ctx.Writer.WriteHeader(http.StatusOK)
	io.WriteString(c.Ctx.Writer, `Patch V1Controller` + fmt.Sprint(c.Ctx.Form))
}

func (c *V1Controller)Delete()  {
	c.Ctx.Writer.WriteHeader(http.StatusOK)
	var reader io.Reader = c.Ctx.Request.Body
	b, e := ioutil.ReadAll(reader)
	if e != nil {

	}

	io.WriteString(c.Ctx.Writer, `Delete V1Controller Form = ` + fmt.Sprint(c.Ctx.Form) + ` & Query = ` + fmt.Sprint(c.Ctx.Query) + ` & Body = ` + string(b))
}

func (c *V1Controller)Options()  {
	c.Ctx.Writer.WriteHeader(http.StatusOK)
	io.WriteString(c.Ctx.Writer, `Options V1Controller` + fmt.Sprint(c.Ctx.Query))
}
