package utils

import (
	"encoding/json"
	"errors"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/FourWD/aot.middleware/model"
	"github.com/FourWD/middleware/infra"
)

func GetNotiApp(message *pubsub.Message) (model.NotiApp, error) {
	infra.AppLog.Event("GetNotiApp Start", map[string]interface{}{
		"message.id":   message.ID,
		"message.data": string(message.Data),
	}, message.ID)

	gMessage := infra.ConventStringToGoogleMessage(string(message.Data))

	var noti model.NotiApp
	if err := json.Unmarshal([]byte(string(gMessage.Message)), &noti); err != nil {
		infra.AppLog.EventError(err, "GetNotiApp json.Unmarshal", nil, message.ID)
		return noti, err
	}

	expire := noti.ActionTime.Add(5 * time.Minute)
	if time.Now().After(expire) {
		infra.AppLog.EventError(nil, "GetNotiApp expired", nil, message.ID)
		return noti, errors.New("noti-app expired")
	}

	infra.AppLog.Event("Data", map[string]interface{}{
		"slip_id":     noti.SlipID,
		"vehicle_id":  noti.VehicleID,
		"driver_id":   noti.DriverID,
		"user_id":     noti.UserID,
		"action_time": noti.ActionTime,
	}, message.ID)

	return noti, nil
}
