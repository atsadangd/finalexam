package customers

import (
	"net/http"

	"github.com/atsadangd/finalexam/database"
	"github.com/gin-gonic/gin"
)

func createCustomersHandler(c *gin.Context) {
	cus := Customer{}
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := database.Conn().QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3)  RETURNING id", cus.Name, cus.Email, cus.Status)

	err := row.Scan(&cus.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, cus)
}

/*
func getTodosHandler(c *gin.Context) {
	status := c.Query("status")

	stmt, err := database.Conn().Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	todos := []Todo{}
	for rows.Next() {
		t := Todo{}

		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		todos = append(todos, t)
	}

	tt := []Todo{}

	for _, item := range todos {
		if status != "" {
			if item.Status == status {
				tt = append(tt, item)
			}
		} else {
			tt = append(tt, item)
		}
	}

	c.JSON(http.StatusOK, tt)
}

func getTodoByIdHandler(c *gin.Context) {
	id := c.Param("id")

	stmt, err := database.Conn().Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	row := stmt.QueryRow(id)

	t := &Todo{}

	err = row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, t)
}

func updateTodosHandler(c *gin.Context) {
	id := c.Param("id")
	stmt, err := database.Conn().Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	row := stmt.QueryRow(id)

	t := &Todo{}

	err = row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := c.ShouldBindJSON(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err = database.Conn().Prepare("UPDATE todos SET status=$2, title=$3 WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := stmt.Exec(id, t.Status, t.Title); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, t)
}
*/
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(AuthMiddleware)

	r.POST("/customers", createCustomersHandler)
	// r.GET("/customers/:id", getCustomersByIdHandler)
	// r.GET("/customers", getCustomersHandler)
	// r.PUT("/customers/:id", updateCustomersHandler)
	// r.DELETE("/customers/:id", DeleteCustomersHandler)

	return r
}
