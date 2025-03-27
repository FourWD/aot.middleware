package utils

import (
	"github.com/FourWD/middleware/common"
)

func SendNotiToDriver(driverID, slipID, typeID string) error {
	logData := map[string]interface{}{
		"driverID": driverID,
		"slipID":   slipID,
		"typeID":   typeID,
	}

	common.Log("SendNotiToDriver", logData, "")
	data := map[string]string{
		"event_code": "NOTI",
	}

	common.SendMessageToSubscriber(driverID, "NOTI", "", data)
	return nil
}
