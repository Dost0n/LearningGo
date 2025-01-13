package main

import (
	"fmt"
	"log"

	"github.com/Dost0n/simple-medium-close/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)

	if err != nil {
		log.Fatalf("Failed to cinnect to postgres: %v", err)
	}
	log.Println("Postgresql Connection successfully done!")

	_ = psqlConn
}
