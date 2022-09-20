package restaurants

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	restaurant2 "sykros-pro/gopro/src/service/restaurant"
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
	restaurantRouter.service = &restaurant2.RestaurantService{
		Db: db,
	}
	restaurant := r.Group(name)
	{
		restaurant.POST("",restaurantRouter.createRestaurant())
		restaurant.GET("/:id", func(context *gin.Context) {

		})
	}
}

func SetUpRestaurantRouters(router *gin.Engine, db *gorm.DB) {
	restaurantRouter := &RestaurantRouter{
		name: "/restaurants",
	}
	restaurantRouter.Setup(router, restaurantRouter.name, db)
}


