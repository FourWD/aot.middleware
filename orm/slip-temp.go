package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type SlipTemp struct {
	orm.GormModel

	AppDriverStatusID string `json:"app_driver_status_id" query:"app_driver_status_id" firestore:"app_driver_status_id" gorm:"type:varchar(2);"`
	AppUserStatusID   string `json:"app_user_status_id" query:"app_user_status_id" firestore:"app_user_status_id" gorm:"type:varchar(2);"`
	HashData          string `json:"hash_data" query:"hash_data" firestore:"hash_data" gorm:"type:text"`

	SlipNo     string    `json:"slip_no" query:"slip_no" firestore:"slip_no" gorm:"type:varchar(36);index"`
	ReceiptNo  string    `json:"receipt_no" query:"receipt_no" firestore:"receipt_no" gorm:"type:varchar(50);index"`
	RrNo       string    `json:"rr_no" query:"rr_no" firestore:"rr_no" gorm:"type:varchar(50); index"`
	SlipDate   time.Time `json:"slip_date" query:"slip_date" firestore:"slip_date" gorm:"default:null;"`
	TravelDate time.Time `json:"travel_date" query:"travel_date" firestore:"travel_date" gorm:"default:null;"`

	SlipTypeID           string `json:"slip_type_id" query:"slip_type_id" firestore:"slip_type_id" gorm:"type:varchar(2);index"`
	SlipSubTypeID        string `json:"slip_sub_type_id" query:"slip_sub_type_id" firestore:"slip_sub_type_id" gorm:"type:varchar(2);"`
	SlipVehicleSubTypeID string `json:"slip_vehicle_sub_type_id" query:"slip_vehicle_sub_type_id" firestore:"slip_vehicle_sub_type_id" gorm:"type:varchar(2);index"`
	IsDuo                bool   `json:"is_duo" query:"is_duo" firestore:"is_duo" gorm:"default:0;type:bool;"`            //  ตั๋วไปกลับ
	DuoSlipID            string `json:"duo_slip_id" query:"duo_slip_id" firestore:"duo_slip_id" gorm:"type:varchar(36)"` // ตั๋วขาแรก
	SlipStatusID         string `json:"slip_status_id" query:"slip_status_id" firestore:"slip_status_id" gorm:"type:varchar(2);default:'00'"`
	// ========================================================================================
	IsPaid          bool      `json:"is_paid" query:"is_paid" firestore:"is_paid" gorm:"default:0; type:bool;"`
	PaidBy          string    `json:"paid_by" query:"paid_by" firestore:"paid_by" gorm:"type:varchar(36);"`
	IsConfirmPaid   bool      `json:"is_confirm_paid" query:"is_confirm_paid" firestore:"is_confirm_paid" gorm:"default:0; type:bool;"`
	ConfirmPaidBy   string    `json:"confirm_paid_by" query:"confirm_paid_by" firestore:"confirm_paid_by" gorm:"type:varchar(36);"`
	ConfirmPaidDate time.Time `json:"confirm_paid_date" query:"confirm_paid_date" firestore:"confirm_paid_date" gorm:"default:null;"`

	PaymentTypeID     string    `json:"payment_type_id" query:"payment_type_id" firestore:"payment_type_id" gorm:"type:varchar(2);"`
	PaymentDate       time.Time `json:"payment_date" query:"payment_date" firestore:"payment_date" gorm:"default:null;"`
	PaymentSlip       string    `json:"payment_slip" query:"payment_slip" firestore:"payment_slip" gorm:"type:varchar(256);"`
	PaymentRef        string    `json:"payment_ref" query:"payment_ref" firestore:"payment_ref" gorm:"type:varchar(50);default:null;"`
	RefundRef         string    `json:"refund_ref" query:"refund_ref" firestore:"refund_ref" gorm:"type:varchar(50);default:null;"`
	RefundRemark      string    `json:"refund_remark" query:"refund_remark" firestore:"refund_remark" gorm:"type:varchar(256);default:null;"`
	RefundAmount      float64   `json:"refund_amount" query:"refund_amount" firestore:"refund_amount" gorm:"default:0; type:decimal(16,4);"`
	CreditCardNo      string    `json:"credit_card_number" query:"credit_card_number" firestore:"credit_card_number" gorm:"default:null; type:varchar(36);"`
	CreditCardTypeID  string    `json:"credit_card_type_id" query:"credit_card_type_id" firestore:"credit_card_type_id" gorm:"type:varchar(2);"`
	CustomerPaymentID string    `json:"customer_payment_id" query:"customer_payment_id" firestore:"customer_payment_id" gorm:"default:null; type:varchar(36);"`
	// BankRefNo         string    `json:"bank_ref_number" query:"bank_ref_number" firestore:"bank_ref_number" gorm:"default:null; type:varchar(20);"`
	// ========================================================================================
	IsConfirm   bool      `json:"is_confirm" query:"is_confirm" firestore:"is_confirm" gorm:"default:0; type:bool;"`
	ConfirmBy   string    `json:"confirm_by" query:"confirm_by" firestore:"confirm_by" gorm:"type:varchar(36);"`
	ConfirmDate time.Time `json:"confirm_date" query:"confirm_date" firestore:"confirm_date" gorm:"default:null;"`
	// ========================================================================================
	SlipVoidTypeID string    `json:"slip_void_type_id" query:"slip_void_type_id" firestore:"slip_void_type_id" gorm:"default:'00'; type:varchar(2);index"`
	VoidRemark     string    `json:"void_remark" query:"void_remark" firestore:"void_remark" gorm:"type:varchar(500);"`
	VoidBy         string    `json:"void_by" query:"void_by" firestore:"void_by" gorm:"type:varchar(36);"`
	VoidDate       time.Time `json:"void_date" query:"void_date" firestore:"void_date" gorm:"default:null;"`
	IsVoidSlip     bool      `json:"is_void_slip" query:"is_void_slip" firestore:"is_void_slip" gorm:"default:0; type:bool;"`
	// ========================================================================================
	SlipCancelTypeID     string    `json:"slip_cancel_type_id" query:"slip_cancel_type_id" firestore:"slip_cancel_type_id" gorm:"default:'00';type:varchar(2);index"`
	SlipUserCancelTypeID string    `json:"slip_user_cancel_type_id" query:"slip_user_cancel_type_id" firestore:"slip_user_cancel_type_id" gorm:"default:'00';type:varchar(2)"`
	CancelRemark         string    `json:"cancel_remark" query:"cancel_remark" firestore:"cancel_remark" gorm:"type:varchar(500);"`
	CancelBy             string    `json:"cancel_by" query:"cancel_by" firestore:"cancel_by" gorm:"type:varchar(36);"`
	CancelDate           time.Time `json:"cancel_date" query:"cancel_date" firestore:"cancel_date" gorm:"default:null;"`
	// ========================================================================================
	CounterID       string  `json:"counter_id" query:"counter_id" firestore:"counter_id" gorm:"type:varchar(5)"`
	OriginPoiID     string  `json:"origin_poi_id" query:"origin_poi_id" firestore:"origin_poi_id" gorm:"type:varchar(36);"`
	OriginName      string  `json:"origin_name" query:"origin_name" firestore:"origin_name" gorm:"type:varchar(150);"`
	ForceOriginName string  `json:"force_origin_name" query:"force_origin_name" firestore:"force_origin_name" gorm:"type:text;"`
	OriginLatitude  float64 `json:"origin_latitude" query:"origin_latitude" firestore:"origin_latitude" gorm:"type:decimal(10,6)"`
	OriginLongitude float64 `json:"origin_longitude" query:"origin_longitude" firestore:"origin_longitude" gorm:"type:decimal(10,6)"`
	// OriginRemark         string  `json:"origin_remark" query:"origin_remark" firestore:"origin_remark" gorm:"type:text"`
	DestinationPoiID     string  `json:"destination_poi_id" query:"destination_poi_id" firestore:"destination_poi_id" gorm:"type:varchar(36);index"`
	DestinationName      string  `json:"destination_name" query:"destination_name" firestore:"destination_name" gorm:"type:varchar(150);"`
	ForceDestinationName string  `json:"force_destination_name" query:"force_destination_name" firestore:"force_destination_name" gorm:"type:text;"`
	DestinationLatitude  float64 `json:"destination_latitude" query:"destination_latitude" firestore:"destination_latitude" gorm:"type:decimal(10,6)"`
	DestinationLongitude float64 `json:"destination_longitude" query:"destination_longitude" firestore:"destination_longitude" gorm:"type:decimal(10,6)"`
	// DestinationRemark    string  `json:"destination_remark" query:"destination_remark" firestore:"destination_remark" gorm:"type:text"`
	Distance float64 `json:"distance" query:"distance" firestore:"distance" gorm:"type:decimal(16,4)"`

	// PriceRateID   string  `json:"price_rate_id" query:"price_rate_id" firestore:"price_rate_id" gorm:"type:varchar(36);"`
	PromotionID   string  `json:"promotion_id" query:"promotion_id" firestore:"promotion_id" gorm:"type:varchar(36);"`
	PromotionCode string  `json:"promotion_code" query:"promotion_code" firestore:"promotion_code" gorm:"type:varchar(36);"`
	Price         float64 `json:"price" query:"price" firestore:"price" gorm:"default:0; type:decimal(16,4);"`
	Discount      float64 `json:"discount" query:"discount" firestore:"discount" gorm:"default:0; type:decimal(16,4);"`
	Wht           float64 `json:"wht" query:"wht" firestore:"wht" gorm:"default:0; type:decimal(16,4);"`
	IsWht         bool    `json:"is_wht" query:"is_wht" firestore:"is_wht" gorm:"type:bool; default:false;"`
	Vat           float64 `json:"vat" query:"vat" firestore:"vat" gorm:"default:0; type:decimal(16,4);"`
	DropOffPrice  float64 `json:"drop_off_price" query:"drop_off_price" firestore:"drop_off_price" gorm:"default:0; type:decimal(16,4);"`
	NetPrice      float64 `json:"net_price" query:"net_price" firestore:"net_price" gorm:"default:0; type:decimal(16,4);"`
	WaitingCharge float64 `json:"waiting_charge" query:"waiting_charge" firestore:"waiting_charge" gorm:"default:0; type:decimal(16,4);"`
	Fee           float64 `json:"fee" query:"fee" firestore:"fee" gorm:"default:0; type:decimal(16,4);"`
	FeeVat        float64 `json:"fee_vat" query:"fee_vat" firestore:"fee_vat" gorm:"default:0; type:decimal(16,4);"`
	EDCBankID     string  `json:"edc_bank_id" query:"edc_bank_id" firestore:"edc_bank_id" gorm:"type:varchar(5);"`
	ReceiveAmount float64 `json:"receive_amount" query:"receive_amount" firestore:"receive_amount" gorm:"default:0; type:decimal(16,4);"`

	RefCode string `json:"ref_code" query:"ref_code" firestore:"ref_code" gorm:"type:varchar(36);"`

	// IsNewCustomer bool      `json:"is_new_customer" query:"is_new_customer" firestore:"is_new_customer" gorm:"type:bool"`
	CustomerID string `json:"customer_id" query:"customer_id" firestore:"customer_id" gorm:"type:varchar(36);"`
	// Code             string    `json:"code" query:"code" firestore:"code" gorm:"type:varchar(20) ; dafault:null ; index"`
	CompanyName string    `json:"company_name" query:"company_name" firestore:"company_name" gorm:"type:varchar(150)"`
	TaxNo       string    `json:"tax_no" query:"tax_no" firestore:"tax_no" gorm:"type:varchar(50)"`
	IsTax       bool      `json:"is_tax" query:"is_tax" firestore:"is_tax" gorm:"type:bool"`
	IsHQ        bool      `json:"is_hq" query:"is_hq" firestore:"is_hq" gorm:"type:bool"`
	Address     string    `json:"address" query:"address" firestore:"address" gorm:"type:text"`
	Postcode    string    `json:"postcode" query:"postcode" firestore:"postcode" gorm:"type:varchar(10)"`
	PhoneNo     string    `json:"phone_no" query:"phone_no" firestore:"phone_no" gorm:"type:varchar(50)"`
	AirlineName string    `json:"airline_name" query:"airline_name" firestore:"airline_name" gorm:"type:varchar(100)"`
	FlightNo    string    `json:"flight_no" query:"flight_no" firestore:"flight_no" gorm:"type:varchar(20)"`
	FlightDate  time.Time `json:"flight_date" query:"flight_date" firestore:"flight_date" gorm:"default:null;"`

	Remark string `json:"remark" query:"remark" firestore:"remark" gorm:"type:text"`
	// BookingNo   string    `json:"booking_no" query:"booking_no" firestore:"booking_no" gorm:"type:varchar(50);index"`
	BookingBy   string    `json:"booking_by" query:"booking_by" firestore:"booking_by" gorm:"type:varchar(36)"`
	BookingDate time.Time `json:"booking_date" query:"booking_date" firestore:"booking_date" gorm:"default:null;"`

	//---FLEET--
	FleetClientID   string `json:"fleet_client_id" query:"fleet_client_id" firestore:"fleet_client_id" gorm:"type:varchar(10); default:null;"`
	BookingHour     string `json:"booking_hour" query:"booking_hour" firestore:"booking_hour" gorm:"type:varchar(2);"`
	Passenger       int    `json:"passenger" query:"passenger" firestore:"passenger" gorm:"type:int"`
	NumberOfLuggage int    `json:"number_of_luggage" query:"number_of_luggage" firestore:"number_of_luggage" gorm:"type:int"`

	PassengerPrefixID  string `json:"passenger_prefix_id" query:"passenger_prefix_id" firestore:"passenger_prefix_id" gorm:"type:varchar(2);"`
	PassengerFirstname string `json:"passenger_firstname" query:"passenger_firstname" firestore:"passenger_firstname" gorm:"type:varchar(200);"`
	PassengerLastname  string `json:"passenger_lastname" query:"passenger_lastname" firestore:"passenger_lastname" gorm:"type:varchar(200);"`
	// PassengerCountryID     string `json:"passenger_country_id" query:"passenger_country_id" firestore:"passenger_country_id" gorm:"type:varchar(10);"`
	PassengerGenderID      string `json:"passenger_gender_id" query:"passenger_gender_id" firestore:"passenger_gender_id" gorm:"type:varchar(10);"`
	PassengerNationalityID string `json:"passenger_nationality_id" query:"passenger_nationality_id" firestore:"passenger_nationality_id" gorm:"type:varchar(10);"`
	PassengerEmail         string `json:"passenger_email" query:"passenger_email" firestore:"passenger_email" gorm:"type:varchar(255);"`
	PassengerPhone         string `json:"passenger_phone" query:"passenger_phone" firestore:"passenger_phone" gorm:"type:varchar(50);"`
	PassengerPassport      string `json:"passenger_passport" query:"passenger_passport" firestore:"passenger_passport" gorm:"type:varchar(30);"`

	ContactPersonPrefixID  string `json:"contact_person_prefix_id" query:"contact_person_prefix_id" firestore:"contact_person_prefix_id" gorm:"type:varchar(10);"`
	ContactPersonFirstname string `json:"contact_person_firstname" query:"contact_person_firstname" firestore:"contact_person_firstname" gorm:"type:varchar(200);"`
	ContactPersonLastname  string `json:"contact_person_lastname" query:"contact_person_lastname" firestore:"contact_person_lastname" gorm:"type:varchar(200);"`
	ContactPersonPhone     string `json:"contact_person_phone " query:"contact_person_phone" firestore:"contact_person_phone" gorm:"type:varchar(50)"`
	ContactPersonEmail     string `json:"contact_person_email" query:"contact_person_email" firestore:"contact_person_email" gorm:"type:varchar(255)"`

	SubPassengerFirstname1 string `json:"sub_passenger_firstname_1" query:"sub_passenger_firstname_1" firestore:"sub_passenger_firstname_1" gorm:"type:varchar(255);column:sub_passenger_firstname_1;"`
	SubPassengerLastname1  string `json:"sub_passenger_lastname_1" query:"sub_passenger_lastname_1" firestore:"sub_passenger_lastname_1" gorm:"type:varchar(255);column:sub_passenger_lastname_1;"`
	// SubPassengerDestination1 string `json:"sub_passenger_destination_1" query:"sub_passenger_destination_1" firestore:"sub_passenger_destination_1" gorm:"type:varchar(255);column:sub_passenger_destination_1;"`

	SubPassengerFirstname2 string `json:"sub_passenger_firstname_2" query:"sub_passenger_firstname_2" firestore:"sub_passenger_firstname_2" gorm:"type:varchar(255);column:sub_passenger_firstname_2;"`
	SubPassengerLastname2  string `json:"sub_passenger_lastname_2" query:"sub_passenger_lastname_2" firestore:"sub_passenger_lastname_2" gorm:"type:varchar(255);column:sub_passenger_lastname_2;"`
	// SubPassengerDestination2 string `json:"sub_passenger_destination_2" query:"sub_passenger_destination_2" firestore:"sub_passenger_destination_2" gorm:"type:varchar(255);column:sub_passenger_destination_2;"`

	// SubPassengerFirstname3   string `json:"sub_passenger_firstname_3" query:"sub_passenger_firstname_3" firestore:"sub_passenger_firstname_3" gorm:"type:varchar(255);column:sub_passenger_firstname_3;"`
	// SubPassengerLastname3    string `json:"sub_passenger_lastname_3" query:"sub_passenger_lastname_3" firestore:"sub_passenger_lastname_3" gorm:"type:varchar(255);column:sub_passenger_lastname_3;"`
	// SubPassengerDestination3 string `json:"sub_passenger_destination_3" query:"sub_passenger_destination_3" firestore:"sub_passenger_destination_3" gorm:"type:varchar(255);column:sub_passenger_destination_3;"`
	// ========================================================================================
	DropOff1       string `json:"drop_off_1" query:"drop_off_1" firestore:"drop_off_1" gorm:"varchar(255);column:drop_off_1"`
	IsCharge1      bool   `json:"is_charge_1" query:"is_charge_1" firestore:"is_charge_1" gorm:"default:0; type:bool;column:is_charge_1;"`
	DropOffRemark1 string `json:"drop_off_remark_1" query:"drop_off_remark_1" firestore:"drop_off_remark_1" gorm:"type:text;column:drop_off_remark_1"`

	DropOff2       string `json:"drop_off_2" query:"drop_off_2" firestore:"drop_off_2" gorm:"varchar(255);column:drop_off_2"`
	IsCharge2      bool   `json:"is_charge_2" query:"is_charge_2" firestore:"is_charge_2" gorm:"default:0; type:bool;column:is_charge_2;"`
	DropOffRemark2 string `json:"drop_off_remark_2" query:"drop_off_remark_2" firestore:"drop_off_remark_2" gorm:"type:text;column:drop_off_remark_2"`

	// DropOff3       string `json:"drop_off_3" query:"drop_off_3" firestore:"drop_off_3" gorm:"varchar(255);column:drop_off_3"`
	// IsCharge3      bool   `json:"is_charge_3" query:"is_charge_3" firestore:"is_charge_3" gorm:"default:0;type:bool;column:is_charge_3;"`
	// DropOffRemark3 string `json:"drop_off_remark_3" query:"drop_off_remark_3" firestore:"drop_off_remark_3" gorm:"type:text;column:drop_off_remark_3"`
	// ========================================================================================
	CarSeat int `json:"car_seat" query:"car_seat" firestore:"car_seat" gorm:"default:0; type:int;column:car_seat;"`
	// RemarkDriver string `json:"remark_driver" query:"remark_driver" firestore:"remark_driver" gorm:"type:text;"`
	// ========================================================================================
	AddressTax         string `json:"address_tax" query:"address_tax" firestore:"address_tax" gorm:"type:text"`                        //ที่อยู่ใบกำกับภาษี
	CompanyTaxName     string `json:"company_tax_name" query:"company_tax_name" firestore:"company_tax_name" gorm:"type:varchar(200)"` //ที่อยู่ใบกำกับภาษี
	DistrictTaxName    string `json:"district_tax_name" query:"district_tax_name" firestore:"district_tax_name" gorm:"type:varchar(200)"`
	SubDistrictTaxName string `json:"sub_district_tax_name" query:"sub_district_tax_name" firestore:"sub_district_tax_name" gorm:"type:varchar(200)"`
	ProvinceTaxName    string `json:"province_tax_name" query:"province_tax_name" firestore:"province_tax_name" gorm:"type:varchar(200)"`
	PostcodeTax        string `json:"postcode_tax" query:"postcode_tax" firestore:"postcode_tax" gorm:"type:varchar(10)"`
	CountryTaxID       string `json:"country_tax_id" query:"country_tax_id" firestore:"country_tax_id" gorm:"type:varchar(10)"`
	CompanyPhoneTax    string `json:"company_phone_tax" query:"company_phone_tax" firestore:"company_phone_tax" gorm:"type:varchar(50)"`
	// ========================================================================================
	FuelByDistance float64 `json:"fuel_by_distance" query:"fuel_by_distance" firestore:"fuel_by_distance" gorm:"type:decimal(10,6)"` // จำนวน ลิตร
	CostByDistance float64 `json:"cost_by_distance" query:"cost_by_distance" firestore:"cost_by_distance" gorm:"type:decimal(10,6)"` // จำนวน บาท
	// RentalRateID      string    `json:"rental_rate_id" query:"rental_rate_id" firestore:"rental_rate_id" gorm:"type:varchar(36)"`
	// RentalPrice       float64   `json:"rental_price" query:"rental_price" firestore:"rental_price" gorm:"type:decimal(16,4)"`
	// RentalFuelRateID  string    `json:"rental_fuel_rate_id" query:"rental_fuel_rate_id" firestore:"rental_fuel_rate_id" gorm:"type:varchar(20)"`
	// RentalFuelLitre   float64   `json:"rental_fuel_litre" query:"rental_fuel_litre" firestore:"rental_fuel_litre" gorm:"type:decimal(16,4)"`
	FuelAverage float64 `json:"fuel_average" query:"fuel_average" firestore:"fuel_average" gorm:"type:decimal(16,4);"`
	// ========================================================================================
	// IsOperationApprove   bool      `json:"is_operation_approve" query:"is_operation_approve" firestore:"is_operation_approve" gorm:"type:bool; default:false;"`
	// OperationApproveBy   string    `json:"operation_approve_by" query:"operation_approve_by" firestore:"operation_approve_by" gorm:"type:varchar(36)"`
	// OperationApproveDate time.Time `json:"operation_approve_date" query:"operation_approve_date" firestore:"operation_approve_date" gorm:"default:null;"`
	// UpdateApproveBy      string    `json:"update_approve_by" query:"update_approve_by" firestore:"update_approve_by" gorm:"type:varchar(36)"`
	// UpdateApproveDate    time.Time `json:"update_approve_date" query:"update_approve_date" firestore:"update_approve_date" gorm:"default:null;"`

	// ReconcileDate time.Time `json:"reconcile_date" query:"reconcile_date" firestore:"reconcile_date" gorm:"default:null;"`
	// ========================================================================================
	IsRequestEdit          bool      `json:"is_request_edit" query:"is_request_edit" firestore:"is_request_edit" gorm:"type:bool; default:false;"`
	RequestEditDate        time.Time `json:"request_edit_date" query:"request_edit_date" firestore:"request_edit_date" gorm:"default:null;"`
	RequestEditApproveBy   string    `json:"request_edit_approve_by" query:"request_edit_approve_by" firestore:"request_edit_approve_by" gorm:"type:varchar(36)"`
	RequestEditApproveDate time.Time `json:"request_edit_approve_date" query:"request_edit_approve_date" firestore:"request_edit_approve_date" gorm:"default:null;"`

	IsRequestCancel   bool      `json:"is_request_cancel " query:"is_request_cancel" firestore:"is_request_cancel" gorm:"type:bool; default:false;"`
	RequestCancelDate time.Time `json:"request_cancel_date" query:"request_cancel_date" firestore:"request_cancel_date" gorm:"default:null;"`
	// ========================================================================================
	AssignVehicleID string `json:"assign_vehicle_id" query:"assign_vehicle_id" firestore:"assign_vehicle_id" gorm:"type:varchar(36);"`
	// AssignVehicleBy string    `json:"assign_vehicle_by" query:"assign_vehicle_by" firestore:"assign_vehicle_by" gorm:"type:varchar(36);"`
	AssignDriverID string    `json:"assign_driver_id" query:"assign_driver_id" firestore:"assign_driver_id" gorm:"default:null; type:varchar(36); index"`
	AssignBy       string    `json:"assign_by" query:"assign_by" firestore:"assign_by" gorm:"type:varchar(36);"`
	AssignDate     time.Time `json:"assign_date" query:"assign_date" firestore:"assign_date" gorm:"default:null;"`
	// JobAssignStatusID string  `json:"job_assign_status_id" query:"job_assign_status_id" firestore:"job_assign_status_id" gorm:"default:'00'; type:varchar(36)"`
	// ========================================================================================
	// AcceptJobDate time.Time `json:"accept_job_date" query:"accept_job_date" firestore:"accept_job_date" gorm:"default:null;"`
	IsAccept        bool      `json:"is_accept" query:"is_accept" firestore:"is_accept" gorm:"default:0; type:bool;"`
	AcceptDate      time.Time `json:"accept_date" query:"accept_date" firestore:"accept_date" gorm:"default:null;"`
	AcceptLatitude  float64   `json:"accept_latitude" query:"accept_latitude" firestore:"accept_latitude" gorm:"type:decimal(10,6);"`
	AcceptLongitude float64   `json:"accept_longitude" query:"accept_longitude" firestore:"accept_longitude" gorm:"type:decimal(10,6);"`

	IsPickup        bool      `json:"is_pickup" query:"is_pickup" firestore:"is_pickup" gorm:"default:0; type:bool;"`
	PickupDate      time.Time `json:"pickup_date" query:"pickup_date" firestore:"pickup_date" gorm:"default:null;"`
	PickupLatitude  float64   `json:"pickup_latitude" query:"arrive_latitude" firestore:"arrive_latitude" gorm:"type:decimal(10,6);"`
	PickupLongitude float64   `json:"pickup_longitude" query:"arrive_longitude" firestore:"arrive_longitude" gorm:"type:decimal(10,6);"`

	IsArrive        bool      `json:"is_arrive" query:"is_arrive" firestore:"is_arrive" gorm:"default:0; type:bool;"`
	ArriveDate      time.Time `json:"arrive_date" query:"arrive_date" firestore:"arrive_date" gorm:"default:null;"`
	ArriveLatitude  float64   `json:"arrive_latitude" query:"arrive_latitude" firestore:"arrive_latitude" gorm:"type:decimal(10,6);"`
	ArriveLongitude float64   `json:"arrive_longitude" query:"arrive_longitude" firestore:"arrive_longitude" gorm:"type:decimal(10,6);"`

	IsDropOff1        bool      `json:"is_drop_off_1" query:"is_drop_off_1" firestore:"is_drop_off_1" gorm:"type:bool;column:is_drop_off_1"`
	DropOffDate1      time.Time `json:"drop_off_date_1" query:"drop_off_date_1" firestore:"drop_off_date_1" gorm:"default:null; column:drop_off_date_1;"`
	DropOffLatitude1  float64   `json:"drop_off_latitude_1" query:"drop_off_latitude_1" firestore:"drop_off_latitude_1" gorm:"type:decimal(10,6);column:drop_off_latitude_1"`
	DropOffLongitude1 float64   `json:"drop_off_longitude_1" query:"drop_off_longitude_1" firestore:"drop_off_longitude_1" gorm:"type:decimal(10,6);column:drop_off_longitude_1"`

	IsDropOff2        bool      `json:"is_drop_off_2" query:"is_drop_off_2" firestore:"is_drop_off_2" gorm:"type:bool;column:is_drop_off_2"`
	DropOffDate2      time.Time `json:"drop_off_date_2" query:"drop_off_date_2" firestore:"drop_off_date_2" gorm:"default:null; column:drop_off_date_2;"`
	DropOffLatitude2  float64   `json:"drop_off_latitude_2" query:"drop_off_latitude_2" firestore:"drop_off_latitude_2" gorm:"type:decimal(10,6);column:drop_off_latitude_2"`
	DropOffLongitude2 float64   `json:"drop_off_longitude_2" query:"drop_off_longitude_2" firestore:"drop_off_longitude_2" gorm:"type:decimal(10,6);column:drop_off_longitude_2"`

	// IsDropOff3        bool      `json:"is_drop_off_3" query:"is_drop_off_3" firestore:"is_drop_off_3" gorm:"type:bool;column:is_drop_off_3"`
	// DropOffDate3      time.Time `json:"drop_off_date_3" query:"drop_off_date_3" firestore:"drop_off_date_3" gorm:"default:null; column:drop_off_date_3;"`
	// DropOffLatitude3  float64   `json:"drop_off_latitude_3" query:"drop_off_latitude_3" firestore:"drop_off_latitude_3" gorm:"type:decimal(10,6);column:drop_off_latitude_3"`
	// DropOffLongitude3 float64   `json:"drop_off_longitude_3" query:"drop_off_longitude_3" firestore:"drop_off_longitude_3" gorm:"type:decimal(10,6);column:drop_off_longitude_3"`

	IsDeliver        bool      `json:"is_deliver" query:"is_deliver" firestore:"is_deliver" gorm:"default:0; type:bool;"`
	DeliverDate      time.Time `json:"deliver_date" query:"deliver_date" firestore:"deliver_date" gorm:"default:null;"`
	DeliverLatitude  float64   `json:"deliver_latitude" query:"deliver_latitude" firestore:"deliver_latitude" gorm:"type:decimal(10,6);"`
	DeliverLongitude float64   `json:"deliver_longitude" query:"deliver_longitude" firestore:"deliver_longitude" gorm:"type:decimal(10,6);"`

	IsComplete   bool      `json:"is_complete" query:"is_complete" firestore:"is_complete" gorm:"default:0; type:bool;"`
	CompleteDate time.Time `json:"complete_date" query:"complete_date" firestore:"complete_date" gorm:"default:null;"`

	ActualDistance float64 `json:"actual_distance" query:"actual_distance" firestore:"actual_distance" gorm:"type:decimal(16,4)"`
	Duration       int     `json:"duration" query:"duration" firestore:"duration" gorm:"type:int"`

	SlipExtraTypeID     string `json:"slip_extra_type_id" query:"slip_extra_type_id" firestore:"slip_extra_type_id" gorm:"type:varchar(2); default:'00';"`
	IsRequestLadyDriver bool   `json:"is_request_lady_driver" query:"is_request_lady_driver" firestore:"is_request_lady_driver" gorm:"type:bool; default:false;"`
	// ========================================================================================
}
