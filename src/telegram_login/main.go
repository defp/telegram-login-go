package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	})

	r.GET("/auth", func(c *gin.Context) {
		username := c.Query("username")
		id := c.Query("id")
		c.String(http.StatusOK, "Hello %s %s", username, id)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
