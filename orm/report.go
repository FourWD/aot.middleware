package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Report struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	ReportCode string `json:"report_code" query:"report_code" gorm:"type:varchar(5)"`
	AotCode    string `json:"aot_code" query:"aot_code" gorm:"type:varchar(36)"`
	Name       string `json:"name" query:"name"  gorm:"type:varchar(150)"`
	orm.GormRowOrder
}
