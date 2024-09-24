package utils

import (
	"github.com/FourWD/aot.middleware/orm"

	"github.com/FourWD/middleware/common"
)

type Payload struct {
	orm.Slip
	DriverCode  string `json:"driver_code"`
	VehicleCode string `json:"vehicle_code"`
}

func SnapSlipToAppDriver(slipID string, driverID, appDriverStatusID string) *orm.AppDriver {
	type Payload struct {
		orm.Slip
		DriverCode  string `json:"driver_code"`
		VehicleCode string `json:"vehicle_code"`
	}

	slipPayload := new(Payload)
	sql := `SELECT CONCAT(d.firstname, ' ', d.lastname) as driver_name, d.driver_image_path,
	vc.name as vehicle_color_name, vm.name as vehicle_name,v.license_plate, s.*
        FROM slips s
    left join drivers d on s.assign_driver_id = d.id
    left join vehicles v on s.assign_vehicle_id = v.id
    LEFT JOIN vehicle_colors vc ON v.vehicle_color_id = vc.id
    LEFT JOIN vehicle_models vm ON s.slip_vehicle_sub_model_id = vm.id
        WHERE s.id = ?
    ORDER BY s.created_at DESC`
	common.Database.Raw(sql, slipID).Debug().Scan(&slipPayload)

	appDriver := new(orm.AppDriver)
	appDriver.ID = driverID
	appDriver.SlipID = slipID
	appDriver.Duration = slipPayload.Duration
	appDriver.CustomerID = slipPayload.CustomerID
	appDriver.AppDriverStatusID = appDriverStatusID
	appDriver.OriginName = slipPayload.OriginName
	appDriver.ForceOriginName = slipPayload.ForceOriginName
	appDriver.OriginLatitude = slipPayload.OriginLatitude
	appDriver.OriginLongitude = slipPayload.OriginLongitude
	appDriver.OriginRemark = slipPayload.OriginRemark
	appDriver.RemarkDriver = slipPayload.RemarkDriver
	appDriver.Remark = slipPayload.Remark
	appDriver.DestinationName = slipPayload.DestinationName
	appDriver.ForceDestinationName = slipPayload.ForceDestinationName
	appDriver.DestinationLatitude = slipPayload.DestinationLatitude
	appDriver.DestinationLongitude = slipPayload.DestinationLongitude
	appDriver.DestinationRemark = slipPayload.DestinationRemark
	appDriver.Distance = slipPayload.Distance
	appDriver.TravelDate = slipPayload.TravelDate
	appDriver.AssignDriverID = slipPayload.AssignDriverID
	appDriver.AssignVehicleID = slipPayload.AssignVehicleID
	appDriver.IsPickup = slipPayload.IsPickup
	appDriver.PickupDate = slipPayload.PickupDate
	appDriver.SlipTypeID = slipPayload.SlipTypeID
	appDriver.SlipNo = slipPayload.SlipNo
	appDriver.JobNo = slipPayload.JobNo
	appDriver.SlipVehicleModelID = slipPayload.SlipVehicleModelID
	appDriver.PassengerFirstname = slipPayload.PassengerFirstname
	appDriver.PassengerLastname = slipPayload.PassengerLastname

	appDriver.DropOff1 = slipPayload.DropOff1
	appDriver.DropOffLatitude1 = slipPayload.DropOffLatitude1
	appDriver.DropOffLongitude1 = slipPayload.DropOffLongitude1
	appDriver.DropOffRemark1 = slipPayload.DropOffRemark1
	appDriver.IsDropOff1 = slipPayload.IsDropOff1
	appDriver.DropOffDate1 = slipPayload.DropOffDate1
	appDriver.DropOff2 = slipPayload.DropOff2
	appDriver.DropOffLatitude2 = slipPayload.DropOffLatitude2
	appDriver.DropOffLongitude2 = slipPayload.DropOffLongitude2
	appDriver.IsDropOff2 = slipPayload.IsDropOff2
	appDriver.DropOffDate2 = slipPayload.DropOffDate2
	appDriver.DropOffRemark2 = slipPayload.DropOffRemark2
	appDriver.DropOff3 = slipPayload.DropOff3
	appDriver.DropOffLatitude3 = slipPayload.DropOffLatitude3
	appDriver.DropOffLongitude3 = slipPayload.DropOffLongitude3
	appDriver.IsDropOff3 = slipPayload.IsDropOff3
	appDriver.DropOffDate3 = slipPayload.DropOffDate3
	appDriver.DropOffRemark3 = slipPayload.DropOffRemark3
	appDriver.IsArrive = slipPayload.IsArrive
	appDriver.ArriveDate = slipPayload.ArriveDate
	appDriver.ArriveLatitude = slipPayload.ArriveLatitude
	appDriver.ArriveLongitude = slipPayload.ArriveLongitude
	appDriver.IsDeliver = slipPayload.IsDeliver
	appDriver.DeliverDate = slipPayload.DeliverDate
	appDriver.IsComplete = slipPayload.IsComplete
	appDriver.CompleteDate = slipPayload.CancelDate
	appDriver.AcceptJobDate = slipPayload.ArriveDate
	return appDriver
}
