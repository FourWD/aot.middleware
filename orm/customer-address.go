package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type CustomerAddress struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	CustomerID string `json:"customer_id"  query:"customer_id" gorm:"type:varchar(36); "`

	IsHQ     bool   `json:"is_hq" query:"is_hq" gorm:"type:tinyint(1)"`
	Address  string `json:"address" query:"address" gorm:"type:text"`
	Postcode string `json:"postcode" query:"postcode" gorm:"type:varchar(5)"`
	PhoneNo  string `json:"phone_no" query:"phone_no" gorm:"type:varchar(10)"`
}
