package main

import "net/http"
import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run()

}
