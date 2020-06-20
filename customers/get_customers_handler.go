package customers

import (
	"net/http"

	"github.com/atsadangd/finalexam/database"
	"github.com/atsadangd/finalexam/types"
	"github.com/gin-gonic/gin"
)

func GetCustomersHandler(c *gin.Context) {

	rows, err := database.GetCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	custs := []types.Customer{}
	for rows.Next() {
		cus := types.Customer{}

		err := rows.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		custs = append(custs, cus)
	}

	c.JSON(http.StatusOK, custs)
}
