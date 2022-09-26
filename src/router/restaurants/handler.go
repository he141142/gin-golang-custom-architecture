package restaurants

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
	"sykros-pro/gopro/src/common"
	"sykros-pro/gopro/src/constants/response"
	"sykros-pro/gopro/src/database/model"
	"sykros-pro/gopro/src/dto"
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
		println(paginateHelper.Page, paginateHelper.Limit, paginateHelper.Offset)
		restaurantsData, _ := restaurantRouter.service.GetAllRestaurant(restaurantRouter.service.Db, paginateHelper)
		dataresponse := response.BaseResponseData[dto.AllRestaurantDto](&dto.AllRestaurantDto{
			restaurantsData,
		})
		response.SystemResponse(http.StatusOK, dataresponse, context)
	}
}

func (restaurantRouter *RestaurantRouter) getRestaurantById() func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		idCvt, err := strconv.Atoi(id)
		if err != nil {
			appErr := common.InvalidTypeOfParam("id", reflect.Int, err)
			response.SystemResponse(appErr.StatusCode, appErr, context)
			return
		}
		fmt.Println(idCvt)
		restaurantsData, err := restaurantRouter.service.GetRestaurantById(restaurantRouter.service.Db, idCvt)
		if err != nil {
			appErr := common.ErrInternalServerError(err)
			response.SystemResponse(appErr.StatusCode, appErr, context)
			return
		}
		dataresponse := response.BaseResponseData[dto.RestaurantDto](
			restaurantsData,
		)
		response.SystemResponse(http.StatusOK, dataresponse, context)
	}
}

func (restaurantRouter *RestaurantRouter) updateRestaurantById() func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		idCvt, err := strconv.Atoi(id)
		if err != nil {
			appErr := common.InvalidTypeOfParam("id", reflect.Int, err)
			response.SystemResponse(appErr.StatusCode, appErr, context)
			return
		}
		var updateData dto.UpdateRestaurantDto
		err = context.ShouldBind(&updateData)
		fmt.Println("entry")

		if err != nil {
			appErr := common.ErrInternalServerError(err)
			response.SystemResponse(appErr.StatusCode, appErr, context)
			return
		}
		//fmt.Println(updateData)
		restaurantsData, err := restaurantRouter.service.
			UpdateRestaurantById(restaurantRouter.
				service.Db, idCvt, &updateData)

		if err != nil {
			appErr := common.ErrInternalServerError(err)
			response.SystemResponse(appErr.StatusCode, appErr, context)
			return
		}
		dataresponse := response.BaseResponseData[dto.RestaurantDto](
			restaurantsData,
		)
		response.SystemResponse(http.StatusOK, dataresponse, context)
	}
}
