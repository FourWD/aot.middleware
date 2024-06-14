package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Nationality struct {
	ID string `query:"id" json:"id" gorm:"type:varchar(3);primary_key;"`
	orm.GormModel
	Name        string `query:"name" json:"name" gorm:"type:varchar(150)"`
	NameEn      string `query:"name_en" json:"name_en" gorm:"type:varchar(255)"`
	Nationality string `query:"nationality" json:"nationality" gorm:"type:varchar(50)"`
	orm.GormRowOrder
}
