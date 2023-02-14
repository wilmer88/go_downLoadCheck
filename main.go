package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wilmer88/go_downLoadCheck/controllers"
	// "gorm-test/controllers"
	"net/http"
	 "github.com/gin-contrib/cors"
)

func main() {

	
	r := setupRouter()
	_ = r.Run(":8080")
	
}

func setupRouter() *gin.Engine {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	r.Use(cors.New(config))
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.New()
	r.POST("/lafamily", userRepo.CreateUser)
	r.GET("/lafamily.herokuapp.com/", userRepo.GetUsers)
	r.GET("/lafamily/:id", userRepo.GetUser)
	r.PUT("/lafamily/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/:id", userRepo.DeleteUser)

	return r
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}