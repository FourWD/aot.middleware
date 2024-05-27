package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Expressway struct { //no CRUD
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel
	Name             string  `json:"name" query:"name" gorm:"type:varchar(50)"`
	Price            float64 `json:"price" query:"price" gorm:"default:null; type:decimal(16,4); comment:'ราคาค่าทางด่วน' "`
	ExpresswayTypeID string  `json:"expressway_type_id" query:"expressway_type_id" gorm:"type:varchar(2)"`
	orm.GormRowOrder
}
