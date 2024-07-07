package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"latihan-gofiber/configs"
	"latihan-gofiber/controllers"
	"latihan-gofiber/models"
	"log"
	"net/http"
)

func main() {

	db := configs.Connect()

	err := db.AutoMigrate(&models.User{}, models.Book{})
	if err != nil {
		return
	}

	app := fiber.New()

	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./assets"),
	}))

	//simple route
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	//app.Use(jwtware.New(jwtware.Config{
	//	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	//}))

	authRoute := app.Group("/auth")

	//Start Auth
	// LOGIN
	authRoute.Post("/login", controllers.Login)

	//REGISTER
	authRoute.Post("/register", controllers.Register)

	//END AUTH

	apiRoute := app.Group("/api")

	apiRoute.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	bookRoute := apiRoute.Group("/book")

	bookRoute.Get("/:id?", controllers.GetBook)
	bookRoute.Post("/", controllers.PostBook)
	bookRoute.Put("/:id/update", controllers.PutBook)
	bookRoute.Delete("/:id/delete", controllers.DeleteBook)
	bookRoute.Post("/:id/upload", controllers.UploadBookImage)

	log.Fatal(app.Listen(":3000"))
}
