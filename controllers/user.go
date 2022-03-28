package controllers

import (
	"delivery/domains/models"
	"delivery/response"
	"delivery/services/logics"
	"delivery/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userLogic logics.IUserService
	validate  *validator.Validate
}

func InitUserController(userLogic logics.IUserService) *UserController {
	if utils.IsNil(userLogic) {
		userLogic = logics.InitUserService(nil, nil, nil, false)
	}

	controller := UserController{
		userLogic: userLogic,
	}

	return &controller
}

func (h *UserController) Purchase(c *gin.Context) {
	var (
		request = MapRequest(c, &models.BaseRequest{BodyData: &models.UserPurchaseRequest{}}, []string{"search_type"})
	)

	// validate payload
	h.validate = validator.New()
	err := h.validate.Struct(request.BodyData)
	if err != nil {
		response := response.Response{}
		response.Error(c, err.Error())
		return
	}

	bodyData := request.BodyData.(*models.UserPurchaseRequest)

	userLogic := logics.InitUserService(nil, nil, nil, true)
	res, err := userLogic.PurchaseDish(bodyData)
	if err != nil {
		response := response.Response{}
		response.Error(c, err.Error())
	} else {
		response := response.Response{Data: res}
		response.Success(c)
	}
}
