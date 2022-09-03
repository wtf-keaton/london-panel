package http_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"template/internal/http_router/admin"
	"template/internal/http_router/index"
	"template/internal/http_router/middleware"
	"template/pgk/session_manager"
)

func Serve() (app *fiber.App) {
	viewEngine := html.New("./ui/templates", ".html")
	viewEngine.Reload(true)

	app = fiber.New(fiber.Config{Views: viewEngine})

	app.Static("/assets", "./ui/assets")

	app.Use(func(c *fiber.Ctx) error {
		c.Bind(fiber.Map{
			"Authorized": session_manager.IsAuthorized(c),
			"User":       session_manager.GetUser(c),
		})
		return c.Next()
	})

	app.Get("/", index.AuthPage)
	app.Post("/auth/login", admin.LoginIn)

	mainGroup := app.Group("/admin", middleware.AuthCheck)
	mainGroup.Get("/", admin.Homepage)
	mainGroup.Get("/users", admin.Users)
	mainGroup.Get("/cheats", admin.Cheats)

	mainGroup.Post("/generateKeys", admin.GenerateKeys)
	mainGroup.Post("/clearKeyHardware", admin.ClearKeyHardwareID)
	mainGroup.Post("/deleteKey", admin.DeleteKey)
	mainGroup.Post("/createCheat", admin.CreateCheat)

	mainGroup.Get("/changeCheatStatus/:cheat", admin.ChangeCheatStatus)
	mainGroup.Get("/deleteCheat/:cheat", admin.DeleteCheat)

	mainGroup.Post("/banHardware", admin.BanHardware)
	mainGroup.Post("/unbanHardware", admin.UnbanHardware)
	mainGroup.Post("/createUser", admin.CreateUser)
	mainGroup.Post("/deleteUser", admin.DeleteUser)

	return
}
