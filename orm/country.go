package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Country struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(3);primary_key;"` // alpha3
	orm.GormModel

	Name          string `db:"name" json:"name" gorm:"type:varchar(150)"`
	NameEn        string `db:"name_en" json:"name_en" gorm:"type:varchar(150)"`
	Nationality   string `db:"nationality" json:"nationality" gorm:"type:varchar(150)"`
	NationalityEn string `db:"nationality_en" json:"nationality_en" gorm:"type:varchar(150)"`
}
