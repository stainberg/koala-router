package koala

import (
	"net/http"
	"io"
	"time"
	"strconv"
	"strings"
)

type RouterHandler struct {
	Root *Namespace
}

func (handler *RouterHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	time := time.Now().UnixNano()
	w.Header().Set("sequence", Base64Encode(strconv.FormatInt(time, 10)))
	ctx := new(Context)
	ctx.Init()
	ctx.Request = r
	ctx.Writer = w
	find := handler.match(ctx)
	if !find {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `no match`)
	}
}

func (handler *RouterHandler)match(c *Context) bool {
	r := c.Request
	if strings.EqualFold(r.URL.Path, "/") {
		return _match([]string{"/"}, handler.Root, c)
	} else {
		pre := strings.Split(r.URL.Path, "/")
		pre[0] = "/"
		return _match(pre, handler.Root, c)
	}
}

func _match(s []string, n *Namespace, c *Context) bool {
	if len(s) == 0 {
		if n.handler != nil {
			n.handler.Match(c)
			return true
		}
	} else if len(s) == 1 {
		if s[0] == n.prefix && n.handler != nil {
			n.handler.Match(c)
			return true
		}
		if n.prefix[:1] == ":" && n.handler != nil {
			c.Vars.Add(n.prefix[1:], s[0])
			n.handler.Match(c)
			return true
		}
	} else {
		if s[0] == n.prefix {
			s = s[1:]
			for _, ns := range n.linkTree {
				if _match(s, ns, c) {
					return true
				}
			}
		} else {
			if n.prefix[:1] == ":" {
				ss := s[1:]
				c.Vars.Add(n.prefix[1:], s[0])
				for _, ns := range n.linkTree {
					if _match(ss, ns, c) {
						return true
					}
				}
			}
		}
	}
	return false
}

