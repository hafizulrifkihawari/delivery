package app

import (
	"delivery/app/gateway"
	"delivery/controllers"
)

type controllerRoutes struct{}

// RegisterRoutes is used to register url routes API
func initControllers() *controllerRoutes {
	return &controllerRoutes{}
}
func registerRoutes() {
	var (
		controllerList = initControllers()
	)

	deliveryRouter(controllerList)
}

func deliveryRouter(c *controllerRoutes) {
	authNoToken := router.Use(gateway.ErrorHandler)
	{
		authNoToken.GET("health_check", controllers.HealthCheck)
	}
}
