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
	GetKeyDate        time.Time `json:"get_key_date"  query:"get_key_date" gorm:"default:null; comment:'วันเวลารับกุญแจ'"`
	StartKilometer    int       `json:"start_kilometer" query:"start_kilometer" gorm:"default:null; type:int comment:'กม. เริ่มต้น'"`
	GetKeyApprover    string    `json:"get_key_approver" query:"get_key_approver" gorm:"default:null; type:varchar(50); comment:'ผู้อนุมัติรับกุญแจ'"`
	IsReturnKey       bool      `json:"is_return_key" query:"is_return_key" gorm:"type:bool; comment:'คืนกุญแจยัง'"`
	ReturnKeyDate     time.Time `json:"return_key_date" query:"return_key_date" gorm:"default:null; comment:'วันเวลาคืนกุญแจ'"`
	ReturnKeyApprover string    `json:"return_key_approver"  query:"return_key_approver" gorm:"default:null; type:varchar(36); comment:'ผู้อนุมัติคืนกุญแจ'"`
	EndKilometer      int       `json:"end_kilometer" query:"end_kilometer" gorm:"default:null; type:int comment:'กม. สิิ้นสุด'"`
	SlipID            string    `json:"slip_id" query:"slip_id" gorm:"type:varchar(36); comment:'รับงานอะไรอยู่'"`
	WorkShiftID       string    `json:"work_shift_id"  query:"work_shift_id" gorm:"default:null; type:varchar(2);"`
}
