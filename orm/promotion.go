package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type Promotion struct { //no CRUD 01 = เงินสด 02 = creditcard 03= chuqe 04 = banktransfer
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	FleetClientID string `json:"fleet_client_id"  query:"fleet_client_id" gorm:"type:varchar(10);"`

	StartDate       time.Time `json:"start_time" query:"start_time"`
	EndDate         time.Time `json:"end_time" query:"end_time"`
	Name            string    `json:"name" query:"name" gorm:"type:varchar(150)"`
	PromotionType   string    `json:"promotion_type" query:"promotion_type" gorm:"type:varchar(5);"`
	PromotionAmount float64   `json:"promotion_amount" query:"promotion_amount" gorm:"type:decimal(10,4)"`
	orm.GormRowOrder
}
