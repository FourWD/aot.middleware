package orm

import orm "github.com/FourWD/middleware/model"

type VehicleBrand struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name     string `db:"name" json:"type_name"`
	NameEn   string `db:"name_en" json:"type_name_en"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
