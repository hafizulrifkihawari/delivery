package app

import (
	"delivery/app/gateway"
	"delivery/controllers"
)

const (
	ParentRoute string = "delivery"
)

type controllerRoutes struct {
	restaurantController *controllers.RestaurantController
	userController       *controllers.UserController
}

// RegisterRoutes is used to register url routes API
func initControllers() *controllerRoutes {
	return &controllerRoutes{
		restaurantController: controllers.InitRestaurantController(nil),
		userController:       controllers.InitUserController(nil),
	}
}
func registerRoutes() {
	var (
		controllerList = initControllers()
	)

	deliveryRouter(controllerList)
}

func deliveryRouter(c *controllerRoutes) {
	auth := router.Group(ParentRoute).Use(gateway.ErrorHandler)
	{
		auth.GET("health-check", controllers.HealthCheck)
		auth.GET("restaurant/:search_type", c.restaurantController.ListRestaurant)
		auth.POST("purchase", c.userController.Purchase)
	}
}
