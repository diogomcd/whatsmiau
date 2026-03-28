package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/verbeux-ai/whatsmiau/lib/whatsmiau"
	"github.com/verbeux-ai/whatsmiau/repositories/instances"
	"github.com/verbeux-ai/whatsmiau/server/controllers"
	"github.com/verbeux-ai/whatsmiau/server/middleware"
	"github.com/verbeux-ai/whatsmiau/services"
)

func Manager(group *echo.Group) {
	repo := instances.NewRedis(services.Redis())
	w := whatsmiau.Get()
	tmpl := controllers.ParseManagerTemplates()

	controller := controllers.NewManager(repo, w, tmpl)

	group.Static("/static", "server/static")

	group.GET("", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusFound, "/manager/instances")
	})

	group.GET("/login", controller.LoginPage)
	group.POST("/login", controller.Login)

	auth := group.Group("",
		middleware.Simplify(middleware.ManagerAuth),
		middleware.Simplify(middleware.ManagerCSRF),
	)

	auth.POST("/logout", controller.Logout)
	auth.GET("/instances", controller.ListInstances)
	auth.POST("/instances", controller.CreateInstance)
	auth.GET("/instances/:id", controller.GetInstance)
	auth.GET("/instances/:id/status", controller.StatusBadge)
	auth.POST("/instances/:id/connect", controller.ConnectInstance)
	auth.GET("/instances/:id/qr-poll", controller.PollQRCode)
	auth.POST("/instances/:id/logout", controller.LogoutInstance)
	auth.DELETE("/instances/:id", controller.DeleteInstance)
	auth.PUT("/instances/:id", controller.UpdateInstance)
}
