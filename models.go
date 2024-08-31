package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"github.com/gofiber/fiber/v2"
)

type Menu struct {
	gorm.Model
	Name     string `json:"name"`
	Price    int `json:"price"`
	Category string `json:"category"`
}

func CreateMenu(db *gorm.DB, c *fiber.Ctx) error {
	menu := new(Menu)

	err := c.BodyParser(menu)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result := db.Create(menu)
	if result.Error != nil {
	  return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"Message": "CreateMenu successful",
	})
}

func GetMenu(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")

	var menu Menu
	
	result := db.First(&menu , id)
	if result.Error != nil {
		log.Fatalf("Error Get Menu : %v", result.Error)
	}
	fmt.Println("Get Menu successfully")

	return c.JSON(menu)
}

func GetMenus(db *gorm.DB , c *fiber.Ctx) error {
	var menus []Menu
	result := db.Find(&menus)

	if result.Error != nil {
		log.Fatalf("Error Get Menu : %v", result.Error)
	}

	return c.JSON(menus)
}

// Update
func UpdateMenu(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	
	menu := new(Menu)
	
	db.First(&menu, id)
	if err := c.BodyParser(menu); err != nil {
	  return c.SendStatus(fiber.StatusBadRequest)
	}
	db.Save(&menu)
	return c.JSON(menu)
}


// soft Del is Didn't actually delete it. but is still in the database.
func DeleteMenu(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")

	var menu Menu

	result := db.Delete(&menu, id)
	if result.Error != nil {
	  return c.SendStatus(fiber.StatusBadRequest)
	}
	
	return c.JSON(fiber.Map{
		"message ": "Del successful",
	})
}

// Hard Del
func Delete(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")

	var menu Menu

	result := db.Unscoped().Delete(&menu, id)
	if result.Error != nil {
	  return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message ": "Del successful",
	})
}

// Get by Where 
// Return slice because easy and safety
// func GetMenuByName(db *gorm.DB , NameMenu string) []Menu {
// 	var menu []Menu
// 	result := db.Where("Name = ?", NameMenu).Find(&menu)
// 	if result.Error != nil {
// 		return nil
// 	}

// 	return menu
// }

// func getMenuSorted(db *gorm.DB) ([]Menu, error) {
// 	var menu []Menu
// 	result := db.Order("created_at desc").Find(&menu)
// 	if result.Error != nil {
// 	  return nil , result.Error
// 	}
// 	return menu , nil
// }