package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type LogCancelSlip struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	SlipID             string `json:"slip_id" query:"id" gorm:"type:varchar(36)"`
	DriverID           string `json:"driver_id" query:"id" gorm:"type:varchar(36)"`
	SlipCancelReasonID string `json:"slip_cancel_reason_id" query:"id" gorm:"type:varchar(36)"`
}
