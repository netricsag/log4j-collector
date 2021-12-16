package controllers

import (
	"github.com/bluestoneag/log4j-collector/database"
	"github.com/bluestoneag/log4j-collector/models"
	"github.com/bluestoneag/log4j-collector/util"
	"github.com/gofiber/fiber/v2"
)

func GetReports(c *fiber.Ctx) error {
	db := database.DBConn
	// Get all reports with their vulnerable files
	var reports []models.Report
	db.Preload("VulnerableFiles").Find(&reports)
	return c.Status(200).JSON(reports)
}

func GetReport(c *fiber.Ctx) error {
	db := database.DBConn
	var report models.Report
	// Get report with its vulnerable files where ID is equal to the ID in the URL
	if err := db.Preload("VulnerableFiles").First(&report, c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(map[string]interface{}{"message": "Report not found"})
	}
	return c.Status(200).JSON(report)
}

func CreateReport(c *fiber.Ctx) error {
	db := database.DBConn
	var report models.Report
	if err := c.BodyParser(&report); err != nil {
		return c.Status(500).JSON(map[string]interface{}{"message": "Error parsing request body"})
	}
	// Check if server name is empty
	if report.ServerName == "" {
		return c.Status(500).JSON(map[string]interface{}{"message": "Server name cannot be empty"})
	}
	// Check if report already exists in database
	var existingReport models.Report
	if err := db.Where("server_name = ?", report.ServerName).First(&existingReport); err == nil {
		util.WarningLogger.Printf("Report with server name %s does not exists in database", report.ServerName)
	}
	if existingReport.ID != 0 {
		// Report already exists in database, delete all existing vulnerable files
		db.Model(&existingReport).Association("VulnerableFiles").Clear()

		// Check if vulnerable files are empty
		if len(report.VulnerableFiles) == 0 {
			// No vulnerable files, delete report
			db.Delete(&existingReport)
			return c.Status(200).JSON(map[string]interface{}{"message": "Report deleted"})
		}

		// Add new vulnerable files to existing report
		db.Model(&existingReport).Association("VulnerableFiles").Append(report.VulnerableFiles)
		return c.Status(200).JSON(existingReport)
	}
	// Report does not exist in database, create it
	db.Create(&report)
	return c.Status(200).JSON(report)
}

func UpdateReport(c *fiber.Ctx) error {
	db := database.DBConn
	// Get report with its vulnerable files where ID is equal to the ID in the URL
	var report models.Report
	if err := c.BodyParser(&report); err != nil {
		return c.Status(500).JSON(map[string]interface{}{"message": "Error parsing request body"})
	}
	// Get existing report
	var existingReport models.Report
	if err := db.Preload("VulnerableFiles").First(&existingReport, c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(map[string]interface{}{"message": "Report not found"})
	}
	if existingReport.ID != 0 {
		// Delete all existing vulnerable files
		db.Model(&existingReport).Association("VulnerableFiles").Clear()

		// Check if vulnerable files are empty
		if len(report.VulnerableFiles) == 0 {
			// No vulnerable files, delete report
			db.Delete(&existingReport)
			return c.Status(200).JSON(map[string]interface{}{"message": "Report deleted"})
		}

		// Add new vulnerable files to existing report
		db.Model(&existingReport).Association("VulnerableFiles").Append(report.VulnerableFiles)
		// Update existing with new values
		db.Model(&existingReport).Updates(report)
		return c.Status(200).JSON(existingReport)
	} else {
		return c.Status(500).JSON(map[string]interface{}{"message": "Report not found"})
	}
}

func DeleteReport(c *fiber.Ctx) error {
	db := database.DBConn
	// Get report with its vulnerable files where ID is equal to the ID in the URL
	var report models.Report
	if err := db.Preload("VulnerableFiles").First(&report, c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(map[string]interface{}{"message": "Report not found"})
	}
	// Delete all existing vulnerable files
	db.Model(&report).Association("VulnerableFiles").Clear()
	// Delete report
	db.Delete(&report)
	return c.Status(200).JSON(map[string]interface{}{"message": "Report deleted"})
}
