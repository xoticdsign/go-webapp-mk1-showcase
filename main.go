package main

import (
	"go-webapp-mark1-showcase/gorm"
	"go-webapp-mark1-showcase/server/handlers"
	"go-webapp-mark1-showcase/server/middleware"
	"time"

	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

// APP

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(errors.New("err from main.go: godotenv.Load failed. " + err.Error()))
	}

	err = gorm.ConfigPostgreSQL()
	if err != nil {
		log.Fatal(errors.New("err from main.go: gorm.ConfigPostgreSQL failed. " + err.Error()))
	}

	app := fiber.New(fiber.Config{
		Views:        html.New("./server/templates", ".html"),
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
		AppName:      "go-webapp-mark1-showcase",
	})

	app.Use("/", middleware.Auth)

	app.Get("/", handlers.MainPage)
	app.Get("/login", handlers.Login)
	app.Post("/login/login-submit", handlers.LoginSubmit)
	app.Get("/registration", handlers.Registration)
	app.Post("/registration/registration-submit", handlers.RegistrationSubmit)
	app.Get("/logout", handlers.Logout)
	app.Get("/profile", handlers.Profile)
	app.Get("/profile/edit", handlers.ProfileEdit)
	app.Post("/profile/edit/edit-submit", handlers.EditSubmit)

	err = app.ListenTLS("0.0.0.0:6524", "./cert.crt", "./cert.key")
	if err != nil {
		log.Fatal(errors.New("err from main.go: app.ListenTLS failed. " + err.Error()))
	}
}
