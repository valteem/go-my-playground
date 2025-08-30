package database

import (
	"context"
	"fmt"
	"greeting-app/config"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(cfg *config.Config) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error
	DB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	if err = DB.Ping(context.Background()); err != nil {
		log.Fatal("Unable to ping database:", err)
	}

	log.Println("Successfully connected to database")

	createTables()
}

func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(30) DEFAULT 'user' CHECK (role IN ('admin', 'greetings_manager', 'user')),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS greetings (
			id SERIAL PRIMARY KEY,
			message TEXT NOT NULL,
			created_by INTEGER REFERENCES users(id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		// Insert default greetings if table is empty
		`INSERT INTO greetings (message, created_by) 
		 SELECT 'Looks like a nice day', NULL 
		 WHERE NOT EXISTS (SELECT 1 FROM greetings WHERE message = 'Looks like a nice day')`,

		`INSERT INTO greetings (message, created_by) 
		 SELECT 'Pretty nice morning, uh?', NULL 
		 WHERE NOT EXISTS (SELECT 1 FROM greetings WHERE message = 'Pretty nice morning, uh?')`,

		`INSERT INTO greetings (message, created_by) 
		 SELECT 'Have a good day', NULL 
		 WHERE NOT EXISTS (SELECT 1 FROM greetings WHERE message = 'Have a good day')`,

		`INSERT INTO greetings (message, created_by) 
		 SELECT 'Check back later if you like it', NULL 
		 WHERE NOT EXISTS (SELECT 1 FROM greetings WHERE message = 'Check back later if you like it')`,
	}

	for _, query := range queries {
		_, err := DB.Exec(context.Background(), query)
		if err != nil {
			log.Fatal("Error creating tables:", err)
		}
	}

	log.Println("Database tables created successfully")
}
