package main

import (
	"net/http"

	"github.com/nasjp/ginclone"
)

func main() {
	r := ginclone.New()

	r.GET("/", func(c *ginclone.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gin</h1>")
	})

	r.GET("/hello", func(c *ginclone.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *ginclone.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *ginclone.Context) {
		c.JSON(http.StatusOK, map[string]string{"filepath": c.Param("filepath")})
	})

	r.Run(":6969")
}
