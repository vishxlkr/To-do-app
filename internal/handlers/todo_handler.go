package handlers

import (
	"net/http"
	"todo/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateTodoInput struct {
	Title     string `json:"title" binding:"required"`
	Completed string `json:"completed"`
}

func CreateTodoHandler(pool *pgxpool) gin.HandlerFunc{

	return func (c *gin.Context){
		var input CreateTodoInput
		
		if err := c.ShouldBindJSON(&input); err!=nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()
			})
		}
		
		// if there is no error then this is continue 

		todo , err :=  repository.CreateTodo(pool )
	}

}
