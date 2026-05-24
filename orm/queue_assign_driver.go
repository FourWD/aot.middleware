package orm

import (
	"time"

	"gorm.io/gorm"
)

type QueueAssignDriver struct {
	ID        string         `json:"id"  query:"id" gorm:"type:varchar(36);primary_key;"`
	CreatedAt time.Time      `json:"created_at" query:"created_at" firestore:"created_at" gorm:"<-:create"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" query:"deleted_at" firestore:"deleted_at" gorm:"index"`
	SlipID    string         `json:"slip_id" query:"slip_id"  gorm:"type:varchar(36);unique"`
	DriverID  string         `json:"driver_id" query:"driver_id" gorm:"type:varchar(36);"`
	VehicleID string         `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);"`
}
