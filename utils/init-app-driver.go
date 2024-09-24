package utils

import (
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
	// appDriver.SlipID = driverID
	appDriver.AppDriverStatusID = config.APP_DRIVER_STATUS.AVAILABLE
	_, err := common.FirebaseClient.Collection("drivers").Doc(driverID).Set(common.FirebaseCtx, appDriver)

	clearDriverSlip(driverID)

	return err
}

func clearDriverSlip(driverID string) error {
	updateData := map[string]interface{}{ // update slip
		"slip_id": "",
	}

	return common.Database.Model(&orm.Driver{}).Where("id = ?", driverID).Updates(updateData).Error
}
