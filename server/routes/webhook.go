package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/verbeux-ai/whatsmiau/lib/whatsmiau"
	"github.com/verbeux-ai/whatsmiau/repositories/instances"
	"github.com/verbeux-ai/whatsmiau/server/controllers"
	"github.com/verbeux-ai/whatsmiau/services"
)

func Webhook(group *echo.Group) {
	redisInstance := instances.NewRedis(services.Redis())
	controller := controllers.NewWebhooks(redisInstance, whatsmiau.Get())

	group.POST("/set/:instance", controller.Set)
	group.GET("/find/:instance", controller.Find)
}
