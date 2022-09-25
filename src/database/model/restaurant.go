package model

import "time"

type Restaurant struct {
	Id               int                    `json:"id" gorm:"column:id;"`
	Name             string                 `json:"name" gorm:"column:name;"`
	OwnerId          uint64                 `json:"owner_id" gorm:"column:owner_id;"`
	Addr             string                 `json:"address" gorm:"column:address;"`
	CityId           uint64                 `json:"city_id" gorm:"column:city_id;"`
	Lat              float64                `json:"lat" gorm:"column:lat;"`
	Lng              float64                `json:"lng" gorm:"column:lng;"`
	Logo             map[string]interface{} `json:"logo" gorm:"column:logo;"`
	ShippingFeePerKm float64                `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km;"`
	Status           int8                   `json:"status" gorm:"column:status;"`
	CreatedAt        time.Time              `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt        time.Time              `json:"updated_at" gorm:"column:updated_at;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}
