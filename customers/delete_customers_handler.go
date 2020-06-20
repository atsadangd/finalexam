package customers

import (
	"net/http"

	"github.com/atsadangd/finalexam/database"
	"github.com/gin-gonic/gin"
)

func DeleteCustomersHandler(c *gin.Context) {
	id := c.Param("id")

	err := database.DeleteCustomersById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, "customer deleted.")
}
