package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type DriverSlip struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	SlipNo                string    `json:"slip_no"  query:"slip_no" gorm:"type:varchar(50); index"`
	ReceiptNo             string    `json:"receipt_no"  query:"receipt_no" gorm:"type:varchar(50); index"`
	RrNo                  string    `json:"rr_no"  query:"rr_no" gorm:"type:varchar(50); index"`
	SlipAt                time.Time `json:"slip_at" query:"slip_at" gorm:"default:null; comment:'วันที่ลูกค้าจ่ายตัง' "`
	TravelAt              time.Time `json:"travel_at" query:"travel_at" gorm:"default:null; comment:'วันทีวิ่งจริง' "`
	ReconcileAt           time.Time `json:"reconcile_at" query:"reconcile_at" gorm:"default:null; comment:'วันที่คีย์ตั๋วเทียบ' "`
	SlipSubTypeID         string    `json:"slip_sub_type_id" query:"slip_sub_type_id" gorm:"type:varchar(2);"`
	SlipVehicleSubModelID string    `json:"slip_vehicle_sub_model_id" query:"slip_vehicle_sub_model_id" gorm:"type:varchar(36); comment:'ประเภทรถตามหน้าตั๋วที่ซื้อ'"`
	IsPickup              bool      `json:"is_pickup" query:"is_pickup" gorm:"default:0; type:bool; comment:'รับลูกค้าหรือยัง' "`
	PickupAt              time.Time `json:"pickup_at" query:"pickup_at" gorm:"default:null; comment:'วันเวลาที่รับลูกค้า'  "`
	CounterID             string    `json:"counter_id" query:"counter_id" gorm:"type:varchar(36)"`
	OriginPoiID           string    `json:"origin_poi_id" query:"origin_poi_id" gorm:"type:varchar(36);"`
	ForceOriginName       string    `json:"force_origin_name" query:"force_origin_name" gorm:"type:varchar(150) ; comment:'กรณีที่ไม่มี POI' "`
	DestinationPoiID      string    `json:"destination_poi_id" query:"destination_poi_id" gorm:"type:varchar(36);"`
	ForceDestinationName  string    `json:"force_destination_name" query:"force_destination_name" gorm:"type:varchar(150) ; comment:'กรณีที่ไม่มี POI' "`
	Distance              float64   `json:"distance" query:"distance" gorm:"type:decimal(16,4)"`
	PriceRateID           string    `json:"price_rate_id" query:"price_rate_id" gorm:"type:varchar(36);"`
	PromotionID           string    `json:"promotion_id" query:"promotion_id" gorm:"type:varchar(36);"`
	PromotionRef          string    `json:"promotion_ref" query:"promotion_ref" gorm:"type:varchar(36); comment:'กรณี duoslip เก็บเลขที่ slip ถ้าkansaiเก็บ เลขของ kansai' "`
	Price                 float64   `json:"price" query:"price" gorm:"default:null; type:decimal(16,4); comment:'ราคาค่าบริการ' "`
	Discount              float64   `json:"discount" query:"discount" gorm:"default:null; type:decimal(16,4); comment:'ส่วนลด' "`
	Wht                   float64   `json:"wht" query:"wht" gorm:"default:null; type:decimal(16,4); comment:'ภาษีหัก ที่จ่าย' "`
	Vat                   float64   `json:"vat" query:"vat" gorm:"default:null; type:decimal(16,4); comment:'ภาษี' "`
	NetPrice              float64   `json:"net_price" query:"net_price" gorm:"default:null; type:decimal(16,4); comment:'รวมราคาค่าบริการทั้งหมด' "`
	IsPaid                bool      `json:"is_paid" query:"is_paid" gorm:"default:0; type:bool; comment:'ลูกค้าจ่ายหรือยัง' "`
	RefCode               string    `json:"ref_code" query:"ref_code" gorm:"type:varchar(36) comment:'ref ในหน้า status fleet'"`
	PaymentTypeID         string    `json:"payment_type_id" query:"payment_type_id" gorm:"type:varchar(2);"`
	PaymentAt             time.Time `json:"payment_at" query:"payment_at" gorm:"default:null; comment:'วันเวลาที่จ่าย' "`
	CreditCardNo          string    `json:"credit_card_number" query:"credit_card_number" gorm:"default:null; type:varchar(36);"`
	CreditCardTypeID      string    `json:"credit_card_type_id" query:"credit_card_type_id" gorm:"type:varchar(2);"`
	BankRefNo             string    `json:"bank_ref_no" query:"bank_ref_no" gorm:"default:null; type:varchar(20);"`
	IsVoid                bool      `json:"is_void" query:"is_void" gorm:"default:0; type:bool; comment:'ยกเลิก?' "`
	IsVoidTypeID          string    `json:"is_void_type_id" query:"is_void_type_id" gorm:"type:varchar(36);"`
	VoidRemark            string    `json:"void_remark" query:"void_remark" gorm:"type:varchar(500);"`
	Voijsony              string    `json:"void_by" query:"void_by" gorm:"type:varchar(36);"`
	VoidAt                time.Time `json:"void_at" query:"void_at" gorm:"default:null; comment:'วันเวลาที่ยกเลิก' "`
	IsCancel              bool      `json:"is_cancel" query:"is_cancel" gorm:"default:0; type:bool; comment:'ยกเลิก?'"`
	IsCancelTypeID        string    `json:"is_cancel_type_id" query:"is_cancel_type_id" gorm:"type:varchar(36);"`
	CancelRemark          string    `json:"cancel_remark" query:"cancel_remark" gorm:"type:varchar(500);"`
	CancelBy              string    `json:"cancel_by" query:"cancel_by" gorm:"type:varchar(36);"`
	CancelAt              time.Time `json:"cancel_at" query:"cancel_at" gorm:"default:null; comment:'วันเวลาที่ยกเลิก' "`
	IsNewCustomer         bool      `json:"is_new_customer" query:"is_new_customer" gorm:"type:bool"`
	CustomerID            string    `json:"customer_id" query:"customer_id" gorm:"type:varchar(36);"`
	Code                  string    `json:"code"  query:"code" gorm:"type:varchar(20) ;default:null;index"`
	CompanyName           string    `json:"company_name" query:"company_name" gorm:"type:varchar(150)"`
	TaxNo                 string    `json:"tax_no" query:"tax_no" gorm:"type:varchar(20)"`
	IsHQ                  bool      `json:"is_hq" query:"is_hq" gorm:"type:bool"`
	Address               string    `json:"address" query:"address" gorm:"type:text"`
	Postcode              string    `json:"postcode" query:"postcode" gorm:"type:varchar(5)"`
	PhoneNo               string    `json:"phone_no" query:"phone_no" gorm:"type:varchar(10)"`
	FightNo               string    `json:"fight_no" query:"fight_no" gorm:"type:varchar(20)"`
	FightAt               time.Time `json:"fight_at" query:"fight_at" gorm:"default:null; comment:'วันเวลาที่บิน'"`
	RentalRateID          string    `json:"rental_rate_id" query:"rental_rate_id" gorm:"type:varchar(36)"`
	RentalPrice           float64   `json:"rental_price" query:"rental_price" gorm:"type:decimal(16,4)"`
	RentalFuelRateID      string    `json:"rental_fuel_rate_id" query:"rental_fuel_rate_id" gorm:"type:varchar(20)"`
	RentalFuelLitre       float64   `json:"rental_fuel_litre" query:"rental_fuel_litre" gorm:"type:decimal(16,4)"`
	// RentalFuelPrice  float64 `json:"rental_fuel_price" query:"rental_fuel_price" gorm:"type:decimal(16,4)"`
	AssignVehicleID string    `json:"assign_vehicle_id" query:"assign_vehicle_id" gorm:"type:varchar(36); comment:'รถที่วิ่งงานจริง'"`
	AssignVehicleBy string    `json:"assign_vehicle_by" query:"assign_vehicle_by" gorm:"type:varchar(36);"`
	AssignDriverID  string    `json:"assign_driver_id" query:"assign_driver_id" gorm:"default:null; type:varchar(36); "`
	ArrivedAt       time.Time `json:"arrived_at" query:"arrived_at" gorm:"default:null; comment:'วันเวลาที่มาถึง' "`
	IsComplete      bool      `json:"is_complete" query:"is_complete" gorm:"default:0; type:bool; comment:'เสร็จสมบูรณ์' "`
	Remark          string    `json:"remark" query:"remark" gorm:"type:varchar(150) comment:'คำอธิบาย'"`
	BookingNo       string    `json:"booking_no"  query:"booking_no" gorm:"type:varchar(50); index"`
	BookingBy       string    `json:"booking_by" query:"booking_by" gorm:"type:varchar(36); comment:'จองโดย'"`
	BookingAt       time.Time `json:"booking_at" query:"booking_at" gorm:"default:null;  comment:'วันที่จอง' "`
}
