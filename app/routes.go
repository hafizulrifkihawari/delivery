package app

import (
	"delivery/app/gateway"
	"delivery/controllers"
)

type controllerRoutes struct {
	restaurantController *controllers.RestaurantController
}

// RegisterRoutes is used to register url routes API
func initControllers() *controllerRoutes {
	return &controllerRoutes{
		restaurantController: controllers.InitRestaurantController(nil),
	}
}
func registerRoutes() {
	var (
		controllerList = initControllers()
	)

	deliveryRouter(controllerList)
}

func deliveryRouter(c *controllerRoutes) {
	auth := router.Use(gateway.ErrorHandler)
	{
		auth.GET("health-check", controllers.HealthCheck)
		auth.GET("restaurant/:search_type", c.restaurantController.ListRestaurant)
	}
}
