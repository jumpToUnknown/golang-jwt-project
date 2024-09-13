package main

import(
	routes "golang-jwt-token-project/routes"
	"os"
	"github.com/gin-gonic/gin"
)

func main(){
  port = os.Getenv("PORT")

	if port==""{
		port="8000"
	}

	router := gin.New()
	router.User(gin.logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.M{"success":"Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.M{"success":"Access granted for api-2"})
	})

	router.RUN(":" + port)
}