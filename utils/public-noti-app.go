package utils

import (
	"time"

	"github.com/FourWD/aot.middleware/model"
	"github.com/FourWD/middleware/common"
)

func publicNotiApp(topic string, group string, noti *model.NotiApp) error {
	noti.ActionTime = time.Now()

	message := group + "@" + common.StructToString(noti)
	logFields := map[string]interface{}{
		"slip_id":    noti.SlipID,
		"vehicle_id": noti.VehicleID,
		"driver_id":  noti.DriverID,
		"user_id":    noti.UserID,
		"topic":      topic,
		"group":      group,
		"message":    message,
	}
	common.Log("publicNotiApp", logFields, "")

	if _, err := common.GooglePublicMessage(topic, message); err != nil {
		return err
	}

	return nil
}
