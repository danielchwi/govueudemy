package routes

import (
	"govue/controllers"
	"govue/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Put("/api/user/info", controllers.UpdateInfo)
	app.Put("/api/user/password", controllers.UpdatePassword)

	app.Get("/api/users", controllers.GetUsers)
	app.Post("/api/user", controllers.CreateUser)
	app.Get("/api/user/:id", controllers.GetUser)
	app.Put("/api/user/:id", controllers.UpdateUser)
	app.Delete("/api/user/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.GetRoles)
	app.Post("/api/role", controllers.CreateRole)
	app.Get("/api/role/:id", controllers.GetRole)
	app.Put("/api/role/:id", controllers.UpdateRole)
	app.Delete("/api/role/:id", controllers.DeleteRole)

	app.Get("/api/products", controllers.GetProducts)
	app.Post("/api/product", controllers.CreateProduct)
	app.Get("/api/product/:id", controllers.GetProduct)
	app.Put("/api/product/:id", controllers.UpdateProduct)
	app.Delete("/api/product/:id", controllers.DeleteProduct)

	app.Post("/api/upload/", controllers.Upload)
	app.Static("api/uploads/", "./uploads")

	app.Get("/api/orders/export", controllers.Export)
	app.Get("/api/orders", controllers.GetOrders)
	app.Get("/api/orders/chart", controllers.Chart)
}
