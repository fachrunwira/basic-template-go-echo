package config

import "os"

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_CONNECTION"),
	}
}
