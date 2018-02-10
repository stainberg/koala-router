package koala

import (
	"net/http"
	"net/url"
	"errors"
	"mime"
	"io"
	"io/ioutil"
	"fmt"
)

type maxBytesReader struct {
	w   http.ResponseWriter
	r   io.ReadCloser // underlying reader
	n   int64         // max bytes remaining
	err error         // sticky error
}

func (l *maxBytesReader) Close() error {
	return l.r.Close()
}

func (l *maxBytesReader) Read(p []byte) (n int, err error) {
	if l.err != nil {
		return 0, l.err
	}
	if len(p) == 0 {
		return 0, nil
	}
	if int64(len(p)) > l.n+1 {
		p = p[:l.n+1]
	}
	n, err = l.r.Read(p)

	if int64(n) <= l.n {
		l.n -= int64(n)
		l.err = err
		return n, err
	}

	n = int(l.n)
	l.n = 0
	type requestTooLarger interface {
		requestTooLarge()
	}
	if res, ok := l.w.(requestTooLarger); ok {
		res.requestTooLarge()
	}
	l.err = errors.New("http: request body too large")
	return n, l.err
}

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Form    url.Values
	Query   url.Values
	Vars    url.Values
	MIME    string
	Body    string
}

func (ctx *Context) Init() {
	ctx.Form = make(url.Values)
	ctx.Query = make(url.Values)
	ctx.Vars = make(url.Values)
}

func (ctx *Context) ParseForm() error {
	var queryValues url.Values
	defer func() {
		copyValues(ctx.Form, ctx.Request.PostForm)
		copyValues(ctx.Query, queryValues)
	}()

	var err error
	if ctx.Request.PostForm == nil {
		if ctx.Request.Method == POST || ctx.Request.Method == PUT || ctx.Request.Method == PATCH || ctx.Request.Method == DELETE || ctx.Request.Method == OPTIONS {
			ctx.Request.PostForm, err = ctx.parsePostForm(ctx.Request)
		}
		if ctx.Request.PostForm == nil {
			ctx.Request.PostForm = make(url.Values)
		}
	}
	if ctx.Request.Form == nil {
		if len(ctx.Request.PostForm) > 0 {
			ctx.Request.Form = make(url.Values)
			copyValues(ctx.Request.Form, ctx.Request.PostForm)
		}
		if ctx.Request.URL != nil {
			queryValues = ctx.Request.URL.Query()
		}
		if queryValues == nil {
			queryValues = make(url.Values)
		}
		if ctx.Request.Form == nil {
			ctx.Request.Form = queryValues
		} else {
			copyValues(ctx.Request.Form, queryValues)
		}
	}
	return err
}

func (ctx *Context) parsePostForm(r *http.Request) (vs url.Values, err error) {
	if r.Body == nil {
		err = errors.New("missing form body")
		return
	}
	ct := r.Header.Get("Content-Type")
	if ct == "" {
		ct = "application/octet-stream"
	}
	ct, _, err = mime.ParseMediaType(ct)
	ctx.MIME = ct
	switch {
	case ct == "application/x-www-form-urlencoded":
		var reader io.Reader = r.Body
		maxFormSize := int64(1<<63 - 1)
		if _, ok := r.Body.(*maxBytesReader); !ok {
			maxFormSize = int64(10 << 20) // 10 MB is a lot of text.
			reader = io.LimitReader(r.Body, maxFormSize+1)
		}
		b, e := ioutil.ReadAll(reader)
		if e != nil {
			if err == nil {
				err = e
			}
			break
		}
		if int64(len(b)) > maxFormSize {
			err = errors.New("http: POST too large")
			return
		}
		vs, e = url.ParseQuery(string(b))
		if err == nil {
			err = e
		}
	case ct == "multipart/form-data":

	case ct == "application/json" || ct == "text/plain" || ct == "application/xml" || ct == "text/xml" || ct == "text/html" || ct == "application/javascript":
		var reader io.Reader = r.Body
		maxFormSize := int64(1<<63 - 1)
		if _, ok := r.Body.(*maxBytesReader); !ok {
			maxFormSize = int64(10 << 20) // 10 MB is a lot of text.
			reader = io.LimitReader(r.Body, maxFormSize+1)
		}
		b, e := ioutil.ReadAll(reader)
		if e != nil {
			if err == nil {
				err = e
			}
			break
		}
		if int64(len(b)) > maxFormSize {
			err = errors.New("http: POST too large")
			return
		}
		fmt.Printf("len = %d, length = %d", len(b), r.ContentLength)
		ctx.Body = string(b)
	}
	return
}


func copyValues(dst, src url.Values) {
	for k, vs := range src {
		for _, value := range vs {
			dst.Add(k, value)
		}
	}
}