package restaurant

import (
	"gorm.io/gorm"
	"sykros-pro/gopro/src/database/model"
	"sykros-pro/gopro/src/share/logger"
	"sykros-pro/gopro/src/utils"
)

type restaurantService interface {
	createRestaurant(restaurant *model.Restaurant) (error, model.Restaurant)
	//GetAllRestaurant(db *gorm.DB) []*RestaurantDto
	GetAllRestaurant(db *gorm.DB, p *utils.PaginateHelper) (*RestaurantDtoPaginated, error)
	GetRestaurantById(db *gorm.DB,id int)
}

type RestaurantService struct {
	restaurantService
	Db     *gorm.DB
	Logger logger.LoggerService
}
