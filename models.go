package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Menu struct {
	gorm.Model
	Name     string `json:"name"`
	Price    int `json:"price"`
	Category string `json:"category"`
}

func CreateMenu(db *gorm.DB, menu *Menu) error{
	result := db.Create(menu)
	if result.Error != nil {
	  return result.Error
	}
	return nil
}

func GetMenu(db *gorm.DB, id int) *Menu {
	var menu Menu
	result := db.First(&menu , id)
	if result.Error != nil {
		log.Fatalf("Error Get Menu : %v", result.Error)
	}
	fmt.Println("Get Menu successfully")

	return &menu
}

func GetMenus(db *gorm.DB) []Menu {
	var menus []Menu
	result := db.Find(&menus)
	if result.Error != nil {
		log.Fatalf("Error Get Menu : %v", result.Error)
	}
	fmt.Println("Get Menu successfully")

	return menus
}

// Update
func UpdateMenu(db *gorm.DB, menu *Menu) error{
	result := db.Model(&menu).Updates(&menu)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


// soft Del is Didn't actually delete it. but is still in the database.
func DeleteMenu(db *gorm.DB, id int) error {
	var menu Menu
	result := db.Delete(&menu, id)
	if result.Error != nil {
	  return result.Error
	}
	
	return nil
}

// Hard Del
func Delete(db *gorm.DB, id int) error {
	var menu Menu
	result := db.Unscoped().Delete(&menu, id)
	if result.Error != nil {
	  return result.Error
	}

	return nil
}

// Get by Where 
// Return slice because easy and safety
func GetMenuByName(db *gorm.DB , NameMenu string) []Menu {
	var menu []Menu
	result := db.Where("Name = ?", NameMenu).Find(&menu)
	if result.Error != nil {
		return nil
	}

	return menu
}

func getMenuSorted(db *gorm.DB) ([]Menu, error) {
	var menu []Menu
	result := db.Order("created_at desc").Find(&menu)
	if result.Error != nil {
	  return nil , result.Error
	}
	return menu , nil
}