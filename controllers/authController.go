package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/amteja/lofig/database"
	"github.com/amteja/lofig/env"
	"github.com/amteja/lofig/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterWithEmail(c *fiber.Ctx) error {
	var data models.RegisterDTO

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data.Email).First(&user)

	if user.Id != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Email already in use",
		})
	}

	if data.Email == "" || data.Password == "" || data.Name == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "All fields are required",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	user = models.User{
		Name:  data.Name,
		Email: data.Email,
	}

	database.DB.Create(&user)

	key := models.Key{
		Id:             "email:" + data.Email,
		HashedPassword: string(password),
		UserID:         user.Id,
		IsPrimary:      true,
	}

	database.DB.Create(&key)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func LoginWithEmail(c *fiber.Ctx) error {

	var data models.LoginDTO

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Fetch the user

	var user models.User

	database.DB.Preload("Keys").Preload("Sessions").Where("email = ?", data.Email).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Check for the Key in user model
	var key models.Key

	//itterate over the keys
	for _, k := range user.Keys {
		if k.Id == "email:"+data.Email {
			key = k
		}
	}

	log.Println(user.Sessions)

	// Compare the password
	passwordMatch := bcrypt.CompareHashAndPassword([]byte(key.HashedPassword), []byte(data.Password))

	if passwordMatch != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(env.Get("JWT_SECRET"))[0:32])

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	database.DB.Create(&models.Session{
		UserID:  user.Id,
		Expires: time.Now().Add(time.Hour * 24).Unix(),
	})

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    user,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	tokenClaims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	database.DB.Where("id = ?", tokenClaims.Issuer).First(&user)

	return c.JSON(user)

}
