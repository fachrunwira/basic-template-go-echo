package main

import (
	"log"
	"os"

	"github.com/fachrunwira/basic-template-go-echo/middlewares/ratelimiting"
	"github.com/fachrunwira/basic-template-go-echo/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Failed to open env files: %v", errEnv)
		return
	}

	appPort := os.Getenv("APP_PORT")

	e := echo.New()

	// Whitelist IP
	// allowedIPs := []string{}
	// e.Use(ipwhitelisting.IPWhitelist(allowedIPs))

	// Rate Limiter
	limiter := ratelimiting.NewClientLimiter(5, 10)
	e.Use(limiter.Middleware())

	e.Use(middleware.CORS())

	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":" + appPort))
}
