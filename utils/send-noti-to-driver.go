package utils

import (
	"github.com/FourWD/middleware/common"
)

func SendNotiToDriver(driverID, slipID, typeID string) {
	title := "📢 งานเข้าแล้ว!"
	body := "คุณมีงานใหม่เข้าแล้ว กรุณาตรวจสอบงาน"
	topic := driverID // auto register by flutter

	logData := map[string]interface{}{
		"topic":    topic,
		"driverID": driverID,
		"slipID":   slipID,
		"typeID":   typeID,
		"body":     body,
		"title":    title,
	}
	common.Log("SendNotiToDriver", logData, "")

	data := map[string]string{
		// "event_code": "NT001",
	}

	common.SendMessageToSubscriber(topic, title, body, data)
}

func SendNotiToQueueDriver(driverID, slipID, title, body string) {
	topic := driverID

	logData := map[string]interface{}{
		"topic":    topic,
		"driverID": driverID,
		"slipID":   slipID,
		"title":    title,
		"body":     body,
	}
	common.Log("SendNotiToDriver", logData, "")

	data := map[string]string{}
	common.SendMessageToSubscriber(topic, title, body, data)
}
