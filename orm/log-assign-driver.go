package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type LogAssignDriver struct {
	ID string `json:"id" gorm:"type:varchar(36); unique;"`
	orm.GormModel
	SlipID   string `json:"slip_id" gorm:"type:varchar(36)"`
	DriverID string `json:"driver_id" gorm:"type:varchar(36)"`
	IsCancel bool   `json:"is_cancel" gorm:"type:tinyint(1)"`
}
