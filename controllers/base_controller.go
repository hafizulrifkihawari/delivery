package controllers

//region imports
import (
	"delivery/domains/models"
	"delivery/utils"

	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
)

//endregion imports

//region structs

//endregion structs

//region funtions

func MapRequest(ctx *gin.Context, request *models.BaseRequest, keys []string) *models.BaseRequest {
	valMap := make(map[string]interface{})
	for i := range keys {
		valMap[keys[i]] = ctx.Param(keys[i])
	}
	valMap["token"] = ctx.GetHeader("Authorization")

	// get query param
	queryParams := make(map[string]string)
	for k, v := range ctx.Request.URL.Query() {
		if len(v) == 1 && len(v[0]) != 0 {
			queryParams[k] = v[0]
		}
	}

	request.QueryParam = queryParams

	mapper.MapperMap(valMap, request)
	err := ctx.ShouldBindJSON(request.BodyData)
	go utils.PrintLog((err))

	return request
}

//endregion functions
