package routes

import (
	"github.com/amteja/lofig/controllers"
	"github.com/amteja/lofig/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {

	app.Post("/auth/register", controllers.RegisterWithEmail)
	app.Post("/auth/login", controllers.LoginWithEmail)
	app.Post("/auth/logout", controllers.Logout)
	app.Get("/auth/user", controllers.User)

}

func PostRoutes(a *fiber.App) {
	route := a.Group("/posts")

	route.Post("/", middlewares.JWTProtected(), controllers.CreatePost)
	route.Get("/", middlewares.JWTProtected(), controllers.GetPosts)
	route.Get("/:id", middlewares.JWTProtected(), controllers.GetPost)
	route.Put("/:id", middlewares.JWTProtected(), controllers.UpdatePost)
	route.Delete("/:id", middlewares.JWTProtected(), controllers.DeletePost)
}

func UserRoutes(a *fiber.App) {
	route := a.Group("/users")

	route.Get("/:id", controllers.GetUser)
	route.Post("/f/:id", middlewares.JWTProtected(), controllers.AddUserFollower)
}
