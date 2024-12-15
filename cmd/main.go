package main

import (
	"TODOLIST/internal/db"
	"TODOLIST/internal/handlers"
	"TODOLIST/internal/middleware"
	"TODOLIST/internal/repository"
	"TODOLIST/internal/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {

	dbConn, err := db.ConnectDB()

	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	defer dbConn.Close()

	createTable(dbConn)

	bookrepo := repository.NewBookRepositoryDB(dbConn)
	bookHandler := handlers.NewBookHandler(bookrepo)

	router := gin.Default()

	router.Use(middleware.Logger())

	routes.RegisterRoutes(router, bookHandler)

	router.Run(":8081")
}

func createTable(db *sqlx.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			author TEXT NOT NULL UNIQUE
		);
	`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
	fmt.Println("Table created successfully (if not already exists).")
}
