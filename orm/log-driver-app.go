package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type LogDriverApp struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`

	model.GormModel

	SlipID string `json:"slip_id" query:"slip_id" gorm:"type:varchar(36)"`

	IsAccept     bool      `json:"is_accept" query:"is_accept" gorm:"type:bool"`
	AppceptDate  time.Time `json:"appcept_date" query:"appcept_date" `
	IsPickup     bool      `json:"is_pickup" query:"is_pickup" gorm:"type:bool"`
	PickDate     time.Time `json:"pick_date" query:"pick_date" `
	IsArrival    bool      `json:"is_arrival" query:"is_arrival" gorm:"type:bool"`
	ArrivalDate  time.Time `json:"arrival_date" query:"arrival_date" `
	IsDeliver    bool      `json:"is_deliver" query:"is_deliver" gorm:"type:bool"`
	DeliveryDate time.Time `json:"deliver_date" query:"deliver_date" `
	IsComplete   bool      `json:"is_complete" query:"is_complete" gorm:"type:bool"`
	CompleteDate time.Time `json:"complete_date" query:"complete_date" `
	IsReject     bool      `json:"is_reject" query:"is_reject" gorm:"type:bool"`
	RejectDate   time.Time `json:"reject_date" query:"reject_date" `
	IsCancel     bool      `json:"is_cancel" query:"is_cancel" gorm:"type:bool"`
	CancelDate   time.Time `json:"cancel_date" query:"cancel_date" `
}
