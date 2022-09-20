package restaurant

import (
	"sykros-pro/gopro/src/database/model"
	"gorm.io/gorm"
)

func (rs *RestaurantService) CreateRestaurant(data *model.Restaurant, db *gorm.DB) (error, *model.Restaurant) {
	if err := db.Create(data).Error; err != nil {
		return err, nil
	}
	return nil, data
}
