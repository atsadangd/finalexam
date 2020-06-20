package customers

import (
	"net/http"

	"github.com/atsadangd/finalexam/database"
	"github.com/gin-gonic/gin"
)

func GetCustomersHandler(c *gin.Context) {

	err, rows := database.GetCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	custs := []Customer{}
	for rows.Next() {
		cus := Customer{}

		err := rows.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		custs = append(custs, cus)
	}

	c.JSON(http.StatusOK, custs)
}