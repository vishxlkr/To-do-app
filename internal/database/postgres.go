package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseUrl string) (*pgxpool.Pool, error) {

	var ctx context.Context = context.Background()

	var config *pgxpool.Config
	var err error
	config, err = pgxpool.ParseConfig(databaseUrl)

	if err != nil {
		log.Printf("Unable to parse the DATABSE URL:%v", err)
		return nil, err
	}

	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Printf("Unable to create connection pool : %v", err)
	}

	err = pool.Ping(ctx)

	if err != nil {
		log.Printf("Unable to ping Database: %v", err)
		pool.Close()
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL database.")

	return pool, nil

}
