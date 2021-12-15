package controllers

import (
	"github.com/bluestoneag/log4j-collector/database"
	"github.com/bluestoneag/log4j-collector/models"
	"github.com/gofiber/fiber/v2"
)

func GetReports(c *fiber.Ctx) error {
	db := database.DBConn
	var reports []models.Report
	db.Find(&reports)
	return c.JSON(reports)
}

func PostReport(c *fiber.Ctx) error {
	db := database.DBConn
	var report models.Report
	if err := c.BodyParser(&report); err != nil {
		return err
	}
	db.Create(&report)
	return c.JSON(report)
}
