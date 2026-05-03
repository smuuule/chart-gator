package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/smuuule/char-gator/internal/api"
	"github.com/smuuule/char-gator/internal/db"
)

func setupDatabase() *pgxpool.Pool {
	dbURL := os.Getenv("DATABASE_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Database connection failed: %v\n", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	fmt.Println("Successfully connected to PostgreSQL database")
	return pool
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	if os.Getenv("DATABASE_URL") == "" {
		_ = os.Setenv("DATABASE_URL", "postgres://gator:password@localhost:5432/chartgator")
	}

	pool := setupDatabase()
	defer pool.Close()

	queries := db.New(pool)
	srv := api.NewServer(queries)
	router := srv.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on :%s\n", port)
	_ = router.Run(":" + port)
}
