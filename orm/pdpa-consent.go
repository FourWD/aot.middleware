package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type PDPAConsent struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	CustomerID         string    `json:"customer_id"  query:"customer_id" gorm:"type:varchar(36);default:null ;"`
	PDPADetail01       string    `json:"pdpa_detail_01" query:"pdpa_detail_01" gorm:"default:null;column:pdpa_detail_01"`
	PDPADetail02       string    `json:"pdpa_detail_02" query:"pdpa_detail_02" gorm:"default:null;column:pdpa_detail_02"`
	PDPADetail03       string    `json:"pdpa_detail_03" query:"pdpa_detail_03" gorm:"default:null;column:pdpa_detail_03"`
	PDPADetail04       string    `json:"pdpa_detail_04" query:"pdpa_detail_04" gorm:"default:null;column:pdpa_detail_04"`
	PDPADetail05       string    `json:"pdpa_detail_05" query:"pdpa_detail_05" gorm:"default:null;column:pdpa_detail_05"`
	IsAccept           bool      `json:"is_accept" query:"is_accept" gorm:"type:tinyint(1);comment:'ลูกค้าเป็นคนกดรับหรือไม่รับ'"`
	AcceptAt           time.Time `json:"accept_at"  query:"accept_at"`
	IsCancel01         bool      `json:"is_cancel_01" query:"is_cancel_01" gorm:"type:tinyint(1);column:is_cancel_01"`
	IsCancel02         bool      `json:"is_cancel_02"  query:"is_cancel_02" gorm:"type:tinyint(1);column:is_cancel_02"`
	IsCancel03         bool      `json:"is_cancel_03" query:"is_cancel_03" gorm:"type:tinyint(1);column:is_cancel_03"`
	IsCancel04         bool      `json:"is_cancel_04"  query:"is_cancel_04" gorm:"type:tinyint(1);column:is_cancel_04"`
	IsCancel05         bool      `json:"is_cancel_05" query:"is_cancel_05" gorm:"type:tinyint(1);column:is_cancel_05"`
	PDPACancelReason01 string    `json:"pdpa_cancel_reason_01" query:"pdpa_cancel_reason_01" gorm:"default:null;column:pdpa_cancel_reason_01"`
	PDPACancelReason02 string    `json:"pdpa_cancel_reason_02" query:"pdpa_cancel_reason_02" gorm:"default:null;column:pdpa_cancel_reason_02"`
	PDPACancelReason03 string    `json:"pdpa_cancel_reason_03" query:"pdpa_cancel_reason_03" gorm:"default:null;column:pdpa_cancel_reason_03"`
	PDPACancelReason04 string    `json:"pdpa_cancel_reason_04" query:"pdpa_cancel_reason_04" gorm:"default:null;column:pdpa_cancel_reason_04"`
	PDPACancelReason05 string    `json:"pdpa_cancel_reason_05" query:"pdpa_cancel_reason_05" gorm:"default:null;column:pdpa_cancel_reason_05"`
	CancelAt           time.Time `json:"cancel_at"  query:"cancel_at" gorm:"default:null; type:datetime;"`
	Remark01           string    `json:"remark_01" query:"remark_01" gorm:"default:null; column:remark_01"`
	Remark02           string    `json:"remark_02" query:"remark_02" gorm:"default:null; column:remark_02"`
	Remark03           string    `json:"remark_03" query:"remark_03" gorm:"default:null; column:remark_03"`
	Remark04           string    `json:"remark_04" query:"remark_04" gorm:"default:null; column:remark_04"`
	Remark05           string    `json:"remark_05" query:"remark_05" gorm:"default:null; column:remark_05"`
}
