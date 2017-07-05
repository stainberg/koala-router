package test

import (
	"koala"
	"test/models"
)

func init() {
	ns := koala.NewNamespace("v1",
		koala.NSNamespace("koala",
			koala.NSController(new(models.KoalaController)),
		),
		koala.NSController(new(models.V1Controller)),
	)
	koala.RegisterNamespace(ns)
	koala.Index(new(models.MainController))
}