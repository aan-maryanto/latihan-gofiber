package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"latihan-gofiber/configs"
	"latihan-gofiber/dtos/requests"
	"latihan-gofiber/models"
	"latihan-gofiber/utils"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Login(ctx *fiber.Ctx) error {
	db := configs.Connect()
	request := new(requests.LoginRequest)
	user := new(models.User)
	if err := ctx.BodyParser(request); err != nil {
		return err
	}
	db.Where("email = ?", request.Email).First(user)
	if user == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password+user.Salt))
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	claims := jwt.MapClaims{
		"name":  user.Name,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.JSON(fiber.Map{"token": token})
}

func Register(ctx *fiber.Ctx) error {
	db := configs.Connect()
	request := new(requests.RegisterRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}
	salt := utils.RandomString(4, charset)
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password+salt), 12)
	if err != nil {
		return err
	}
	user := models.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  string(password),
		Salt:      salt,
		IsActive:  true,
		IsDeleted: false,
		CreatedAt: time.Now(),
	}

	db.Create(&user)
	return ctx.JSON(user)
}
