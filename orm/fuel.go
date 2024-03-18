package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Fuel struct { //no CRUD
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name       string `db:"name" json:"name" gorm:"type:varchar(50)"`
	FuelTypeID string `db:"fuel_type_id" json:"fuel_type_id" gorm:"type:varchar(2);"`
	RowOrder   int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
