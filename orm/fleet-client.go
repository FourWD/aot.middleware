package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type FleetClient struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(10);primary_key;"`
	orm.GormModel

	FleetID string `db:"fleet_id" json:"fleet_id" gorm:"type:varchar(4)"`

	Name               string `db:"name" json:"name" gorm:"type:varchar(50)"`
	RefCode            string `db:"ref_code" json:"ref_code" gorm:"type:varchar(50)"`
	Remark             string `db:"remark" json:"remark" gorm:"type:varchar(200)"`
	RowOrder           int    `db:"row_order" json:"row_order" gorm:"type:int"`
	Address            string `db:"address" json:"address" gorm:"type:varchar(150)"`
	PhoneNo            string `db:"phone_no" json:"phone_no" gorm:"type:varchar(10)"` //
	Email              string `db:"email" json:"email" gorm:"type:varchar(50)"`
	Status             bool   `db:"status" json:"status" gorm:"type:bool"`
	ContactPerson      string `db:"contact_person" json:"contact_person" gorm:"type:varchar(50)"`
	ContactPersonNo    string `db:"contact_person_no" json:"contact_person_no" gorm:"type:varchar(10)"`
	ContactPersonEmail string `db:"contact_person_email" json:"contact_person_email" gorm:"type:varchar(50)"`
}
