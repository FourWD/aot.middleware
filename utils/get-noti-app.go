package utils

import (
	"encoding/json"
	"errors"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/FourWD/aot.middleware/model"
	"github.com/FourWD/middleware/common"
)

func GetNotiApp(message *pubsub.Message) (model.NotiApp, error) {
	common.Log("GetNotiApp Start", map[string]interface{}{
		"message.id":   message.ID,
		"message.data": string(message.Data),
	}, message.ID)

	gMessage := common.ConventStringToGoogleMessage(string(message.Data))

	var noti model.NotiApp
	if err := json.Unmarshal([]byte(string(gMessage.Message)), &noti); err != nil {
		common.LogError("GetNotiApp json.Unmarshal error: "+err.Error(), nil, message.ID)
		return noti, err
	}

	expire := noti.ActionTime.Add(5 * time.Minute)
	if time.Now().After(expire) {
		common.LogError("GetNotiApp expired", nil, message.ID)
		return noti, errors.New("noti-app expired")
	}

	common.Log("Data", map[string]interface{}{
		"slip_id":     noti.SlipID,
		"vehicle_id":  noti.VehicleID,
		"driver_id":   noti.DriverID,
		"user_id":     noti.UserID,
		"action_time": noti.ActionTime,
	}, message.ID)

	return noti, nil
}
