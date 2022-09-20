package restaurants

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sykros-pro/gopro/src/constants/response"
	"sykros-pro/gopro/src/database/model"
	"sykros-pro/gopro/src/github.com/fatih/structs"
	"sykros-pro/gopro/src/share/errorHandling"
)

func (restaurantRouter *RestaurantRouter) createRestaurant() func(context *gin.Context){
	return func(context *gin.Context) {
		var data model.Restaurant

		if err := errorHandling.ValidateDto(context, &data); err != nil {
			response.SystemResponse(http.StatusBadRequest, map[string]interface {
			}{
				"error": err.Error(),
			}, context)
		}

		err, dataresponse := restaurantRouter.service.CreateRestaurant(&data, restaurantRouter.service.Db)
		if err != nil {
			response.SystemResponse(http.StatusBadRequest, map[string]interface {
			}{
				"error": err.Error(),
			}, context)
		}
		response.SystemResponse(http.StatusOK, structs.Map(dataresponse), context)
	}
}


func (restaurantRouter *RestaurantRouter) getRestaurantById() func(context *gin.Context){
	return func(context *gin.Context) {
		var data model.Restaurant

		if err := errorHandling.ValidateDto(context, &data); err != nil {
			response.SystemResponse(http.StatusBadRequest, map[string]interface {
			}{
				"error": err.Error(),
			}, context)
		}

		err, dataresponse := restaurantRouter.service.CreateRestaurant(&data, restaurantRouter.service.Db)
		if err != nil {
			response.SystemResponse(http.StatusBadRequest, map[string]interface {
			}{
				"error": err.Error(),
			}, context)
		}
		response.SystemResponse(http.StatusOK, structs.Map(dataresponse), context)
	}
}
