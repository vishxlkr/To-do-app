package main

import (
	"log"
	"todo/internal/config"
	"todo/internal/database"
	"todo/internal/handlers"
	"todo/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	var cfg *config.Config
	var err error
	cfg, err = config.Load() // loading env file

	if err != nil {
		log.Fatal("failed to load configuration : ", err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("failed to connect to the database : ", err)
	}

	defer pool.Close() // executed once the main function completes

	var router *gin.Engine = gin.Default()

	router.SetTrustedProxies(nil)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Todo API is running",
			"success":  true,
			"database": "connected",
		})
	})

	//public routes
	router.POST("/auth/register", handlers.CreateUserHandler(pool))
	router.POST("/auth/login", handlers.LoginHandler(pool, cfg))

	protected := router.Group("/todos")

	protected.Use(middleware.AuthMiddleware(cfg))

	protected.POST("", handlers.CreateTodoHandler(pool))
	protected.GET("", handlers.GetAllTodosHandler(pool))
	protected.GET("/:id", handlers.GetToDoByIDHandler(pool))
	protected.PUT("/:id", handlers.UpdateToDoHandler(pool)) // body will also be sent
	protected.DELETE("/:id", handlers.DeleteToDoHandler(pool))

	// middleware test route
	router.GET("/protected-test", middleware.AuthMiddleware(cfg), handlers.TestProtectedHandler())

	router.Run(":" + cfg.Port) // this start the server at given port
}
