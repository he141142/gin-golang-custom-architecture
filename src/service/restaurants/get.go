package restaurants

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sykros-pro/gopro/src/common"
	"sykros-pro/gopro/src/dto"
	"sykros-pro/gopro/src/utils"
	"sykros-pro/gopro/src/utils/helper"
	"time"
)

func (r *RestaurantService) GetAllRestaurant(db *gorm.DB, p *utils.PaginateHelper) (*dto.RestaurantDtoPaginated, error) {
	rawlSql := `SELECT * FROM restaurants`
	countRawlSql := fmt.Sprintf(`%s  LIMIT @limit OFFSET @offset; `, rawlSql)
	errChan := make(chan error)
	restaurantChan := make(chan []*dto.RestaurantDto)
	paginateChan := make(chan *utils.PaginateDto)
	go func() {
		rawlCountSql := fmt.Sprintf(`select count(*) total from ( %s ) as result;`, rawlSql)
		countQuery := db.Raw(rawlCountSql)
		paginateDto, err := p.GetTotalItemsCount(countQuery)
		if err != nil {
			errChan <- err
		}
		paginateChan <- paginateDto
	}()

	go func() {
		query := db.Raw(countRawlSql, sql.Named("limit", p.Limit), sql.Named("offset", p.Offset))
		restaurant, err := r.getRestaurantsHelper(query)
		if err != nil {
			errChan <- err
		}
		restaurantChan <- restaurant
	}()

	var response = &dto.RestaurantDtoPaginated{}
	for i := 0; i < 2; i++ {
		select {
		case res := <-restaurantChan:
			response.Restaurants = res
		case res := <-paginateChan:
			response.Paginate = res
		case res := <-errChan:
			return nil, errors.New(res.Error())
		}
	}

	return response, nil
}

func (r *RestaurantService) getRestaurantsHelper(query *gorm.DB) ([]*dto.RestaurantDto, error) {
	var dbResult []map[string]interface{}
	helper.RawSQLScanner(query, &dbResult)
	fmt.Println(dbResult)
	restaurantsList := make([]*dto.RestaurantDto, 0)
	for _, data := range dbResult {
		if id, found := data["id"]; found && id != nil {
			restaurantRes := &dto.RestaurantDto{}
			r.bindRestaurantBaseData(data, restaurantRes)
			restaurantsList = append(restaurantsList, restaurantRes)
		} else {
			return nil, errors.New("No Restaurants Found")
		}
	}
	return restaurantsList, nil
}

func (r *RestaurantService) bindRestaurantBaseData(source map[string]interface{}, bind *dto.RestaurantDto) {
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
				bind.Addr = v.(string)
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
				bind.Lat = v.(float64)
			}
		case k == "shipping_fee_per_km":
			if v != nil {
				bind.ShippingFeePerKm = v.(float64)
			}
		}
	}
}

func (r *RestaurantService) GetRestaurantById(db *gorm.DB, id int) (*dto.RestaurantDto, error) {
	rawlSql := `SELECT * FROM restaurants where id = @id;`
	query := db.Raw(rawlSql, sql.Named("id", id))
	restaurant, err := r.getRestaurantsHelper(query)
	if err != nil {
		return nil, common.CanNotGetEntity("restaurants", err)
	}
	if len(restaurant) == 0 {
		return nil, common.CanNotGetEntity("restaurants", errors.New("Cant found"))
	}
	return restaurant[0], nil
}
