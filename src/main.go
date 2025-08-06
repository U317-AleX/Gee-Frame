package main

import (
	"gee-frame/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>Hello Gee</h1>\n")
	})

	r.GET("/hello", func(c *gee.Context) {
		name := c.Query("name")
		c.String(200, "hello " + name + "\n" + "you are at /hello\n")
	})

	r.Run(":9999")
}
