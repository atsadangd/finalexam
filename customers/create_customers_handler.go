package customers

import (
	"net/http"

	"github.com/atsadangd/finalexam/database"
	"github.com/gin-gonic/gin"
)

func CreateCustomersHandler(c *gin.Context) {
	cus := Customer{}

	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, id := database.CreateCustomer(cus.Name, cus.Email, cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	cus.ID = id
	c.JSON(http.StatusCreated, cus)
}
