package orm

import orm "github.com/FourWD/middleware/model"

type VehicleType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel
	Name     string `json:"name" query:"type_name"`
	NameEn   string `json:"name_en" query:"type_name_en"`
	RowOrder int    `json:"row_order" query:"row_order" gorm:"type:int"`
	orm.GormRowOrder
}
