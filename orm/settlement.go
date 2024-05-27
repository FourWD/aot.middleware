package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Settlement struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	CounterID      string `json:"counter_id" query:"counter_id" gorm:"type:varchar(36);"`
	StartingTime   string `json:"starting_time" query:"starting_time" gorm:"type:varchar(50) ; comment:'เวลาเริ่ม' "`
	SettlementTime string `json:"settlement_time" query:"settlement_time" gorm:"type:varchar(50) ; comment:'เวลาตัดจ่ายยอด'"`
}
