package restaurant

import (
	"gorm.io/gorm"
	"sykros-pro/gopro/src/database/model"
)

type restaurantService interface {
	createRestaurant(restaurant *model.Restaurant) (error, model.Restaurant)
	getAllRestaurant(db *gorm.DB) []*RestaurantDto
}

type RestaurantService struct {
	restaurantService
	Db *gorm.DB
}
