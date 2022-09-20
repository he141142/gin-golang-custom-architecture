package restaurant

import (
	"sykros-pro/gopro/src/database/model"
	"gorm.io/gorm"
)

type restaurantService interface {
	createRestaurant(restaurant *model.Restaurant) (error,model.Restaurant)
}

type RestaurantService struct {
	restaurantService
	Db *gorm.DB
}

