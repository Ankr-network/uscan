package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(static.Serve("/", static.LocalFile("dist", true)))

	app.StaticFS("/dist", http.Dir("./dist"))

	app.StaticFile("/favicon.ico", "./favicon.ico")

	app.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write((content))
			c.Writer.Flush()
		}
	})
	// run static server
	app.Run("0.0.0.0:8888")
}
