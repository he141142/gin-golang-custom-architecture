package restaurants

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	restaurant2 "sykros-pro/gopro/src/service/restaurant"
	"sykros-pro/gopro/src/share/logger"
)

type handleRestaurant interface {
	createRestaurant() func(context *gin.Context)
	getRestaurantById() func(context *gin.Context)
}

type RestaurantRouter struct {
	handleRestaurant
	name    string
	router  *gin.RouterGroup
	service *restaurant2.RestaurantService
}

func (restaurantRouter *RestaurantRouter) Setup(r *gin.Engine, name string, db *gorm.DB) {
	restaurantRouter.name = name
	var vicePerLogger logger.LoggerService
	vicePerLogger = logger.LogrusSetup("RESTAURANT_SERVICE")
	serviceInitializingMsg := vicePerLogger.GetContext()+" Initializing"
	vicePerLogger.LogWithMsg(serviceInitializingMsg,logger.INFO)
	restaurantRouter.service = &restaurant2.RestaurantService{
		Db:     db,
		Logger: vicePerLogger,
	}
	restaurant := r.Group(name)
	{
		restaurant.POST("", restaurantRouter.createRestaurant())
		restaurant.GET("", restaurantRouter.getAllRestaurants())
		restaurant.GET("/:id", restaurantRouter.getRestaurantById())

	}
}

func SetUpRestaurantRouters(router *gin.Engine, db *gorm.DB) {
	restaurantRouter := &RestaurantRouter{
		name: "/restaurants",
	}
	restaurantRouter.Setup(router, restaurantRouter.name, db)
}
