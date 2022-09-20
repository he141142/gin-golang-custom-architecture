package restaurant

import (
	"gorm.io/gorm"
	"sykros-pro/gopro/src/utils"
	"time"
)

func (r *RestaurantService) GetAllRestaurant(db *gorm.DB) []*RestaurantDto {
	rawlSql := `SELECT * FROM restaurants`
	query := db.Raw(rawlSql)
	return r.getRestaurantsHelper(query)
}

func (r *RestaurantService) getRestaurantsHelper(db *gorm.DB) []*RestaurantDto {
	var dbResult []map[string]interface{}
	utils.RawSQLScanner(db, &dbResult)
	restaurantsList := make([]*RestaurantDto, 0)
	for _, data := range dbResult {
		if id, found := data["id"]; found && id != nil {
			restaurantRes := &RestaurantDto{}
			r.bindRestaurantBaseData(data, restaurantRes)
			restaurantsList = append(restaurantsList, restaurantRes)
		}
	}
	return restaurantsList
}

func (r *RestaurantService) bindRestaurantBaseData(source map[string]interface{}, bind *RestaurantDto) {
	for k, v := range source {
		switch {
		case k == "updated_at":
			if v != nil {
				bind.UpdatedAt = v.(time.Time)
			}
		case k == "created_at":
			if v != nil {
				bind.CreatedAt = v.(time.Time)
			}
		case k == "owner_id":
			if v != nil {
				bind.OwnerId = v.(uint64)
			}
		case k == "address":
			if v != nil {
				bind.Address = v.(string)
			}
		case k == "name":
			if v != nil {
				bind.Name = v.(string)
			}
		case k == "city_id":
			if v != nil {
				bind.CityId = v.(uint64)
			}
		case k == "lat":
			if v != nil {
				bind.Lat = v.(float32)
			}
		case k == "shipping_fee_per_km":
			if v != nil {
				bind.ShippingFeePerKm = v.(float32)
			}
		}
	}
}
