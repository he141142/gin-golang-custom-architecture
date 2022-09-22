package restaurants

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"sykros-pro/gopro/src/constants/response"
	"sykros-pro/gopro/src/database/model"
	"sykros-pro/gopro/src/service/restaurant"
	"sykros-pro/gopro/src/share/errorHandling"
	"sykros-pro/gopro/src/utils"
)

func (restaurantRouter *RestaurantRouter) createRestaurant() func(context *gin.Context) {
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

func (restaurantRouter *RestaurantRouter) getAllRestaurants() func(context *gin.Context) {
	return func(context *gin.Context) {
		paginateHelper := &utils.PaginateHelper{}
		//prepare for paginate parameter
		paginateHelper.Processing(context, utils.LIMIT_OFFSET)
		restaurantsData, _ := restaurantRouter.service.GetAllRestaurant(restaurantRouter.service.Db, paginateHelper)
		dataresponse := response.BaseResponseData[restaurant.AllRestaurantDto](&restaurant.AllRestaurantDto{
			restaurantsData,
		})
		response.SystemResponse(http.StatusOK, dataresponse, context)
	}
}

//func (restaurantRouter *RestaurantRouter) getRestaurantById() func(context *gin.Context) {
//	return func(context *gin.Context) {
//		var data model.Restaurant
//
//		if err := errorHandling.ValidateDto(context, &data); err != nil {
//			response.SystemResponse(http.StatusBadRequest, map[string]interface {
//			}{
//				"error": err.Error(),
//			}, context)
//		}
//
//		err, dataresponse := restaurantRouter.service.CreateRestaurant(&data, restaurantRouter.service.Db)
//		if err != nil {
//			response.SystemResponse(http.StatusBadRequest, map[string]interface {
//			}{
//				"error": err.Error(),
//			}, context)
//		}
//		response.SystemResponse(http.StatusOK, structs.Map(dataresponse), context)
//	}
//}
