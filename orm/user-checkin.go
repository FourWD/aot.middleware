package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type UserCheckIn struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID           string `db:"user_id" json:"user_id" gorm:"type:varchar(36);"`
	CheckInDatetime  string `db:"check_in_at" json:"check_in_at" gorm:"type:datetime;"`
	CheckOutDatetime string `db:"check_out_at" json:"check_out_at" gorm:"type:datetime;"`
	CounterID        string `db:"counter_id" json:"counter_id" gorm:"type:varchar(36);"`
}
