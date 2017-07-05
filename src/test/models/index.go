package models

import (
	"koala"
	"net/http"
	"io"
)

type MainController struct {
	koala.Controller
}

func (c *MainController) URLMapping() {
	c.Mapping(koala.GET, c.Get)
}

func (c *MainController) Get()  {
	c.Ctx.Writer.WriteHeader(http.StatusOK)
	io.WriteString(c.Ctx.Writer, `Get MainController`)
}
