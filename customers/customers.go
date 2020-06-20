package customers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(AuthMiddleware)

	r.POST("/customers", CreateCustomersHandler)
	r.GET("/customers/:id", GetCustomersByIdHandler)
	r.GET("/customers", GetCustomersHandler)
	r.PUT("/customers/:id", UpdateCustomersHandler)
	r.DELETE("/customers/:id", DeleteCustomersHandler)

	return r
}
