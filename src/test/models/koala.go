package models

import (
	"koala"
	"net/http"
	"io"
)

type KoalaController struct {
	koala.Controller
}

func (k *KoalaController) URLMapping() {
	k.Mapping(koala.GET, k.Get)
	k.Mapping(koala.POST, k.Post)
}

func (c *KoalaController) Get()  {
	c.Ctx.Writer.WriteHeader(http.StatusOK)
	io.WriteString(c.Ctx.Writer, `Get KoalaController`)
}

func (c *KoalaController) Post()  {
	c.Ctx.Writer.WriteHeader(http.StatusOK)
	io.WriteString(c.Ctx.Writer, `Post KoalaController`)
}