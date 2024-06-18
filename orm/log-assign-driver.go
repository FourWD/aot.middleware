package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type LogAssignDriver struct {
	ID string `json:"id"  query:"slip_id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	SlipID   string `json:"slip_id" query:"slip_id"  gorm:"type:varchar(36)"`
	DriverID string `json:"driver_id" query:"driver_id" gorm:"type:varchar(36)"`
	IsCancel bool   `json:"is_cancel" query:"is_cancel" gorm:"type:bool"`
}
