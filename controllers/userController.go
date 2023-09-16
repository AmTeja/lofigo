package controllers

import (
	"log"

	"github.com/amteja/lofig/database"
	"github.com/amteja/lofig/models"
	"github.com/amteja/lofig/util"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {

	var user models.User
	var id = c.Params("id")

	database.DB.Find(&user, id)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    user,
	})
}

func AddUserFollower(c *fiber.Ctx) error {

	var ownerUser models.User
	var targetUser models.User

	var targetId = c.Params("id")

	// Get current user id from JWT
	jwt := c.Cookies("jwt")

	ownerUser, error := util.UserFromJWT(jwt)

	if error != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	database.DB.Preload("Followers").Preload("Following").Find(&targetUser, targetId)

	if targetUser.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Check if user is already following
	var isFollowing = false

	for _, follower := range targetUser.Followers {
		log.Println(follower.Id)
		log.Println(ownerUser.Id)

		if follower.FollowerId == ownerUser.Id {
			isFollowing = true
			break
		}
	}

	if isFollowing {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User is already following",
		})
	}

	// Add follower
	var follower = models.Follower{
		UserId:     targetUser.Id,
		FollowerId: ownerUser.Id,
	}

	database.DB.Create(&follower)

	// Add following
	var following = models.Following{
		UserId:      ownerUser.Id,
		FollowingId: targetUser.Id,
	}

	database.DB.Create(&following)

	database.DB.Preload("Followers").Preload("Following").Find(&targetUser, targetId)

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    targetUser,
	})
}
