package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type WorkShift struct { // กะ การเข้างาน
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel
	SourceID  string `json:"source_id" query:"source_id" gorm:"type:varchar(2);"`
	StartTime string `json:"start_time" query:"start_time" gorm:"type:varchar(50);comment:'เวลาเข้ากะ'"`
	EndTime   string `json:"end_time" query:"end_time" gorm:"type:varchar(50);comment:'เวลาเลิกกะ'"`
	orm.GormRowOrder
}
