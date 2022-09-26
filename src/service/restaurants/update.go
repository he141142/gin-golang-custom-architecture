package restaurants

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"sykros-pro/gopro/src/dto"
	"sykros-pro/gopro/src/utils/helper"
)

func (r *RestaurantService) UpdateRestaurantById(db *gorm.DB, id int, updateDto *dto.UpdateRestaurantDto) (*dto.RestaurantDto, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if validateErr := helper.DtoValidate(updateDto); validateErr != nil {
		return nil, errors.New(string(*validateErr))
	}
	_, err := r.GetRestaurantById(db, id)
	if err != nil {
		return nil, err
	}

	updateRestaurantSql := strings.Join(helper.RawlSqlHelper(updateDto), ";")
	fmt.Println(updateRestaurantSql)
	if err := tx.Exec(updateRestaurantSql, sql.Named("id", id)).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	getRestaurantAfterUpdate, _ := r.GetRestaurantById(db, id)

	return getRestaurantAfterUpdate, nil
}
