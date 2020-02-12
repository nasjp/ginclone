package ginclone

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Path:    req.URL.Path,
		Method:  req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}

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
	e.addRoute("GET", pattern, f)
}

func (e *Engine) POST(pattern string, f HandlerFunc) {
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
	c := newContext(w, r)
	e.router.handle(c)
}
