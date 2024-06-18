package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type FleetClient struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(10);primary_key;"`
	orm.GormModel

	FleetID string `json:"fleet_id" query:"fleet_id" gorm:"type:varchar(4)"`

	Name               string    `json:"name" query:"name" gorm:"type:varchar(150)"`
	RefCode            string    `json:"ref_code" query:"ref_code" gorm:"type:varchar(150)"`
	Remark             string    `json:"remark" query:"remark" gorm:"type:varchar(200)"`
	Address            string    `json:"address" query:"address" gorm:"type:varchar(150)"`
	PhoneNo            string    `json:"phone_no" query:"phone_no" gorm:"type:varchar(10)"`
	ContractStartDate  time.Time `json:"contract_start_date" query:"contract_start_date"`
	ContractEndDate    time.Time `json:"contract_end_date" query:"contract_end_date"`
	MinTicket          int       `json:"min_ticket" query:"min_ticket" gorm:"type:int"`
	Discount           float64   `json:"discount" query:"discount" gorm:"default:null;type:decimal(16,4)"`
	Email              string    `json:"email" query:"email" gorm:"type:varchar(100)"`
	Status             bool      `json:"status" query:"status" gorm:"type:bool"`
	ContactPerson      string    `json:"contact_person" query:"contact_person" gorm:"type:varchar(50)"`
	ContactPersonNo    string    `json:"contact_person_no" query:"contact_person_no" gorm:"type:varchar(10)"`
	ContactPersonEmail string    `json:"contact_person_email" query:"contact_person_email" gorm:"type:varchar(50)"`
	orm.GormRowOrder
}
