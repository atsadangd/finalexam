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

// func deleteTodoById(id string) error {
// 	err := database.DeleteCustomersById(id)
// 	if err != nil {
// 		return &errors.Error{
// 			Code:    666,
// 			Message: "business not allow to delete specialitem",
// 		}
// 	}

// 	return nil
// }
