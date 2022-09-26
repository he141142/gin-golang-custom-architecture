package restaurants

import (
	"gorm.io/gorm"
	"sykros-pro/gopro/src/database/model"
)

func (rs *RestaurantService) CreateRestaurant(data *model.Restaurant, db *gorm.DB) (error, *model.Restaurant) {
	if err := db.Create(data).Error; err != nil {
		return err, nil
	}
	return nil, data
}
