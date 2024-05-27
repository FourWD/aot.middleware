package orm

import orm "github.com/FourWD/middleware/model"

type VehicleBrand struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name   string `json:"name" query:"type_name"`
	NameEn string `json:"name_en" query:"type_name_en"`
	orm.GormRowOrder
}
