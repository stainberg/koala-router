# koalart < http router >


## support method
```go
const (
    GET = "GET"
    POST = "POST"
    DELETE = "DELETE"
    PATCH = "PATCH"
    PUT = "PUT"
    OPTIONS = "OPTIONS"
)
```

## Installation

### Install:
```
go get -u github.com/stainberg/koalart
```

### Import:
```go
import "github.com/stainberg/koalart"
```

### Quickstart

#### create model
```go
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

#### create router
```go
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

#### use
```go
package main

import (
    _ "router"//init router
    "github.com/stainberg/koalart"
)


func main() {
    koala.Run("8888")
}
```