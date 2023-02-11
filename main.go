package main

import (
	"net/http"
	"github.com/wilmer88/go_downLoadCheck/controllers"
   "github.com/gin-gonic/contrib/static"
   "github.com/gin-gonic/gin"
)
func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

server := gin.Default()

server.Run(":8080")
   r := gin.Default()
   r.GET("/hello", func(c *gin.Context) {
       c.String(200, "Hello")
   })
   api := r.Group("/api")
   api.GET("/ping", func(c *gin.Context) {
       c.JSON(200, gin.H{
           "message": "ping",
       })
   })
   r.Use(static.Serve("/", static.LocalFile("./views", true)))
   r.Run()
}