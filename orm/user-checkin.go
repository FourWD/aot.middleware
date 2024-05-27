package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type UserCheckIn struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID           string    `db:"user_id" json:"user_id" gorm:"type:varchar(36);"`
	CheckInDatetime  time.Time `db:"check_in_at" json:"check_in_at"`
	CheckOutDatetime time.Time `db:"check_out_at" json:"check_out_at"`
	CounterID        string    `db:"counter_id" json:"counter_id" gorm:"type:varchar(36);"`
}
