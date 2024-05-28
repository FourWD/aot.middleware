package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type VehicleAllocation struct { // จำนวนรถที่แบ่งไว้ขายบนเว็บไซต์
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	SourceID      string `json:"source_id" query:"source_id" gorm:"type:varchar(2);"`
	StartDate     string `json:"start_time" query:"start_time" gorm:"type:varchar(50) ; comment:'วันที่เริ่ม'"`
	EndDate       string `json:"end_time" query:"end_time" gorm:"type:varchar(50) ; comment:'วันที่จบ'"`
	Hour00        int    `json:"ihour00" query:"ihour00" gorm:"type:int(2); column:ihour00; comment:'รถล็อคไว้เวลา เที่ยงคืน'"`
	Hour01        int    `json:"ihour01" query:"ihour01" gorm:"type:int(2); column:ihour01; comment:'รถล็อคไว้เวลา ตี 1'"`
	Hour02        int    `json:"ihour02" query:"ihour02" gorm:"type:int(2); column:ihour02; comment:'รถล็อคไว้เวลา ตี 2'"`
	Hour03        int    `json:"ihour03" query:"ihour03" gorm:"type:int(2); column:ihour03; comment:'รถล็อคไว้เวลา ตี 3'"`
	Hour04        int    `json:"ihour04" query:"ihour04" gorm:"type:int(2); column:ihour04; comment:'รถล็อคไว้เวลา ตี 4'"`
	Hour05        int    `json:"ihour05" query:"ihour05" gorm:"type:int(2); column:ihour05; comment:'รถล็อคไว้เวลา ตี 5'"`
	Hour06        int    `json:"ihour06" query:"ihour06" gorm:"type:int(2); column:ihour06; comment:'รถล็อคไว้เวลา 6 โมงเช้า'"`
	Hour07        int    `json:"ihour07" query:"ihour07" gorm:"type:int(2); column:ihour07; comment:'รถล็อคไว้เวลา 7 โมงเช้า'"`
	Hour08        int    `json:"ihour08" query:"ihour08" gorm:"type:int(2); column:ihour08; comment:'รถล็อคไว้เวลา 8 โมงเช้า'"`
	Hour09        int    `json:"ihour09" query:"ihour09" gorm:"type:int(2); column:ihour09; comment:'รถล็อคไว้เวลา 9 โมงเช้า'"`
	Hour10        int    `json:"ihour10" query:"ihour10" gorm:"type:int(2); column:ihour10; comment:'รถล็อคไว้เวลา 10 โมงเช้า'"`
	Hour11        int    `json:"ihour11" query:"ihour11" gorm:"type:int(2); column:ihour11; comment:'รถล็อคไว้เวลา 11 โมงเช้า'"`
	Hour12        int    `json:"ihour12" query:"ihour12" gorm:"type:int(2); column:ihour12; comment:'รถล็อคไว้เวลา เที่ยงวัน'"`
	Hour13        int    `json:"ihour13" query:"ihour13" gorm:"type:int(2); column:ihour13; comment:'รถล็อคไว้เวลา บ่าย 1'"`
	Hour14        int    `json:"ihour14" query:"ihour14" gorm:"type:int(2); column:ihour14; comment:'รถล็อคไว้เวลา บ่าย 2'"`
	Hour15        int    `json:"ihour15" query:"ihour15" gorm:"type:int(2); column:ihour15; comment:'รถล็อคไว้เวลา บ่าย 3'"`
	Hour16        int    `json:"ihour16" query:"ihour16" gorm:"type:int(2); column:ihour16; comment:'รถล็อคไว้เวลา บ่าย 4'"`
	Hour17        int    `json:"ihour17" query:"ihour17" gorm:"type:int(2); column:ihour17; comment:'รถล็อคไว้เวลา 5 โมงเย็น'"`
	Hour18        int    `json:"ihour18" query:"ihour18" gorm:"type:int(2); column:ihour18; comment:'รถล็อคไว้เวลา 6 โมงเย็น'"`
	Hour19        int    `json:"ihour19" query:"ihour19" gorm:"type:int(2); column:ihour19; comment:'รถล็อคไว้เวลา 1 ทุ่ม'"`
	Hour20        int    `json:"ihour20" query:"ihour20" gorm:"type:int(2); column:ihour20; comment:'รถล็อคไว้เวลา 2 ทุ่ม'"`
	Hour21        int    `json:"ihour21" query:"ihour21" gorm:"type:int(2); column:ihour21; comment:'รถล็อคไว้เวลา 3 ทุ่ม'"`
	Hour22        int    `json:"ihour22" query:"ihour22" gorm:"type:int(2); column:ihour22; comment:'รถล็อคไว้เวลา 4 ทุ่ม'"`
	Hour23        int    `json:"ihour23" query:"ihour23" gorm:"type:int(2); column:ihour23; comment:'รถล็อคไว้เวลา 5 ทุ่ม'"`
	VehicleTypeID string `json:"vehicle_type_id" query:"vehicle_type_id" gorm:"type:varchar(36);"`
}
