package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Report struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	ReportCode string `db:"report_code" json:"report_code" gorm:"type:varchar(5)"`
	AotCode    string `db:"aot_code" json:"aot_code" gorm:"type:varchar(36)"`
	Name       string `json:"name" query:"name" db:"name" gorm:"type:varchar(150)"`
	RowOrder   int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
