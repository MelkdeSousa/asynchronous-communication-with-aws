package main

import (
	"net/http"

	"github.com/melkdesousa/wottva/users/api/handlers"
	"github.com/melkdesousa/wottva/users/internal/database"
	"github.com/melkdesousa/wottva/users/internal/messaging"
)

func main() {

	conn := database.DbConnection()

	defer conn.Close()

	// create a new user repository
	userRepo := database.NewPgUserRepo(conn)

	// create a new broker
	broker := messaging.NewSNSBroker()

	// create a new user controller
	userController := handlers.NewUserController(userRepo, broker)

	app := http.NewServeMux()
	app.Handle("/users", userController)
	app.Handle("/users/", userController)
	http.ListenAndServe(":8080", app) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
