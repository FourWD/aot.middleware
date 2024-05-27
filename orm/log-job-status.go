package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type LogJobStatus struct {
	ID string `json:"id" gorm:"type:varchar(36); unique;"`
	orm.GormModel
	SlipID      string `json:"slip_id" gorm:"type:varchar(36)"`
	DriverID    string `json:"driver_id" gorm:"type:varchar(36)"`
	JobStatusID string `json:"job_status_id" gorm:"type:varchar(10)"`
}
