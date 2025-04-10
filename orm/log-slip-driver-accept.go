package orm

import (
	"time"
)

type LogSlipDriverAccept struct {
	ID        string    `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	SlipID    string    `json:"slip_id" query:"slip_id"  gorm:"type:varchar(36)"`
	DriverID  string    `json:"driver_id" query:"driver_id" gorm:"type:varchar(36)"`
	CreatedAt time.Time `json:"created_at" query:"created_at" gorm:"<-:create"`
	DeletedAt time.Time `json:"deleted_at" query:"deleted_at" gorm:"default:null;"`
}
