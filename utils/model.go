package utils

import "time"

//--customer
type PayloadCustomer struct {
	CustomerID      string    `json:"customer_id"`
	Code            string    `json:"code"`
	PaymentTypeID   string    `json:"payment_type_id"`
	Promotion       string    `json:"promotion"`
	PrefixID        string    `json:"prefix_id"`
	Firstname       string    `json:"firstname"`
	Lastname        string    `json:"lastname"`
	RefCode         string    `json:"ref_code"`
	BirthDay        time.Time `json:"birthday"`
	IDCardNo        string    `json:"id_card_no"`
	CompanyName     string    `json:"company_name"`
	NationalityID   string    `json:"nationality_id"`
	PassportNo      string    `json:"passport_no"`
	RegisterFrom    string    `json:"register_from"`
	TaxNo           string    `json:"tax_no"`
	IsTax           bool      `json:"is_tax"`
	Telephone       string    `json:"telephone"`
	RunningNo       int       `json:"running_no"`
	BranchID        string    `json:"branch_id"`
	BranchName      string    `json:"branch_name"`
	AddressCompany  string    `json:"address_company"`
	PostCodeCompany string    `json:"postcode_company"`
	PhoneCompany    string    `json:"phone_company"`
	BookingDate     string    `json:"booking_date"`
}

type SlipPDF struct {
	ID                    string  `json:"id"`
	SlipNo                string  `json:"slip_no"`
	ReceiptNo             string  `json:"receipt_no"`
	SlipVehicleTypeID     string  `json:"slip_vehicle_type_id"`
	SlipVehicleSubModelID string  `json:"slip_vehicle_sub_model_id"`
	PaymentTypeID         string  `json:"payment_type_id"`
	OriginPoiID           string  `json:"origin_poi_id"`
	ForceOriginName       string  `json:"force_origin_name"`
	DestinationPoiID      string  `json:"destination_poi_id"`
	ForceDestinationName  string  `json:"force_destination_name"`
	Distance              float64 `json:"distance"`
	DistanceRoundTrip     float64 `json:"distance_round_trip"`
	Price                 float64 `json:"price"`
	Discount              float64 `json:"discount"`
	NetPrice              float64 `json:"net_price"`
	Litre                 float64 `json:"rental_fuel_litre"`
}
