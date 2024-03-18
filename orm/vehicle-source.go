package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type VehicleSource struct { // 01 = VRP 02=BenzThonglor
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name     string `db:"name" json:"name" gorm:"type:varchar(150)"`
	NameEn   string `db:"name_en" json:"name_en" gorm:"type:varchar(200)"`
	PhoneNo  string `db:"phone_no" json:"position" gorm:"type:varchar(10)"`
	Remark   string `db:"remark" json:"remark" gorm:"type:varchar(200)"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
