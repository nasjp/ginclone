package ginclone

import (
	"fmt"
	"log"
	"strings"
)

type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*node
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	ss := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, s := range ss {
		if s != "" {
			parts = append(parts, s)
			if s[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parsePattern(pattern), 0)
	r.handlers[method+"-"+pattern] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	searchParts := parsePattern(path)
	n := root.search(searchParts, 0)
	if n == nil {
		return nil, nil
	}

	params := make(map[string]string)
	for i, part := range parsePattern(n.pattern) {
		switch part[0] {
		case ':':
			params[part[1:]] = searchParts[i]
		case '*':
			if len(part) <= 1 {
				continue
			}
			params[part[1:]] = strings.Join(searchParts[i:], "/")
			break
		}
	}

	return n, params
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n == nil {
		fmt.Fprintf(c.Writer, "404 NOT FOUND: %s\n", c.Path)
		return
	}

	log.Printf("%s %4s", c.Method, c.Path)
	c.Params = params
	r.handlers[c.Method+"-"+n.pattern](c)
}
