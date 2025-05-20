package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fachrunwira/basic-template-go-echo/config"
)

var DB *sql.DB

func Connect() error {
	db_cfg := config.LoadDBConfig()

	var err error

	DB, err = sql.Open(connectionType(db_cfg.Driver), dsn(&db_cfg))

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	return nil
}

func connectionType(driver string) string {
	switch driver {
	case "pgsql":
		return "postgres"
	case "mysql":
		return "mysql"
	}

	return "unknown"
}

func dsn(cfg *config.DBConfig) string {
	switch cfg.Driver {
	case "pgsql":
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	}

	return ""
}
