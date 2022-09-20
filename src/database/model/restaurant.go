package model

type Restaurant struct {
	Id int `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:address;"`
}

func (Restaurant) TableName() string{
	return "restaurants"
}