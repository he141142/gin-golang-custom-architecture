package restaurant

import (
	"gorm.io/gorm"
	"sykros-pro/gopro/src/utils"
	"time"
)

func (r *RestaurantService) getAllRestaurant(db *gorm.DB) []*RestaurantDto{
	rawlSql := `SELECT * FROM restaurants`
	query := db.Raw(rawlSql)
	return r.getRestaurantsHelper(query)
}

func (r *RestaurantService) getRestaurantsHelper(db *gorm.DB) []*RestaurantDto {
	var dbResult []map[string]interface{}
	utils.RawSQLScanner(db, &dbResult)
	//var restaurantId = make(map[uint64]uint64)
	restaurantsList := make([]*RestaurantDto,0)

	for _, data := range dbResult {
			if id,found := data["id"]; found && id != nil{
				restaurantRes := &RestaurantDto{}
				restaurantRes.UpdatedAt = data["updated_at"].(time.Time)
				restaurantRes.OwnerId = data["owner_id"].(uint64)
				restaurantRes.Address = data["address"].(string)
				restaurantRes.Name = data["name"].(string)
				restaurantRes.CityId = data["city_id"].(uint64)
				restaurantRes.Lat = data["lat"].(float32)
				restaurantRes.ShippingFeePerKm = data["shipping_fee_per_km"].(float32)
				restaurantRes.CreatedAt = data["created_at"].(time.Time)
				restaurantsList = append(restaurantsList, restaurantRes)
			}
	}
	return restaurantsList
}
