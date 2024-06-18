package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Fleet struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(4);primary_key;"`
	orm.GormModel

	Name           string `json:"name" query:"name" gorm:"type:varchar(50)"`
	BusinessTypeID string `json:"business_type_id" query:"business_type_id" gorm:"type:varchar(36)"`

	RefCode string `json:"ref_code" query:"ref_code" gorm:"type:varchar(50)"`
	Address string `json:"address" query:"address" gorm:"type:varchar(150)"`
	PhoneNo string `json:"phone_no" query:"phone_no" gorm:"type:varchar(20)"`
	Email   string `json:"email" query:"email" gorm:"type:varchar(100)"`
	Status  bool   `json:"status" query:"status" gorm:"type:bool"`
	orm.GormRowOrder
}
