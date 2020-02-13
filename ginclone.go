package ginclone

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, pattern string, f HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	e.router.addRoute(method, pattern, f)
}

func (e *Engine) GET(pattern string, f HandlerFunc) {
	e.addRoute(http.MethodGet, pattern, f)
}

func (e *Engine) POST(pattern string, f HandlerFunc) {
	e.addRoute(http.MethodPost, pattern, f)
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
	c := newContext(w, r)
	e.router.handle(c)
}
