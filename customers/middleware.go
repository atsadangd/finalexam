package customers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	fmt.Println("start #AuthMiddleware")
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you don't have permission"})
		c.Abort()
		return
	}

	c.Next()

	fmt.Println("end #AuthMiddleware")

}
