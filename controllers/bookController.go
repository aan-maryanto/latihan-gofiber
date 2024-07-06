package controllers

import (
	"github.com/gofiber/fiber/v2"
	"latihan-gofiber/configs"
	"latihan-gofiber/dtos/requests"
	"latihan-gofiber/models"
	"time"
)

func GetBook(ctx *fiber.Ctx) error {
	db := configs.Connect()
	var book = new(models.Book)
	var books []models.Book
	if ctx.Params("id") != "" {
		id := ctx.Params("id")
		result := db.First(&book, id)
		if result.Error != nil {
			return ctx.SendStatus(404)
		}
		return ctx.JSON(book)
	}
	results := db.Find(&books)
	if results.Error != nil {
		return ctx.SendStatus(404)
	}
	return ctx.JSON(books)

}

func PostBook(ctx *fiber.Ctx) error {
	db := configs.Connect()
	request := new(requests.BookRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}
	book := models.Book{
		Title:       request.Title,
		Author:      request.Author,
		Publisher:   request.Publisher,
		Category:    request.Category,
		Year:        request.Year,
		ISBN:        request.ISBN,
		Description: request.Description,
		Image:       request.Image,
		CreatedAt:   time.Now(),
		CreatedBy:   "System",
	}
	db.Create(&book)
	return ctx.JSON(book)
}

func PutBook(ctx *fiber.Ctx) error {
	book := new(models.Book)
	db := configs.Connect()
	request := new(requests.BookRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}
	db.First(&book, ctx.Params("id"))
	book.Title = request.Title
	book.Author = request.Author
	book.Publisher = request.Publisher
	book.Category = request.Category
	book.Year = request.Year
	book.ISBN = request.ISBN
	book.Description = request.Description
	book.Image = request.Image
	book.UpdatedAt = time.Now()
	book.UpdatedBy = "System"

	db.Save(&book)
	return ctx.JSON(book)
}

func DeleteBook(ctx *fiber.Ctx) error {
	db := configs.Connect()
	var book = new(models.Book)
	err := db.First(&book, ctx.Params("id"))
	if err != nil {
		return ctx.SendStatus(404)
	}
	db.Delete(&book)
	return ctx.SendStatus(200)
}
