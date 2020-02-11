package ginclone

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
	router map[string]http.HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]http.HandlerFunc)}
}

func (e *Engine) addRoute(method string, pattern string, handler http.HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	e.router[key] = handler
}

func (e *Engine) GET(pattern string, handler http.HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler http.HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP(ResponseWriter, *Request)
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	handler, ok := e.router[key]
	if !ok {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
		return
	}
	handler(w, req)
}
