package orm

type AppDriver struct {
	Slip
	SlipID             string `json:"slip_id" firestore:"slip_id" query:"slip_id" gorm:"-"`
	VehicleModelID     string `json:"vehicle_model_id" firestore:"vehicle_model_id" gorm:"-"`
	VehicleModelName   string `json:"vehicle_model_name" firestore:"vehicle_model_name" query:"vehicle_model_name" gorm:"-"`
	VehicleSubTypeID   string `json:"vehicle_sub_type_id" firestore:"vehicle_sub_type_id" gorm:"-"`
	VehicleSubTypeName string `json:"vehicle_sub_type_name" firestore:"vehicle_sub_type_name" gorm:"-"`
	IsOriginAirport    bool   `json:"is_origin_airport" firestore:"is_origin_airport" query:"is_origin_airport" gorm:"-"`
}

// VehicleName        string `json:"vehicle_name" firestore:"vehicle_name" query:"vehicle_name" gorm:"-"` // vehicle_model_name
