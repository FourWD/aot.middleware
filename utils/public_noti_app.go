package utils

import (
	"time"

	"github.com/FourWD/aot.middleware/model"
	"github.com/FourWD/middleware/infra"
	"github.com/FourWD/middleware/kit"
)

func publicNotiApp(topic string, group string, noti *model.NotiApp) error {
	noti.ActionTime = time.Now()

	message := group + "@" + kit.StructToString(noti)
	logFields := map[string]interface{}{
		"slip_id":    noti.SlipID,
		"vehicle_id": noti.VehicleID,
		"driver_id":  noti.DriverID,
		"user_id":    noti.UserID,
		"topic":      topic,
		"group":      group,
		"message":    message,
	}
	infra.AppLog.Event("publicNotiApp", logFields, "")

	if _, err := infra.GooglePublicMessage(topic, message); err != nil {
		return err
	}

	return nil
}
