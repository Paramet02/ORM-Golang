package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"
  )

// config database
const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		SlowThreshold: time.Second, // Slow SQL threshold
		LogLevel:      logger.Info, // Log level
		Colorful:      true,        // Enable color
		},
	)
	
	// connect database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // add Logger
	})

	// panic : is a builtin function that we can use to stop the flow if a critical situation arises.
	if err != nil {
		panic("fail to connect database")
	}

	println(db)
	fmt.Println("Connect Database successful")

	// create database 
	// func autoMigrate can not del database but only create database
	// db.AutoMigrate(&Menu{})


	// setup fiber
	app := fiber.New()
	
	app.Get("/menus", func(c *fiber.Ctx) error {
		return c.JSON(GetMenus(db))
	})
	
	app.Get("/menu/:id" , func(c *fiber.Ctx) error {
		id := c.Params("id")
		Menu := GetMenu(db , id)
	})

	app.Listen(":8000")

	
}
