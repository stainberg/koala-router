# koalart
http router

Installation

## support method
```
const (
    GET = "GET"
    POST = "POST"
    DELETE = "DELETE"
    PATCH = "PATCH"
    PUT = "PUT"
    OPTIONS = "OPTIONS"
)
```

## Install:
```
go get -u github.com/stainberg/koalart
```

## Import:
```
import "github.com/stainberg/koalart"
```

## Quickstart

### create model
```
package models

import (
    "github.com/stainberg/koalart"
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

```

### create router
```
package router

import (
    "github.com/stainberg/koalart"
    "models"
)


func init() {
    ns := koala.NewNamespace("koala",
        koala.NSController(new(models.KoalaController)),
    )
    koala.RegisterNamespace(ns)
}
```
