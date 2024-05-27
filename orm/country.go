package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Country struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(3);primary_key;"` // alpha3
	orm.GormModel

	Name          string `json:"name" query:"name" gorm:"type:varchar(150)"`
	NameEn        string `json:"name_en" query:"name_en" gorm:"type:varchar(150)"`
	Nationality   string `json:"nationality" query:"nationality" gorm:"type:varchar(150)"`
	NationalityEn string `json:"nationality_en" query:"nationality_en" gorm:"type:varchar(150)"`
	orm.GormRowOrder
}
