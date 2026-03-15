package handlers

import (
	"net/http"
	"strconv"

	"todo/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateTodoInput struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

type UpdateToDoInput struct {
	Title string `json:"title" `
	// we are using the pointer bool to get the already present value in the todo
	// &true --------> set completed as -> true
	// &false -------> set completed as -> false
	// nil ----------> set completed as -> not provided -> use default value false
	Completed *bool `json:"completed"`
}

func CreateTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {
		var input CreateTodoInput

		// parse json body and converts it into go struct and save it in input variable
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// if there is no error then this is save it to databse (repository is used)
		todo, err := repository.CreateTodo(pool, input.Title, input.Completed)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// if everything goes well, then we return 201 status and the todo that is saved by the client
		c.JSON(http.StatusCreated, todo)
	}

}

func GetAllTodosHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {
		todos, err := repository.GetAllTodos(pool)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, todos)

	}
}

func GetToDoByIDHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {
		idStr := c.Param("id")
		// "2" ------> 2, nil
		// "a" ------> 0 , error (invalid syntax)

		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid todo ID",
			})
			return
		}

		todo, err := repository.GetToDoByID(pool, id)

		if err != nil {
			if err == pgx.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Todo not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, todo)

	}
}

// here before creating this handler -> create struct for updated todo
func UpdateToDoHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {

	}
}
