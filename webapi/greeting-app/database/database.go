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

}
