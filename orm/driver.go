package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type Driver struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	SourceID                string    `json:"source_id" query:"source_id" gorm:"type:varchar(2);"`
	DriverCode              string    `json:"driver_code" query:"driver_code" gorm:"type:varchar(36);unique;"`
	PrefixID                string    `json:"prefix_id" query:"prefix_id" firestore:"prefix_id" gorm:"type:varchar(2);"`
	Firstname               string    `json:"firstname" query:"firstname" gorm:"type:varchar(100)"`
	Lastname                string    `json:"lastname" query:"lastname" gorm:"type:varchar(100)"`
	Nickname                string    `json:"nickname" query:"nickname" gorm:"type:varchar(50)"`
	Birthday                time.Time `json:"birthday" query:"birthday" gorm:"default:null;"`
	Email                   string    `json:"email" query:"email" gorm:"type:varchar(100)"`
	IDCardNo                string    `json:"idcard_no" query:"idcard_no" gorm:"type:varchar(13)"`
	DriverLicenseNo         string    `json:"driver_license_no" query:"driver_license_no" gorm:"type:varchar(13)"`
	DriverLicenseExpireDate time.Time `json:"driver_license_expire_date" query:"driver_license_expire_date" gorm:"default:null;"`
	PhoneNo                 string    `json:"phone_no" query:"phone_no" gorm:"type:varchar(50)"`
	Address                 string    `json:"address" query:"address" gorm:"type:varchar(150)"`
	WorkingStartDate        time.Time `json:"working_start_date" query:"working_start_date" gorm:"default:null;"`
	WorkingEndDate          time.Time `json:"working_end_date" query:"working_end_date" gorm:"default:null;"`
	Remark                  string    `json:"remark" query:"remark" gorm:"type:varchar(200)"`
	DriverImagePath         string    `json:"driver_image_path" query:"driver_image_path" gorm:"type:varchar(200)"`
	DefaultVehicleModelID   string    `json:"default_vehicle_model_id" query:"default_vehicle_model_id" gorm:"default:null;type:varchar(2);"`
	CurrentVehicleID        string    `json:"current_vehicle_id" query:"current_vehicle_id" gorm:"default:null; type:varchar(36);"`
	SlipID                  string    `json:"slip_id" query:"slip_id" gorm:"type:varchar(36);"`
	GenderID                string    `json:"gender_id" query:"gender_id" gorm:"type:varchar(2);"`
	IsOnline                bool      `json:"is_online" query:"is_online" gorm:"type:bool"`
	IsActive                bool      `json:"is_active" query:"is_active" gorm:"type:bool"`
	IsAppActive             bool      `json:"is_app_active" query:"is_app_active" gorm:"type:bool"`
	LastOnlineDate          time.Time `json:"last_online_date" query:"last_online_date" gorm:"default:null;"`
	QueueNo                 string    `json:"queue_no" query:"queue_no" gorm:"type:varchar(10)"` // stamp from running_nos
	QueueDate               time.Time `json:"queue_date" query:"queue_date" gorm:"default:null;"`
	TerminalQueueDate       time.Time `json:"terminal_queue_date" query:"terminal_queue_date" gorm:"default:null;"`
	GateID                  string    `json:"gate_id" query:"gate_id" gorm:"type:varchar(2)"`
	ResignDate              time.Time `json:"resign_date" query:"resign_date" gorm:"default:null;"`
	RefCode                 string    `json:"ref_code" query:"ref_code" gorm:"type:varchar(10);"`
	Point                   float64   `json:"point" query:"point" gorm:"int;"` //แต้มทั้งหมดของ driver
	// StopLocationID          int       `json:"stop_location_id" query:"stop_location_id" gorm:"default:null; type:int;"`
}
