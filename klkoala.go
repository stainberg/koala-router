package koala

import (
	"net/http"
	"log"
)

const (
	GET     = "GET"
	POST    = "POST"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	PUT     = "PUT"
	OPTIONS = "OPTIONS"
)

type LinkedNamespace func(*Namespace)

type Namespace struct {
	handler ControllerInterface
	prefix string
	linkTree []*Namespace

}

var r *RouterHandler

func init() {
	r = new(RouterHandler)
}

func Run(port string) {
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}

func NewNamespace(prefix string, params ...LinkedNamespace) *Namespace {
	ns := new(Namespace)
	ns.linkTree = make([]*Namespace, 0)
	ns.prefix = prefix
	for _, p := range params {
		p(ns)
	}
	return ns
}

func NSNamespace(prefix string, params ...LinkedNamespace) LinkedNamespace {
	return func(ns *Namespace) {
		n := NewNamespace(prefix, params...)
		ns.Namespace(n)
	}
}

func NSController(controller ControllerInterface) LinkedNamespace {
	return func(ns *Namespace) {
		ns.Controller(controller)
	}
}

func RegisterNamespace(n *Namespace)  {
	if r.Root == nil {
		ns := new(Namespace)
		ns.linkTree = make([]*Namespace, 0)
		ns.linkTree = append(ns.linkTree, n)
		ns.prefix = "/"
		r.Root = ns
	} else {
		r.Root.linkTree = append(r.Root.linkTree, n)
	}
}

func Index(controller ControllerInterface) {
	if r.Root == nil {
		ns := new(Namespace)
		ns.linkTree = make([]*Namespace, 0)
		ns.prefix = "/"
		r.Root = ns
	}
	r.Root.Controller(controller)
}

func (n *Namespace) Namespace(ns *Namespace) *Namespace {
	n.linkTree = append(n.linkTree, ns)
	return n
}

func (n *Namespace) Controller (controller ControllerInterface) {
	controller.Init()
	controller.URLMapping()
	n.handler = controller
}

func printNamespace(n *Namespace) {
	println(n.prefix, n.handler)
	for _, v := range n.linkTree {
		printNamespace(v)
	}
}