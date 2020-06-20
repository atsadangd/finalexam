package customers

import (
	"net/http"

	"github.com/atsadangd/finalexam/database"
	"github.com/gin-gonic/gin"
)

func GetCustomersByIdHandler(c *gin.Context) {
	id := c.Param("id")

	err, row := database.GetCustomersByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	cus := &Customer{}
	err = row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cus)
}
