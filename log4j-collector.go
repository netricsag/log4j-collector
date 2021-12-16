package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/bluestoneag/log4j-collector/database"
	"github.com/bluestoneag/log4j-collector/models"
	"github.com/bluestoneag/log4j-collector/routes"
	"github.com/bluestoneag/log4j-collector/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbName     string
	dbPath     string
	dbLocation string
	err        error
)

func init() {
	util.InitLoggers()
	dbName = os.Getenv("DB_NAME")
	if dbName == "" {
		util.WarningLogger.Println("DB_NAME is not set, setting it to default value: log4j-collector.db")
		dbName = "log4j-collector.db"
	}

	dbPath = os.Getenv("DB_PATH")
	if dbPath == "" {
		util.WarningLogger.Println("DB_PATH is not set, setting it to default value: ./db")
		dbPath = "./db"
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		util.WarningLogger.Println("DB_PATH does not exist, creating it")
		err = os.MkdirAll(dbPath, 0755)
		if err != nil {
			util.ErrorLogger.Println("Failed to create DB_PATH")
			os.Exit(1)
		}
	}

	dbLocation = filepath.Join(dbPath, filepath.Base(dbName))
	util.InfoLogger.Printf("DB_LOCATION: %s", dbLocation)

	_, err = os.Stat(dbLocation)
	if os.IsNotExist(err) {
		util.WarningLogger.Println("Database file does not exist, creating it")
		file, err := os.Create(dbLocation)
		if err != nil {
			util.ErrorLogger.Println(err)
			os.Exit(1)
		}
		defer file.Close()
	} else if err != nil {
		currentTime := time.Now().Local()
		err = os.Chtimes(dbLocation, currentTime, currentTime)
		if err != nil {
			util.ErrorLogger.Println(err)
			os.Exit(1)
		}
	}

	database.DBConn, err = gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
	if err != nil {
		util.ErrorLogger.Println(err)
		os.Exit(1)
	}

	database.DBConn.AutoMigrate(&models.Report{})
	database.DBConn.AutoMigrate(&models.VulnerableFile{})
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("github.com/bluestoneag/log4j-collector")
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.Setup(app)
	util.InfoLogger.Printf("Started web server on port 8080")

	app.Listen(":8080")

	// defer database.DBConn.Close()
}
