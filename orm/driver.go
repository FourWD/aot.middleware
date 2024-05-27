package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type Driver struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	SourceID                string    `db:"source_id" json:"source_id" gorm:"type:varchar(2);"`
	DriverCode              string    `db:"driver_code" json:"driver_code" gorm:"type:varchar(36)"`
	Firstname               string    `db:"firstname" json:"firstname" gorm:"type:varchar(100)"`
	Lastname                string    `db:"lastname" json:"lastname" gorm:"type:varchar(100)"`
	Nickname                string    `db:"nickname" json:"nickname" gorm:"type:varchar(50)"`
	Birthday                time.Time `db:"birthday" json:"birthday" `
	Email                   string    `db:"email" json:"email" gorm:"type:varchar(50)"`
	IDCardNo                string    `db:"idcard_no" json:"idcard_no" gorm:"type:varchar(13)"`
	DriverLicenseNo         string    `db:"driver_license_no" json:"driver_license_no" gorm:"type:varchar(13)"`
	DriverLicenseExpireDate time.Time `db:"driver_license_expire_date" json:"driver_license_expire_date" gorm:"type:varchar(50)"`
	PhoneNo                 string    `db:"phone_no" json:"phone_no" gorm:"type:varchar(10)"`
	Address                 string    `db:"address" json:"address" gorm:"type:varchar(150)"`
	WorkingStartDate        time.Time `db:"working_start_date" json:"working_start_date" `
	WorkingEndDate          time.Time `db:"working_end_date" json:"working_end_date" `
	Remark                  string    `db:"remark" json:"remark" gorm:"type:varchar(200)"`
	DefaultVehicleID        string    `db:"default_vehicle_id" json:"default_vehicle_id" gorm:"default:null;type:varchar(36); "`
	CurrentVehicleID        string    `db:"current_vehicle_id" json:"current_vehicle_id" gorm:"default:null;type:varchar(36); "`
}
