package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/verbeux-ai/whatsmiau/server/middleware"
)

func Load(app *echo.Echo) {
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	app.Pre(middleware.Simplify(middleware.Auth))

	V1(app.Group("/v1"))
}

func V1(group *echo.Group) {
	Root(group)
	Instance(group.Group("/instance"))
	Message(group.Group("/instance/:instance/message"))
	Chat(group.Group("/instance/:instance/chat"))

	ChatEVO(group.Group("/chat"))
	MessageEVO(group.Group("/message"))
	Webhook(group.Group("/webhook"))
}
