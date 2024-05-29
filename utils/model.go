package calculatorHandler

type GetPricePayload struct {
	FleetClientID string  `json:"fleet_client_id"`
	Distance      float64 `json:"distance"`
}

type Response struct {
	Vehicles []Vehicle `json:"vehicles"`
	Distance float64   `json:"distance"`
	// NetPrice float64   `json:"netprice"`
}

type Vehicle struct {
	VehicleSourceName   string  `json:"vehicle_source_name"`
	VehicleModelID      string  `json:"vehicle_model_id"`
	VehicleModelName    string  `json:"vehicle_model_name"`
	VehicleModelNameEn  string  `json:"vehicle_model_name_en"`
	VehicleSubModelName string  `json:"vehicle_sub_model_name_en"`
	VehicleSourceID     string  `json:"vehicle_source_id"`
	VehicleSubModelID   string  `json:"vehicle_sub_model_id"`
	Price               float64 `json:"price"`
	Discount            float64 `json:"discount"`
	NetPrice            float64 `json:"net_price"`
	VehicleImage        string  `json:"vehicle_image"`
	// by_vehicle_model_id_01
}
