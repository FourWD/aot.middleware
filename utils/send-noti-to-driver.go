package utils

// import (
// 	"time"

// 	"github.com/FourWD/middleware/common"
// )

// func SendNotiToDriver(driverID, slipID, typeID string) error {
// 	title := "📢 งานเข้าแล้ว!"
// 	body := "คุณมีงานใหม่เข้าแล้ว กรุณาตรวจสอบงาน"
// 	topic := driverID // auto register by flutter

// 	logData := map[string]interface{}{
// 		"topic":    topic,
// 		"driverID": driverID,
// 		"slipID":   slipID,
// 		"typeID":   typeID,
// 		"body":     body,
// 		"title":    title,
// 	}
// 	common.Log("SendNotiToDriver", logData, "")

// 	data := map[string]string{
// 		"event_code": "D001",
// 		"expire":     time.Now().Add(5 * time.Minute).Format(common.DATE_FORMAT_MINUTE),
// 	}

// 	return common.SendMessageToSubscriber(topic, title, body, data)
// }

// func SendNotiToQueueDriver(driverID, slipID, typeID string, index int) {
// 	title := "📢 แจ้งเตือนคิวของคุณ"
// 	var body string

// 	if index == 0 {
// 		body = "ถึงคิวของคุณแล้ว"
// 	} else {
// 		body = fmt.Sprintf("อีก %d คิวจะถึงคุณ", index)
// 	}

// 	topic := driverID

// 	logData := map[string]interface{}{
// 		"topic":    topic,
// 		"driverID": driverID,
// 		"slipID":   slipID,
// 		"typeID":   typeID,
// 		"title":    title,
// 		"body":     body,
// 	}
// 	common.Log("SendNotiToDriver", logData, "")

// 	data := map[string]string{}
// 	common.SendMessageToSubscriber(topic, title, body, data)
// }
