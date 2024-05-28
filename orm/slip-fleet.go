package orm

type SlipFleet struct {
	Slip
	FleetClientID string `json:"fleet_client_id"  query:"fleet_client_id" gorm:"type:varchar(10) ; default:null ; "`
	Hour          string `query:"hour" json:"hour" gorm:"type:varchar(2); comment:'ระยะเวลาเช่าของระบบ Fleet '"`

	PassengerName             string `json:"passenger_name"  query:"passenger_name" gorm:"type:varchar(50);"`
	PassengerLastName         string `json:"passenger_last_name"  query:"passenger_last_name" gorm:"type:varchar(200);"`
	ContactPersonName         string `json:"contact_person_name"  query:"contact_person_name" gorm:"type:varchar(200);"`
	ContactPersonPhone        string `json:" contact_person_phone " query:" contact_person_phone " gorm:"type:varchar(20)"`
	SubPassengerName01        string `json:"sub_passenger_name_01"  query:"sub_passenger_name_01" gorm:"type:varchar(50);"`
	SubPassengerLastName01    string `json:"sub_passenger_lastname_01"  query:"sub_passenger_lastname_01" gorm:"type:varchar(50);"`
	SubPassengerDestination01 string `json:"sub_passenger_destination_01"  query:"sub_passenger_destination_01" gorm:"type:varchar(150);"`
	IsCharge01                bool   `json:"is_charge_01" query:"is_charge_01" gorm:"default:0; type:tinyint(1); column:is_charge_01;"`
	SubPassengerName02        string `json:"sub_passenger_name_02"  query:"sub_passenger_name_02" gorm:"type:varchar(50);"`
	SubPassengerLastName02    string `json:"sub_passenger_lastname_02"  query:"sub_passenger_lastname_02" gorm:"type:varchar(50);"`
	SubPassengerDestination02 string `json:"sub_passenger_destination_02"  query:"sub_passenger_destination_02" gorm:"type:varchar(150);"`
	IsCharge02                bool   `json:"is_charge_02" query:"is_charge_02" gorm:"default:0; type:tinyint(1); column:is_charge_02;"`
	SubPassengerName03        string `json:"sub_passenger_name_03"  query:"sub_passenger_name_03" gorm:"type:varchar(50);"`
	SubPassengerLastName03    string `json:"sub_passenger_lastname_03"  query:"sub_passenger_lastname_03" gorm:"type:varchar(50);"`
	SubPassengerDestination03 string `json:"sub_passenger_destination_03"  query:"sub_passenger_destination_03" gorm:"type:varchar(150);"`
	IsCharge03                bool   `json:"is_charge_03" query:"is_charge_03" gorm:"default:0; type:tinyint(1); column:is_charge_03;"`
}
