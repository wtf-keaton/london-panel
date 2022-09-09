package http_router

import (
	"template/internal/http_router/admin"
	"template/internal/http_router/api"
	"template/internal/http_router/index"
	"template/internal/http_router/middleware"
	"template/pgk/session_manager"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
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
	mainGroup.Get("/keys", admin.Keys)
	mainGroup.Get("/keys/banned", admin.KeysBanned)
	mainGroup.Get("/banned_hwids", admin.BannedHardware)

	mainGroup.Post("/generateKeys", admin.GenerateKeys)
	mainGroup.Post("/addDays", admin.AddDaysAll)
	mainGroup.Get("/clearKeyHardware/:key", admin.ClearKeyHardwareID)
	mainGroup.Get("/deleteKey/:key", admin.DeleteKey)
	mainGroup.Get("/banKey/:key", admin.BanKey)
	mainGroup.Post("/createCheat", admin.CreateCheat)
	mainGroup.Post("/uploadFile/:cheat", admin.UploadFile)
	mainGroup.Get("/changeCheatStatus/:cheat", admin.ChangeCheatStatus)
	mainGroup.Get("/deleteCheat/:cheat", admin.DeleteCheat)
	mainGroup.Post("/banHardware", admin.BanHardware)
	mainGroup.Get("/unbanHardware/:hardware", admin.UnbanHardware)
	mainGroup.Post("/createUser", admin.CreateUser)
	mainGroup.Get("/deleteUser/:user", admin.DeleteUser)

	apiGroup := app.Group("/api/v1")
	apiGroup.Post("/checkKey", api.CheckKey)
	apiGroup.Post("/getDll", api.GetCheatFile)
	apiGroup.Post("/getDriver", api.GetDriverFile)
	apiGroup.Post("/activateKey", api.ActivateKey)
	apiGroup.Post("/banHardware", api.BanHardware)
	apiGroup.Post("/keyInformation", api.KeyInformation)

	return
}
