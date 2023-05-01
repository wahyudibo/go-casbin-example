package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/wahyudibo/go-casbin-example/internal/config"
	"github.com/wahyudibo/go-casbin-example/internal/database/mysql"
	"github.com/wahyudibo/go-casbin-example/internal/router"
)

func main() {
	// initializes config
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to initialize config: %+v\n", err)
	}

	db, err := mysql.New(cfg)
	if err != nil {
		log.Fatalf("failed to database client: %+v\n", err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatalf("failed to run database migration: %+v\n", err)
	}

	router := router.New(db.GetConnection())

	router.Run()
}
