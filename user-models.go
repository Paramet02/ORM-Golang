package main

import (
	"os"
	"time"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

// createUser handles user registration
func CreateUser(db *gorm.DB, c *fiber.Ctx) error {
	// Reserve memory space
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create user
	result := db.Create(user)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
func LoginUser(db *gorm.DB, c *fiber.Ctx) error {
	var input User
	var user User
  
	if err := c.BodyParser(&input); err != nil {
	  return err // Handle body parsing error
	}
  
	// Find user by email
	// SELECT * FROM users WHERE name = input.Email ORDER BY id LIMIT 1;
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
	  return c.SendStatus(fiber.StatusInternalServerError) // Handle database error
	}
  
	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
	  return c.SendStatus(fiber.StatusUnauthorized) // Handle invalid password
	}
  
	// Create JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
  
	t, err := token.SignedString([]byte(os.Getenv("jwtSecretKey")))
	if err != nil {
	  return c.SendStatus(fiber.StatusInternalServerError) // Handle JWT signing error
	}

	// Set cookie
	c.Cookie(&fiber.Cookie{
	  Name:     "JWT", // cookie name
	  Value:    t, // value name
	  Expires:  time.Now().Add(time.Hour * 72), // time
	  HTTPOnly: true, // http
	})
  
	return c.JSON(fiber.Map{"message": "success"})
  }
  
  func AuthRequired(c *fiber.Ctx) error  {
	// get value from cookie
	cookie := c.Cookies("JWT")
	key := []byte(os.Getenv("jwtSecretKey"))

	// check token from cookie
	token , err := jwt.ParseWithClaims(cookie , &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	
	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claim := token.Claims.(*jwt.MapClaims)

	fmt.Println(claim)


	// func authRequired ---> GET/menu
	return c.Next()
  }