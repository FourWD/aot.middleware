package utils

import (
	"fmt"

	"github.com/FourWD/aot.middleware/config"
	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func InitAppDriver(driverID string) error {
	// doc, err := common.FirebaseClient.Collection("drivers").Doc(driverID).Get(common.FirebaseCtx)
	// if err != nil {
	// 	return err
	// }

	// var currentDriverData model.AppDriver
	// if err := doc.DataTo(&currentDriverData); err != nil {
	// 	return err
	// }

	// if currentDriverData.AppDriverStatusID == config.APP_DRIVER_STATUS.ON_JOB {
	// 	return nil
	// }

	appDriver := new(orm.AppDriver)
	appDriver.ID = driverID
	appDriver.AppDriverStatusID = config.APP_DRIVER_STATUS.AVAILABLE
	type Payload struct {
		VehicleModelID string `json:"vehicle_model_id"`
	}
	driver := new(Payload)
	sql := `select v.vehicle_model_id from drivers d
	left join vehicles v on d.current_vehicle_id = v.id where d.id = ?`
	common.Database.Raw(sql, driverID).Scan(driver)

	_, err := GetAppDriverFirebase(driverID)
	if err != nil {
		updateData := map[string]interface{}{
			"vehicle_model_id": driver.VehicleModelID,
		}
		docPath := fmt.Sprintf("drivers/%s", driverID)
		if err := common.FirebaseUpdate(common.FirebaseClient, docPath, updateData); err != nil {
			return err
		}

	}

	_, err = common.FirebaseClient.Collection("drivers").Doc(driverID).Set(common.FirebaseCtx, appDriver)

	clearDriverSlip(driverID)

	return err
}

func clearDriverSlip(driverID string) error {
	updateData := map[string]interface{}{ // update slip
		"slip_id": "",
	}

	return common.Database.Model(&orm.Driver{}).Where("id = ?", driverID).Updates(updateData).Error
}
