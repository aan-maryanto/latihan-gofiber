package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"latihan-gofiber/configs"
	"latihan-gofiber/controllers"
	"latihan-gofiber/models"
	"time"
)

func main() {

	type Person struct {
		gorm.Model
		ID        uint64 `gorm:"primaryKey;autoIncrement"`
		Name      string
		Age       int
		Dob       time.Time
		Pob       string
		IsActive  bool `gorm:"default:true"`
		IsDeleted bool `gorm:"default:false"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	type PersonRequest struct {
		Name string    `json:"name" binding:"required" example:"John Doe"`
		Age  int       `json:"age" binding:"required" example:"18"`
		Dob  time.Time `json:"dob" binding:"required" example:"2022-01-01 11:11:11"`
		Pob  string    `json:"pob" binding:"required" example:"https://www.google.com"`
	}

	db := configs.Connect()

	err := db.AutoMigrate(&models.User{}, models.Book{})
	if err != nil {
		return
	}

	app := fiber.New()

	app.Use(db)

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

	personRoute := apiRoute.Group("/person")

	// Person

	personRoute.Get("/:id?", func(ctx *fiber.Ctx) error {
		var persons []Person
		var person = new(Person)

		if ctx.Params("id") != "" {
			id := ctx.Params("id")
			result := db.First(&person, id)
			if result.Error != nil {
				return ctx.SendStatus(404)
			}
			return ctx.JSON(person)
		}

		results := db.Find(&persons)
		if results.Error != nil {
			return ctx.Status(500).JSON(fiber.Map{"error": results.Error})
		}
		return ctx.JSON(persons)
	})

	personRoute.Post("/", func(ctx *fiber.Ctx) error {
		request := new(PersonRequest)
		if err := ctx.BodyParser(request); err != nil {
			return err
		}
		person := Person{Name: request.Name, Age: request.Age, Dob: request.Dob, Pob: request.Pob, CreatedAt: time.Now()}
		db.Create(&person)
		return ctx.JSON(person)
	})

	personRoute.Put("/:id/update", func(ctx *fiber.Ctx) error {
		request := new(PersonRequest)
		person := new(Person)
		if err := ctx.BodyParser(request); err != nil {
			return err
		}
		db.Find(&person, ctx.Params("id"))
		person.Name = request.Name
		person.Age = request.Age
		person.Dob = request.Dob
		person.Pob = request.Pob
		person.UpdatedAt = time.Now()
		db.Save(&person)
		return ctx.JSON(person)
	})

	personRoute.Delete("/:id/delete", func(ctx *fiber.Ctx) error {
		request := new(PersonRequest)
		person := new(Person)
		if err := ctx.BodyParser(request); err != nil {
			return err
		}
		db.Find(&person, ctx.Params("id"))
		db.Delete(&person)
		return ctx.JSON(person)
	})

	// END PERSON

	app.Listen(":3000")
}
