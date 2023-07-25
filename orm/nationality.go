package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type Nationality struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(3);primary_key;"`
	orm.GormModel
	Name        string `db:"name" json:"name" gorm:"type:varchar(150)"`
	NameEn      string `db:"name_en" json:"name_en" gorm:"type:varchar(50)"`
	Nationality string `db:"nationality" json:"nationality" gorm:"type:varchar(50)"`
	RowOrder    int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
