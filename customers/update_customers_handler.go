package customers

import (
	"net/http"

	"github.com/atsadangd/finalexam/database"
	"github.com/gin-gonic/gin"
)

func UpdateCustomersHandler(c *gin.Context) {
	id := c.Param("id")

	row, err := database.GetCustomersByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	cus := &Customer{}
	err = row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := c.ShouldBindJSON(cus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.UpdateCustomersById(id, cus.Name, cus.Email, cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, cus)
}
