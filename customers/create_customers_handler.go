package customers

import (
	"net/http"

	"github.com/atsadangd/finalexam/database"
	"github.com/atsadangd/finalexam/types"
	"github.com/gin-gonic/gin"
)

func CreateCustomersHandler(c *gin.Context) {
	cus := types.Customer{}

	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// id, err := database.CreateCustomer(cus.Name, cus.Email, cus.Status)
	id, err := database.CreateCustomer(cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	cus.ID = id
	c.JSON(http.StatusCreated, cus)
}
