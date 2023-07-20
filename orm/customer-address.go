package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type CustomerAddress struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	CustomerID string `db:"customer_id"  json:"customer_id" gorm:"type:varchar(36); "`

	IsHQ     bool   `db:"is_hq" json:"is_hq" gorm:"type:bool"`
	Address  string `db:"address" json:"address" gorm:"type:text"`
	Postcode string `db:"postcode" json:"postcode" gorm:"type:varchar(5)"`
	PhoneNo  string `db:"phone_no" json:"phone_no" gorm:"type:varchar(10)"`
}
