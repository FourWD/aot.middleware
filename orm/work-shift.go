package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type WorkShift struct { // กะ การเข้างาน
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel
	SourceID  string `db:"source_id" json:"source_id" gorm:"type:varchar(2);"`
	StartTime string `db:"start_time" json:"start_time" gorm:"type:varchar(50);comment:'เวลาเข้ากะ'"`
	EndTime   string `db:"end_time" json:"end_time" gorm:"type:varchar(50);comment:'เวลาเลิกกะ'"`
	orm.GormRowOrder
}
