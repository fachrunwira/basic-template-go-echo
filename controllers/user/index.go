package user

import (
	"log"
	"net/http"

	"github.com/fachrunwira/basic-template-go-echo/db"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID   string
	Name string
}

func Home(c echo.Context) error {

	// example execute query
	err := db.Connect()
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}
	defer db.DB.Close()

	result, err := db.DB.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatalf("failed to fetch data: %v", err)
	}
	defer result.Close()

	var users []User
	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatalf("Row scan failed: %v", err)
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Hello World!",
		"content": users,
	})
}
