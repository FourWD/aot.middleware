package orm

import (
	"github.com/FourWD/middleware/model"
)

type LogDriverLogin struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	model.GormModel
	DriverID   string `json:"driver_id" query:"driver_id" gorm:"type:varchar(36)"`
	VehicleID  string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	IsLogin    bool   `json:"is_login" query:"is_login" gorm:"type:bool"`
	SlipID     string `json:"slip_id" query:"slip_id" gorm:"type:varchar(36)"`
	CustomerID string `json:"customer_id" query:"customer_id" gorm:"type:varchar(36)"`
	Remark     string `json:"remark" query:"remark" gorm:"type:varchar(200)"`
}
