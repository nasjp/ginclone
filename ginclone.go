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

func (e *Engine) addRoute(method string, pattern string, f http.HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	e.router[key] = f
}

func (e *Engine) GET(pattern string, f http.HandlerFunc) {
	e.addRoute("GET", pattern, f)
}

func (e *Engine) POST(pattern string, f http.HandlerFunc) {
	e.addRoute("POST", pattern, f)
}

func (e *Engine) Run(addr string) (err error) {
	addrMsg := addr
	if len(addrMsg) > 0 && addrMsg[0] == ':' {
		addrMsg = "localhost" + addrMsg
	}
	log.Printf("=> %s", addrMsg)
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	f, ok := e.router[key]
	if !ok {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
		return
	}
	f(w, r)
}
