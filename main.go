package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name/doc", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*.html", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param(".html")})
	})

	v1 := r.Group("v1")
	{
		v1.GET("/hello", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello v1</h1>")
		})

		v1.GET("/assets/*.doc", func(ctx *gee.Context) {
			ctx.JSON(http.StatusOK, gee.H{"file": ctx.Param(".doc")})
		})
	}
	v2 := r.Group("v2")
	{
		v2.GET("/hello", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello v2</h1>")
		})

		v2.GET("/assets/*.doc", func(ctx *gee.Context) {
			ctx.JSON(http.StatusOK, gee.H{"file": ctx.Param(".doc")})
		})
	}

	r.Run(":9999")
}
