package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type Promotion struct { //no CRUD 01 = เงินสด 02 = creditcard 03= chuqe 04 = banktransfer
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	FleetClientID string `db:"fleet_client_id"  json:"fleet_client_id" gorm:"type:varchar(10);"`

	StartDate       time.Time `db:"start_time" json:"start_time"`
	EndDate         time.Time `db:"end_time" json:"end_time"`
	Name            string    `db:"name" json:"name" gorm:"type:varchar(150)"`
	PromotionType   string    `db:"promotion_type" json:"promotion_type" gorm:"type:varchar(5);"`
	PromotionAmount float64   `db:"promotion_amount" json:"promotion_amount" gorm:"type:int(5);"`
	orm.GormRowOrder
}
