package main

import (
	"log"
	"todo/internal/config"
	"todo/internal/database"

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
			"message": "Todo API is running",
			"success": true,
		})
	})

	router.Run(":" + cfg.Port) // this start the server at given port
}
