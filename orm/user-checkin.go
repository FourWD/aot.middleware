package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type UserCheckIn struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID           string    `json:"user_id" query:"user_id" gorm:"type:varchar(36);"`
	CheckInDatetime  time.Time `json:"check_in_at" query:"check_in_at"`
	CheckOutDatetime time.Time `json:"check_out_at" query:"check_out_at"`
	CounterID        string    `json:"counter_id" query:"counter_id" gorm:"type:varchar(36);"`
}
