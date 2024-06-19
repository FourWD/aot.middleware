package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type DriverRegister struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DriverID          string    `json:"driver_id" query:"driver_id" gorm:"default:null; type:varchar(36); "`
	VehicleID         string    `json:"vehicle_id" query:"vehicle_id" gorm:"default:null; type:varchar(36); "`
	IsGetKey          bool      `json:"is_get_key" query:"is_get_key" gorm:"default:0; type:bool; comment:'รับกุญแจยัง'"`
	GetKeyAt          time.Time `json:"get_key_at"  query:"get_key_at" gorm:"default:null; comment:'วันเวลารับกุญแจ'"`
	StartKilometerAt  int       `json:"start_kilometer_at" query:"start_kilometer_at" gorm:"default:null; type:int comment:'กม. เริ่มต้น'"`
	GetKeyApprover    string    `json:"get_key_approver" query:"get_key_approver" gorm:"default:null; type:varchar(50); comment:'ผู้อนุมัติรับกุญแจ'"`
	IsReturnKey       bool      `json:"is_return_key" query:"is_return_key" gorm:"type:bool; comment:'คืนกุญแจยัง'"`
	ReturnKeyAt       time.Time `json:"return_key_at" query:"return_key_at" gorm:"default:null; comment:'วันเวลาคืนกุญแจ'"`
	ReturnKeyApprover string    `json:"return_key_approver"  query:"return_key_approver" gorm:"default:null; type:varchar(36); comment:'ผู้อนุมัติคืนกุญแจ'"`
	EndKilometerAt    int       `json:"end_kilometer_at" query:"end_kilometer_at" gorm:"default:null; type:int comment:'กม. สิิ้นสุด'"`
	IsAssignWork      bool      `json:"is_assign_work" query:"is_assign_work" gorm:"type:bool; comment:'รับงานหรือยัง'"`
	WorkShiftID       string    `json:"work_shift_id"  query:"work_shift_id" gorm:"default:null; type:varchar(2); "`
}
