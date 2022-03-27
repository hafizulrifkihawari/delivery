package controllers

import (
	"delivery/domains/models"
	"delivery/response"
	"delivery/services/logics"
	"delivery/utils"

	"github.com/gin-gonic/gin"
)

type RestaurantController struct {
	restaurantLogic logics.IRestaurantService
}

func InitRestaurantController(restaurantLogic logics.IRestaurantService) *RestaurantController {
	if utils.IsNil(restaurantLogic) {
		restaurantLogic = logics.InitRestaurantService(nil)
	}
	controller := RestaurantController{restaurantLogic: restaurantLogic}

	return &controller
}

func (h *RestaurantController) ListRestaurant(c *gin.Context) {
	var (
		request = MapRequest(c, &models.BaseRequest{}, []string{})
	)
	res, err := h.restaurantLogic.ListRestaurantByFilter(request.QueryParam)
	if err != nil {
		response := response.Response{}
		response.Error(c, err.Error())
	} else {
		response := response.Response{Data: res}
		response.Success(c)
	}
}
