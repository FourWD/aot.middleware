package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Province struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name           string `json:"name" query:"name" gorm:"type:varchar(150)"`
	NameEn         string `json:"name_en" query:"name_en" gorm:"type:varchar(150)"`
	ProvinceTypeID string `json:"province_type_id" query:"province_type_id" gorm:"type:varchar(2);"`
}
