package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type Fleet struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(4);primary_key;"`
	orm.GormModel

	Name           string `db:"name" json:"name" gorm:"type:varchar(50)"`
	BusinessTypeID string `db:"business_type_id" json:"business_type_id" gorm:"type:varchar(36)"`

	RefCode  string `db:"ref_code" json:"ref_code" gorm:"type:varchar(50)"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
	Address  string `db:"address" json:"address" gorm:"type:varchar(150)"`
	PhoneNo  string `db:"phone_no" json:"phone_no" gorm:"type:varchar(10)"`
	Email    string `db:"email" json:"email" gorm:"type:varchar(50)"`
	Status   bool   `db:"status" json:"status" gorm:"type:tinyint(2)"`
}
