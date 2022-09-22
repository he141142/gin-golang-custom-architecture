package restaurant

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sykros-pro/gopro/src/utils"
	"time"
)

func (r *RestaurantService) GetAllRestaurant(db *gorm.DB, p *utils.PaginateHelper) (*RestaurantDtoPaginated, error) {
	rawlSql := `SELECT * FROM restaurants`
	countRawlSql := fmt.Sprintf(`%s  LIMIT @limit OFFSET @offset; `, rawlSql)
	errChan := make(chan error)
	restaurantChan := make(chan []*RestaurantDto)
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
		restaurant := r.getRestaurantsHelper(query)
		restaurantChan <- restaurant
	}()

	var response = &RestaurantDtoPaginated{}
	for i := 0; i < 2; i++ {
		select {
		case res := <-restaurantChan:
			response.restaurants = res
		case res := <-paginateChan:
			response.Page = res.Page
			response.Size = res.Size
			response.TotalItems = res.TotalItems
			response.TotalPages = res.TotalPages
			cast := p.SetPagingParam(res, response)
			response = cast.(*RestaurantDtoPaginated)
		case res := <-errChan:
			return nil, errors.New(res.Error())
		}
	}

	return response, nil
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
				bind.Lat = v.(float64)
			}
		case k == "shipping_fee_per_km":
			if v != nil {
				bind.ShippingFeePerKm = v.(float64)
			}
		}
	}
}
