package customers

import (
	"net/http"

	"github.com/atsadangd/finalexam/database"
	"github.com/atsadangd/finalexam/types"
	"github.com/gin-gonic/gin"
)

func GetCustomersByIdHandler(c *gin.Context) {
	id := c.Param("id")

	row, err := database.GetCustomersByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	cus := &types.Customer{}
	err = row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cus)
}
