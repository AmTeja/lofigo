package controllers

import (
	"github.com/amteja/lofig/database"
	"github.com/amteja/lofig/models"
	"github.com/gofiber/fiber/v2"
)

// GetPosts returns all posts
func GetPosts(c *fiber.Ctx) error {

	var posts []models.Post

	database.DB.Preload("User").Find(&posts)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    posts,
	})
}

func GetPost(c *fiber.Ctx) error {

	id := c.Params("id")

	var post models.Post

	database.DB.Preload("User").Find(&post, id)

	if post.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    post,
	})
}

func CreatePost(c *fiber.Ctx) error {

	var post models.Post

	if err := c.BodyParser(&post); err != nil {
		return err
	}

	database.DB.Create(&post)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    post,
	})
}

func UpdatePost(c *fiber.Ctx) error {

	id := c.Params("id")

	var post models.Post

	database.DB.First(&post, id)

	if err := c.BodyParser(&post); err != nil {
		return err
	}

	database.DB.Save(&post)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    post,
	})
}

func DeletePost(c *fiber.Ctx) error {

	id := c.Params("id")

	var post models.Post

	database.DB.Delete(&post, id)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
