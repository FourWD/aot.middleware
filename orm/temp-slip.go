package orm

type TempSlip struct {
	Slip
	DestinationDetail string `json:"destination_detail" query:"destination_detail" gorm:"type:varchar(255)"`
	RoomNo            string `json:"room_no" query:"room_no" gorm:"type:varchar(10)"`
	PickupPassenger   int    `json:"pickup_passenger" query:"pickup_passenger" gorm:"type:int(3)"`
	AirlineName       string `json:"airline_name" query:"airline_name" gorm:"type:varchar(255)"`
}
