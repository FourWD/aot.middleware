package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func GetAppDriverFirebase(driverID string) (orm.AppDriver, error) {
	docPath := fmt.Sprintf("drivers/%s", driverID)
	println(docPath, "dddddpoaskdopasjiokdjansiplldijlnkasijlndnalsijldajsk;da")
	docRef := common.FirebaseClient.Doc(docPath)
	snap, err := docRef.Get(context.Background())
	if err != nil {
		return orm.AppDriver{}, errors.New("cannot get app driver")
	}

	var appDriver orm.AppDriver
	if err := snap.DataTo(&appDriver); err != nil {
		return orm.AppDriver{}, errors.New("cannot get app driver")
	}

	fmt.Printf("CustomerID AppDriver: %+v\n", appDriver.CustomerID)
	fmt.Printf("AssignVehicleID AppDriver: %+v\n", appDriver.AssignVehicleID)

	arriveDate, _ := snap.Data()["arrive_date"].(time.Time)
	pickupDate, _ := snap.Data()["pickup_date"].(time.Time)
	dropOffDate1, _ := snap.Data()["drop_off_date_1"].(time.Time)
	dropOffDate2, _ := snap.Data()["drop_off_date_2"].(time.Time)
	dropOffDate3, _ := snap.Data()["drop_off_date_3"].(time.Time)
	deliverDate, _ := snap.Data()["deliver_date"].(time.Time)
	completeDate, _ := snap.Data()["complete_date"].(time.Time)

	appDriver.ArriveDate = common.UTCToThailandTime(arriveDate.Truncate(time.Second))
	appDriver.PickupDate = common.UTCToThailandTime(pickupDate.Truncate(time.Second))
	appDriver.DropOffDate1 = common.UTCToThailandTime(dropOffDate1.Truncate(time.Second))
	appDriver.DropOffDate2 = common.UTCToThailandTime(dropOffDate2.Truncate(time.Second))
	appDriver.DropOffDate3 = common.UTCToThailandTime(dropOffDate3.Truncate(time.Second))
	appDriver.DeliverDate = common.UTCToThailandTime(deliverDate.Truncate(time.Second))
	appDriver.CompleteDate = common.UTCToThailandTime(completeDate.Truncate(time.Second))

	return appDriver, nil
}
