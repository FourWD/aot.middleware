package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type VehicleAllocation struct { // จำนวนรถที่แบ่งไว้ขายบนเว็บไซต์
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	SourceID      string `db:"source_id" json:"source_id" gorm:"type:varchar(2);"`
	StartDate     string `db:"start_time" json:"start_time" gorm:"type:varchar(50) ; comment:'วันที่เริ่ม'"`
	EndDate       string `db:"end_time" json:"end_time" gorm:"type:varchar(50) ; comment:'วันที่จบ'"`
	Hour00        int    `db:"ihour00" json:"ihour00" gorm:"type:int(2);comment:'รถล็อคไว้เวลา เที่ยงคืน'"`
	Hour01        int    `db:"ihour01" json:"ihour01" gorm:"type:int(2);comment:'รถล็อคไว้เวลา ตี 1'"`
	Hour02        int    `db:"ihour02" json:"ihour02" gorm:"type:int(2);comment:'รถล็อคไว้เวลา ตี 2'"`
	Hour03        int    `db:"ihour03" json:"ihour03" gorm:"type:int(2);comment:'รถล็อคไว้เวลา ตี 3'"`
	Hour04        int    `db:"ihour04" json:"ihour04" gorm:"type:int(2);comment:'รถล็อคไว้เวลา ตี 4'"`
	Hour05        int    `db:"ihour05" json:"ihour05" gorm:"type:int(2);comment:'รถล็อคไว้เวลา ตี 5'"`
	Hour06        int    `db:"ihour06" json:"ihour06" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 6 โมงเช้า'"`
	Hour07        int    `db:"ihour07" json:"ihour07" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 7 โมงเช้า'"`
	Hour08        int    `db:"ihour08" json:"ihour08" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 8 โมงเช้า'"`
	Hour09        int    `db:"ihour09" json:"ihour09" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 9 โมงเช้า'"`
	Hour10        int    `db:"ihour10" json:"ihour10" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 10 โมงเช้า'"`
	Hour11        int    `db:"ihour11" json:"ihour11" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 11 โมงเช้า'"`
	Hour12        int    `db:"ihour12" json:"ihour12" gorm:"type:int(2);comment:'รถล็อคไว้เวลา เที่ยงวัน'"`
	Hour13        int    `db:"ihour13" json:"ihour13" gorm:"type:int(2);comment:'รถล็อคไว้เวลา บ่าย 1'"`
	Hour14        int    `db:"ihour14" json:"ihour14" gorm:"type:int(2);comment:'รถล็อคไว้เวลา บ่าย 2'"`
	Hour15        int    `db:"ihour15" json:"ihour15" gorm:"type:int(2);comment:'รถล็อคไว้เวลา บ่าย 3'"`
	Hour16        int    `db:"ihour16" json:"ihour16" gorm:"type:int(2);comment:'รถล็อคไว้เวลา บ่าย 4'"`
	Hour17        int    `db:"ihour17" json:"ihour17" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 5 โมงเย็น'"`
	Hour18        int    `db:"ihour18" json:"ihour18" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 6 โมงเย็น'"`
	Hour19        int    `db:"ihour19" json:"ihour19" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 1 ทุ่ม'"`
	Hour20        int    `db:"ihour20" json:"ihour20" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 2 ทุ่ม'"`
	Hour21        int    `db:"ihour21" json:"ihour21" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 3 ทุ่ม'"`
	Hour22        int    `db:"ihour22" json:"ihour22" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 4 ทุ่ม'"`
	Hour23        int    `db:"ihour23" json:"ihour23" gorm:"type:int(2);comment:'รถล็อคไว้เวลา 5 ทุ่ม'"`
	VehicleTypeID string `db:"vehicle_type_id" json:"vehicle_type_id" gorm:"type:varchar(36);"`
}
