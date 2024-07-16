package controllers

import (
	"github.com/gofiber/fiber/v2"
	"latihan-gofiber/configs"
	"latihan-gofiber/models"
	"time"
)

func GetMember(ctx *fiber.Ctx) error {
	db := configs.Connect()
	var member = new(models.Member)
	var members []models.Member
	if ctx.Params("id") != "" {
		id := ctx.Params("id")
		result := db.First(&members, "id=?", id)
		if result.Error != nil {
			return ctx.SendStatus(404)
		}
		return ctx.JSON(member)
	}
	result := db.Find(&members)
	if result.Error != nil {
		return ctx.SendStatus(404)
	}
	return ctx.JSON(members)
}

func PostMember(ctx *fiber.Ctx) error {
	db := configs.Connect()
	request := new(models.Member)
	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}
	member := models.Member{
		Name:       request.Name,
		IdentityNo: request.IdentityNo,
		Address:    request.Address,
		Phone:      request.Phone,
		IsActive:   true,
		IsDeleted:  false,
		CreatedAt:  time.Now(),
		CreatedBy:  "System",
	}
	db.Create(&member)
	return ctx.JSON(member)
}

func PutMember(ctx *fiber.Ctx) error {
	member := new(models.Member)
	db := configs.Connect()
	request := new(models.Member)
	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}
	db.Find(&member, "id=?", request.ID)
	member.Name = request.Name
	member.IdentityNo = request.IdentityNo
	member.Address = request.Address
	member.Phone = request.Phone
	member.UpdatedAt = time.Now()
	member.UpdatedBy = "System"

	db.Save(&member)
	return ctx.JSON(member)
}

func DeleteMember(ctx *fiber.Ctx) error {
	db := configs.Connect()
	var member = new(models.Member)
	err := db.First(&member, "id=?", ctx.Params("id"))
	if err != nil {
		return ctx.SendStatus(404)
	}
	db.Delete(&member)
	return ctx.SendStatus(200)
}
