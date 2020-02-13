package ginclone

import (
	"fmt"
	"log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	f, ok := r.handlers[key]
	if !ok {
		fmt.Fprintf(c.Writer, "404 NOT FOUND: %s\n", c.Path)
		return
	}
	log.Printf("%s %4s", c.Method, c.Path)
	f(c)
}
