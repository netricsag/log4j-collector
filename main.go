package main

import (
	"os"
	"time"

	"github.com/bluestoneag/log4j-collector/database"
	"github.com/bluestoneag/log4j-collector/models"
	"github.com/bluestoneag/log4j-collector/routes"
	"github.com/bluestoneag/log4j-collector/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbName string
	err    error
)

func init() {
	util.InitLoggers()
	dbName = os.Getenv("DB_NAME")
	if dbName == "" {
		util.WarningLogger.Println("DB_NAME is not set, setting it to default value: log4j-collector.db")
		dbName = "log4j-collector.db"
	}

	_, err = os.Stat(dbName)
	if os.IsNotExist(err) {
		util.WarningLogger.Println("Database file does not exist, creating it")
		file, err := os.Create(dbName)
		if err != nil {
			util.ErrorLogger.Println(err)
			os.Exit(1)
		}
		defer file.Close()
	} else if err != nil {
		currentTime := time.Now().Local()
		err = os.Chtimes(dbName, currentTime, currentTime)
		if err != nil {
			util.ErrorLogger.Println(err)
			os.Exit(1)
		}
	}

	database.DBConn, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		util.ErrorLogger.Println(err)
		os.Exit(1)
	}

	database.DBConn.AutoMigrate(&models.Report{})
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("github.com/bluestoneag/log4j-collector")
	})

	routes.Setup(app)
	util.InfoLogger.Printf("Started web server on port 8080")

	app.Listen(":8080")

	// defer database.DBConn.Close()
}
