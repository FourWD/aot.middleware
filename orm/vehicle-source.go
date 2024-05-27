package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type VehicleSource struct { // 01 = VRP 02=BenzThonglor
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name    string `json:"name" query:"name" gorm:"type:varchar(150)"`
	NameEn  string `json:"name_en" query:"name_en" gorm:"type:varchar(200)"`
	PhoneNo string `json:"phone_no" query:"position" gorm:"type:varchar(10)"`
	Remark  string `json:"remark" query:"remark" gorm:"type:varchar(200)"`
	orm.GormRowOrder
}
