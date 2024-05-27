package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type PDPAConsent struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	CustomerID string `db:"customer_id"  json:"customer_id" gorm:"type:varchar(36) ; default:null ;"`
	// PdpaDetailID string `db:"pdpa_detail_id" json:"pdpa_detail_id" gorm:"type:varchar(36);"`
	PDPADetail01 string `db:"pdpa_detail_01" json:"pdpa_detail_01" gorm:"default:null; column:pdpa_detail_01"`
	PDPADetail02 string `db:"pdpa_detail_02" json:"pdpa_detail_02" gorm:"default:null; column:pdpa_detail_02"`
	PDPADetail03 string `db:"pdpa_detail_03" json:"pdpa_detail_03" gorm:"default:null; column:pdpa_detail_03"`
	PDPADetail04 string `db:"pdpa_detail_04" json:"pdpa_detail_04" gorm:"default:null; column:pdpa_detail_04"`
	PDPADetail05 string `db:"pdpa_detail_05" json:"pdpa_detail_05" gorm:"default:null; column:pdpa_detail_05"`

	IsAccept bool      `db:"is_accept" json:"is_accept" gorm:"type:tinyint(2) ; comment:'ลูกค้าเป็นคนกดรับหรือไม่รับ'"`
	AcceptAt time.Time `db:"accept_at"  json:"accept_at" gorm:"default:null; type:datetime; comment:'วันที่ลูกค้ากดทำรายการผ่านหน้าเว็บ/แอพ'"`

	// IsCancel bool `db:"is_cancel" json:"is_cancel" gorm:"type:tinyint(2) ; comment:'ลูกค้าโทรมายกเลิก pdpa'"`
	IsCancel01 bool `db:"is_cancel_01" json:"is_cancel_01" gorm:"type:tinyint(2); column:is_cancel_01"`
	IsCancel02 bool `db:"is_cancel_02"  json:"is_cancel_02" gorm:"type:tinyint(2); column:is_cancel_02"`
	IsCancel03 bool `db:"is_cancel_03" json:"is_cancel_03" gorm:"type:tinyint(2); column:is_cancel_03"`
	IsCancel04 bool `db:"is_cancel_04"  json:"is_cancel_04" gorm:"type:tinyint(2); column:is_cancel_04"`
	IsCancel05 bool `db:"is_cancel_05" json:"is_cancel_05" gorm:"type:tinyint(2); column:is_cancel_05"`

	// PDPACancelReasonID string    `db:"pdpa_cancel_reason_id"  json:"pdpa_cancel_reason_id" gorm:"type:varchar(2) ; default:null ;"`
	PDPACancelReason01 string    `db:"pdpa_cancel_reason_01" json:"pdpa_cancel_reason_01" gorm:"default:null; column:pdpa_cancel_reason_01"`
	PDPACancelReason02 string    `db:"pdpa_cancel_reason_02" json:"pdpa_cancel_reason_02" gorm:"default:null; column:pdpa_cancel_reason_02"`
	PDPACancelReason03 string    `db:"pdpa_cancel_reason_03" json:"pdpa_cancel_reason_03" gorm:"default:null; column:pdpa_cancel_reason_03"`
	PDPACancelReason04 string    `db:"pdpa_cancel_reason_04" json:"pdpa_cancel_reason_04" gorm:"default:null; column:pdpa_cancel_reason_04"`
	PDPACancelReason05 string    `db:"pdpa_cancel_reason_05" json:"pdpa_cancel_reason_05" gorm:"default:null; column:pdpa_cancel_reason_05"`
	CancelAt           time.Time `db:"cancel_at"  json:"cancel_at" gorm:"default:null; type:datetime;"`
	// Remark             string    `db:"remark"  json:"remark" gorm:"type:varchar(200) ; default:null ;"`
	Remark01 string `db:"remark_01" json:"remark_01" gorm:"default:null; column:remark_01"`
	Remark02 string `db:"remark_02" json:"remark_02" gorm:"default:null; column:remark_02"`
	Remark03 string `db:"remark_03" json:"remark_03" gorm:"default:null; column:remark_03"`
	Remark04 string `db:"remark_04" json:"remark_04" gorm:"default:null; column:remark_04"`
	Remark05 string `db:"remark_05" json:"remark_05" gorm:"default:null; column:remark_05"`
}
