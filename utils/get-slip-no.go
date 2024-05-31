package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func GetSlipNo(slip Payload) (string, error) {
	vehiclemodel, _ := getVehicleModel(slip.VehicleSubModelNameEn)
	counter := "ZZ"
	year := time.Now().Year()
	month := time.Now().Month()
	// runningnoStr, err := getRuning()
	// if err != nil {
	// 	// common.FiberReviewPayload(c)
	// 	return "", err
	// }
	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	slipno := vehiclemodel + "-" + counter + "-" + fyear + "-" + fmonth + "-"
	// + runningnoStr
	common.Print(fyear, fmonth)
	common.Print("vehiclemodel", vehiclemodel)
	return slipno, nil
}
func getRuning(payload Payload) (string, error) {
	// sql := `SELECT COUNT(*) from slips where running_no <=
	// (SELECT running_no from slips where id = '%s')
	// AND vehicle_sub_model_name = '%s'
	// AND counter = '%s'
	// AND

	// `

	runningslip := new(orm.Customer)

	getRunning := `SELECT * FROM slips ORDER BY running_no DESC LIMIT 1`
	common.Database.Raw(getRunning).Scan(&runningslip)

	common.PrintError("error running ", getRunning)
	return fmt.Sprintf(`%05d`, (runningslip.RunningNo + 1)), nil
}

func getVehicleModel(vehiclesubmodel string) (string, error) {

	vehicle := new(orm.VehicleSubModel)
	getvehicle := fmt.Sprintf(`SELECT * FROM vehicle_sub_models WHERE name_en = '%s'`, vehiclesubmodel)
	common.Database.Raw(getvehicle).Scan(&vehicle)

	common.PrintError("vehiclemodel", vehicle.NameEn)
	return vehicle.NameEn, nil
}

type Payload struct {
	ID            string `json:"id"`
	SlipNo        string `json:"slip_no"`
	ReceiptNo     string `json:"receip_no"`
	RrNo          string `json:"rr_no"`
	SlipAt        string `json:"slip_at"`
	TravelAt      string `json:"travel_at"`
	ReconcileAt   string `json:"reconcile_at"`
	SlipSubTypeID string `json:"slip_sub_type_id"`
	// SlipVehicleSubModelID string `json:"slip_vehicle_sub_model_id"`
	SlipTypeID string `json:"slip_type_id"`

	SlipVehicleSubModelID string  `json:"slip_vehicle_sub_model_id"`
	VehicleSubModelNameEn string  `json:"vehicle_sub_model_name"`
	VehicleSubModelID     string  `json:"vehicle_sub_model_id"`
	NameEn                string  `json:"name_en"`
	IsPickUp              int     `json:"is_pickup"`
	IsHQ                  int     `json:"is_hq"`
	IsTax                 string  `json:"is_tax"`
	CounterID             string  `json:"counter_id"`
	OriginPoiID           string  `json:"origin_poi_id"`
	ForceOriginName       string  `json:"force_origin_name"`
	DestinationPoiID      string  `json:"destination_poi_id"`
	ForceDestinationName  string  `json:"force_destination_name"`
	Address               string  `json:"address"`
	AddressCompany        string  `json:"address_company"`
	PostCodeCompany       string  `json:"postcode_company"`
	PhoneCompany          string  `json:"phone_company"`
	BookingAt             string  `json:"booking_at"`
	TimeAt                string  `json:"time_at"`
	Distance              string  `json:"distance"`
	PriceRateID           string  `json:"price_rate_id"`
	PromotionID           string  `json:"promotion_id"`
	PromotionName         string  `json:"promotion_name"`
	PromotionRef          string  `json:"promotion_ref"`
	Price                 float64 `json:"price"`
	Discount              float64 `json:"discount"`
	Wht                   float64 `json:"wht"`
	Vat                   float64 `json:"vat"`
	NetPrice              float64 `json:"net_price"`
	IsPaID                int     `json:"is_paid"`
	PaymentAt             string  `json:"payment_at"`
	CustomerID            string  `json:"customer_id"`
	Code                  string  `json:"code"`
	PaymentTypeID         string  `json:"payment_type_id"`
	Name                  string  `json:"name"`
	Promotion             string  `json:"promotion"`
	PrefixID              string  `json:"prefix_id"`
	Firstname             string  `json:"firstname"`
	Lastname              string  `json:"lastname"`
	RefCode               string  `json:"ref_code"`
	// BirthDay              string  `json:"birthday"`
	// Avatar                string  `json:"avatar"`
	// Email                 string  `json:"email"`
	IDCardNo      string `json:"id_card_no"`
	CompanyName   string `json:"company_name"`
	NationalityID string `json:"nationality_id"`
	PassportNo    string `json:"passport_no"`
	RegisterFrom  string `json:"register_from"`
	TaxNo         string `json:"tax_no"`
	Telephone     string `json:"telephone"`
	RunningNo     int    `json:"running_no"`
}

type GetRunning struct {
	RunningNo string `json:"running_no"`
}

type Counter struct {
	Name string `json:"name"`
}

type Vehicle struct {
	VehicleSubNameEn string `json:"vehicle_sub_name_en"`
}
