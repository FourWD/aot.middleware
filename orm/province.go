package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Province struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name           string `db:"name" json:"name" gorm:"type:varchar(150)"`
	NameEn         string `db:"name_en" json:"name_en" gorm:"type:varchar(150)"`
	ProvinceTypeID string `db:"province_type_id" json:"province_type_id" gorm:"type:varchar(2);"`
}
