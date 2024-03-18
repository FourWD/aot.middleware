package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Settlement struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	CounterID      string `db:"counter_id" json:"counter_id" gorm:"type:varchar(36);"`
	StartingTime   string `db:"starting_time" json:"starting_time" gorm:"type:varchar(50) ; comment:'เวลาเริ่ม' "`
	SettlementTime string `db:"settlement_time" json:"settlement_time" gorm:"type:varchar(50) ; comment:'เวลาตัดจ่ายยอด'"`
}
