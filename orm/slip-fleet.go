package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type SlipFleet struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SlipNo      string    `db:"slip_no"  json:"slip_no" gorm:"type:varchar(50); index"`
	ReceiptNo   string    `db:"receipt_no"  json:"receipt_no" gorm:"type:varchar(50); index"`
	SlipAt      string    `db:"slip_at" json:"slip_at" gorm:"default:null; type:varchar(50); comment:'วันที่ลูกค้าจ่ายตัง' "`
	TravelAt    time.Time `db:"travel_at" json:"travel_at" gorm:"default:null;comment:'วันทีวิ่งจริง' "`
	ReconcileAt time.Time `db:"reconcile_at" json:"reconcile_at" gorm:"default:null;comment:'วันที่คีย์ตั๋วเทียบ' "`

	SlipTypeID            string `db:"slip_type_id" json:"slip_type_id" gorm:"type:varchar(2);"`
	SlipSubTypeID         string `db:"slip_sub_type_id" json:"slip_sub_type_id" gorm:"type:varchar(2);"`
	SlipVehicleSubModelID string `db:"slip_vehicle_model_id" json:"slip_vehicle_model_id" gorm:"type:varchar(2); comment:'ประเภทรถตามหน้าตั๋วที่ซื้อ'"`

	IsPickup  bool      `db:"is_pickup" json:"is_pickup" gorm:"default:0; type:tinyint(1); comment:'รับลูกค้าหรือยัง' "`
	PickupAt  time.Time `db:"pickup_at" json:"pickup_at"`
	CounterID string    `db:"counter_id" json:"counter_id" gorm:"type:varchar(36)"`

	OriginPoiID               string    `db:"origin_poi_id" json:"origin_poi_id" gorm:"type:varchar(36);"`
	ForceOriginName           string    `db:"force_origin_name" json:"force_origin_name" gorm:"type:varchar(150) ; comment:'กรณีที่ไม่มี POI' "`
	DestinationPoiID          string    `db:"destination_poi_id" json:"destination_poi_id" gorm:"type:varchar(36);"`
	ForceDestinationName      string    `db:"force_destination_name" json:"force_destination_name" gorm:"type:varchar(150) ; comment:'กรณีที่ไม่มี POI' "`
	Distance                  float64   `db:"distance" json:"distance" gorm:"type:decimal(16,4)"`
	PriceRateID               string    `db:"price_rate_id" json:"price_rate_id" gorm:"type:varchar(36);"`
	PromotionID               string    `db:"promotion_id" json:"promotion_id" gorm:"type:varchar(36);"`
	PromotionRef              string    `db:"promotion_ref" json:"promotion_ref" gorm:"type:varchar(36); comment:'กรณี duoslip เก็บเลขที่ slip ถ้าkansaiเก็บ เลขของ kansai' "`
	Price                     float64   `db:"price" json:"price" gorm:"default:null; type:decimal(16,4); comment:'ราคาค่าบริการ' "`
	Discount                  float64   `db:"discount" json:"discount" gorm:"default:null; type:decimal(16,4); comment:'ส่วนลด' "`
	Vat                       float64   `db:"vat" json:"vat" gorm:"default:null; type:decimal(16,4); comment:'ภาษี' "`
	NetPrice                  float64   `db:"net_price" json:"net_price" gorm:"default:null; type:decimal(16,4); comment:'รวมราคาค่าบริการทั้งหมด' "`
	IsPaid                    bool      `db:"is_paid" json:"is_paid" gorm:"default:0; type:tinyint(1); comment:'ลูกค้าจ่ายหรือยัง' "`
	PaymentTypeID             string    `db:"payment_type_id" json:"payment_type_id" gorm:"type:varchar(2);"`
	PaymentAt                 string    `db:"payment_at" json:"payment_at" gorm:"default:null; type:varchar(50); comment:'วันเวลาที่จ่าย' "`
	CreditCardNo              string    `db:"credit_card_number" json:"credit_card_number" gorm:"default:null; type:varchar(36);"`
	CreditCardTypeID          string    `db:"credit_card_type_id" json:"credit_card_type_id" gorm:"type:varchar(2);"`
	Latitude                  float64   `db:"lat" json:"lat" gorm:"type:decimal(10,6);comment:'ระยะพิกัดต้นทาง'"`
	Longitude                 float64   `db:"long" json:"long" gorm:"type:decimal(10,6);comment:'ระยะพิกัดปลายทาง'"`
	IsCancel                  bool      `db:"is_cancel" json:"is_cancel" gorm:"default:0;type:tinyint(1); comment:'ยกเลิก?' "`
	IsCancelTypeID            string    `db:"is_cancel_type_id" json:"is_cancel_type_id" gorm:"type:varchar(36);"`
	CancelRemark              string    `db:"cancel_remark" json:"cancel_remark" gorm:"type:varchar(500);"`
	CancelBy                  string    `db:"cancel_by" json:"cancel_by" gorm:"type:varchar(36);"`
	CancelAt                  time.Time `db:"cancel_at" json:"cancel_at"`
	IsVoid                    bool      `db:"is_void" json:"is_void" gorm:"default:0; type:tinyint(1);comment:'ยกเลิก?' "`
	IsVoidTypeID              string    `db:"is_void_type_id" json:"is_void_type_id" gorm:"type:varchar(36);"`
	VoidRemark                string    `db:"void_remark" json:"void_remark" gorm:"type:varchar(500);"`
	VoidBy                    string    `db:"void_by" json:"void_by" gorm:"type:varchar(36);"`
	VoidAt                    time.Time `db:"void_at" json:"void_at"`
	Code                      string    `db:"code"  json:"code" gorm:"type:varchar(20) ; default:null ; index"`
	FleetClientID             string    `db:"fleet_client_id"  json:"fleet_client_id" gorm:"type:varchar(10) ; default:null ; "`
	CompanyName               string    `db:"company_name" json:"company_name" gorm:"type:varchar(150)"`
	CompanyID                 string    `db:"company_id" json:"company_id" gorm:"type:varchar(10)"`
	BusinessTypeID            string    `db:"business_type_id" json:"business_type_id" gorm:"type:varchar(36)"`
	IsHQ                      bool      `db:"is_hq" json:"is_hq" gorm:"type:tinyint(1)"`
	Address                   string    `db:"address" json:"address" gorm:"type:text"`
	Postcode                  string    `db:"postcode" json:"postcode" gorm:"type:varchar(5)"`
	PhoneNo                   string    `db:"phone_no" json:"phone_no" gorm:"type:varchar(10)"`
	FlightNo                  string    `db:"flight_no" json:"flight_no" gorm:"type:varchar(20)"`
	FlightAt                  time.Time `db:"flight_at" json:"flight_at"`
	RentalRateID              string    `db:"rental_rate_id" json:"rental_rate_id" gorm:"type:varchar(36)"`
	RentalPrice               float64   `db:"rental_price" json:"rental_price" gorm:"type:decimal(16,4)"`
	RentalFuelRateID          string    `db:"rental_fuel_rate_id" json:"rental_fuel_rate_id" gorm:"type:varchar(20)"`
	RentalFuelLitre           float64   `db:"rental_fuel_litre" json:"rental_fuel_litre" gorm:"type:decimal(16,4)"`
	FuelAverage               float64   `db:"fuel_average" json:"fuel_average" gorm:"type:decimal(16,4); comment:'เก็บราคาของค่าน้ำมันเฉลี่ย '"`
	Hour                      int       `db:"hour" json:"hour" gorm:"int; comment:'ระยะเวลาเช่าของระบบ Fleet '"`
	AssignVehicleID           string    `db:"assign_vehicle_id" json:"assign_vehicle_id" gorm:"type:varchar(36); comment:'รถที่วิ่งงานจริง'"`
	AssignVehicleBy           string    `db:"assign_vehicle_by" json:"assign_vehicle_by" gorm:"type:varchar(36);"`
	AssignDriverID            string    `db:"assign_driver_id" json:"assign_driver_id" gorm:"default:null; type:varchar(36); "`
	ArrivedAt                 time.Time `db:"arrived_at" json:"arrived_at" gorm:"default:null; type:varchar(50); comment:'วันเวลาที่มาถึง' "`
	IsCompleted               bool      `db:"is_completed" json:"is_completed" gorm:"default:0; type:tinyint(1); comment:'เสร็จสมบูรณ์' "`
	Remark                    string    `db:"remark" json:"remark" gorm:"type:varchar(255)"`
	BookingNo                 string    `db:"booking_no"  json:"booking_no" gorm:"type:varchar(50); index"`
	BookingBy                 string    `db:"booking_by" json:"booking_by" gorm:"type:varchar(36); comment:'จองโดย'"`
	BookingAt                 time.Time `db:"booking_at" json:"booking_at"`
	PassengerName             string    `db:"passenger_name"  json:"passenger_name" gorm:"type:varchar(50);"`
	PassengerLastName         string    `db:"passenger_last_name"  json:"passenger_last_name" gorm:"type:varchar(200);"`
	ContactPersonName         string    `db:"contact_person_name"  json:"contact_person_name" gorm:"type:varchar(200);"`
	ContactPersonPhone        string    `db:" contact_person_phone " json:" contact_person_phone " gorm:"type:varchar(20)"`
	SubPassengerName01        string    `db:"sub_passenger_name_01"  json:"sub_passenger_name_01" gorm:"type:varchar(50);"`
	SubPassengerLastName01    string    `db:"sub_passenger_lastname_01"  json:"sub_passenger_lastname_01" gorm:"type:varchar(50);"`
	SubPassengerDestination01 string    `db:"sub_passenger_destination_01"  json:"sub_passenger_destination_01" gorm:"type:varchar(150);"`
	IsCharge01                bool      `db:"is_charge_01" json:"is_charge_01" gorm:"default:0; type:tinyint(1);"`
	SubPassengerName02        string    `db:"sub_passenger_name_02"  json:"sub_passenger_name_02" gorm:"type:varchar(50);"`
	SubPassengerLastName02    string    `db:"sub_passenger_lastname_02"  json:"sub_passenger_lastname_02" gorm:"type:varchar(50);"`
	SubPassengerDestination02 string    `db:"sub_passenger_destination_02"  json:"sub_passenger_destination_02" gorm:"type:varchar(150);"`
	IsCharge02                bool      `db:"is_charge_02" json:"is_charge_02" gorm:"default:0; type:tinyint(1);"`
	SubPassengerName03        string    `db:"sub_passenger_name_03"  json:"sub_passenger_name_03" gorm:"type:varchar(50);"`
	SubPassengerLastName03    string    `db:"sub_passenger_lastname_03"  json:"sub_passenger_lastname_03" gorm:"type:varchar(50);"`
	SubPassengerDestination03 string    `db:"sub_passenger_destination_03"  json:"sub_passenger_destination_03" gorm:"type:varchar(150);"`
	IsCharge03                bool      `db:"is_charge_03" json:"is_charge_03" gorm:"default:0; type:tinyint(1);"`
}
