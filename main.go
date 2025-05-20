package main

import (
	"log"
	"os"

	"github.com/fachrunwira/basic-template-go-echo/middlewares/ratelimiting"
	"github.com/fachrunwira/basic-template-go-echo/routes"
	"github.com/natefinch/lumberjack"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Logger using lumberjack
	logFile := &lumberjack.Logger{
		Filename:   "./log/app.log",
		MaxSize:    10,   // Max file size before rotation
		MaxBackups: 3,    // Max old log files to retain
		MaxAge:     28,   // Max days to retain a log files
		Compress:   true, // Compress old files
	}

	log.SetOutput(logFile)

	// Load .env files
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Failed to open env files: %v", errEnv)
		return
	}

	appPort := os.Getenv("APP_PORT")

	e := echo.New()

	e.Logger.SetOutput(logFile)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
		Format: `${time_rfc3339} | ${status} | ${method} | ${uri} | ${latency_human}` + "\n",
	}))

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
