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
