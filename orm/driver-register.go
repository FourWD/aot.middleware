package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type DriverRegister struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DriverID          string `db:"driver_id" json:"driver_id" gorm:"default:null; type:varchar(36); "`
	VehicleID         string `db:"vehicle_id" json:"vehicle_id" gorm:"default:null; type:varchar(36); "`
	IsGetKey          bool   `db:"is_get_key" json:"is_get_key" gorm:"default:0; type:tinyint(1); comment:'รับกุญแจยัง' "`
	GetKeyAt          string `db:"get_key_at"  json:"get_key_at" gorm:"default:null; type:varchar(50); comment:'วันเวลารับกุญแจ' "`
	GetKeyApprover    string `db:"get_key_approver"  json:"get_key_approver" gorm:"default:null; type:varchar(50); comment:'ผู้อนุมัติรับกุญแจ' "`
	IsReturnKey       bool   `db:"is_return_key"  json:"is_return_key" gorm:"default:0; type:tinyint(1); comment:'คืนกุญแจยัง' "`
	ReturnKeyAt       string `db:"return_key_at"  json:"return_key_at" gorm:"default:null; type:varchar(50); comment:'วันเวลาคืนกุญแจ' "`
	ReturnKeyApprover string `db:"return_key_approver"  json:"return_key_approver" gorm:"default:null; type:varchar(50); comment:'ผู้อนุมัติคืนกุญแจ' "`
	WorkShiftID       string `db:"work_shift_id"  json:"work_shift_id" gorm:"default:null; type:varchar(2); "`
}
