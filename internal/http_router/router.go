package http_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"template/internal/http_router/admin"
	"template/internal/http_router/index"
)

func Serve() (app *fiber.App) {
	viewEngine := html.New("./ui/templates", ".html")
	viewEngine.Reload(true)

	app = fiber.New(fiber.Config{Views: viewEngine})

	app.Static("/assets", "./ui/assets")

	app.Get("/", index.Homepage)

	app.Post("/admin/generateKeys", admin.GenerateKeys)
	app.Post("/admin/clearKeyHardware", admin.ClearKeyHardwareID)
	app.Post("/admin/deleteKey", admin.DeleteKey)

	app.Post("/admin/createCheat", admin.CreateCheat)
	app.Post("/admin/changeCheatStatus", admin.ChangeCheatStatus)
	app.Post("/admin/deleteCheat", admin.DeleteCheat)

	app.Post("/admin/banHardware", admin.BanHardware)
	app.Post("/admin/unbanHardware", admin.UnbanHardware)

	return
}
