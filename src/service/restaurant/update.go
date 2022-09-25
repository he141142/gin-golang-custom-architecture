package restaurant

import (
	"gorm.io/gorm"
)

func (r *RestaurantService) UpdateRestaurantById(db *gorm.DB, id int, updateDto *UpdateRestaurantDto) (*RestaurantDto, error) {

	getRestaurant, err := r.GetRestaurantById(db, id)
	if err != nil {
		return nil, err
	}

	return getRestaurant, nil
}
