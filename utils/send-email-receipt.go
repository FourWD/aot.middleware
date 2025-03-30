package utils

import (
	"log"

	midOrm "github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
	"github.com/spf13/viper"
)

func SendEmailReceipt(slipID string, recipient string) error {
	sender := viper.GetString("mail.sender")

	slip := new(midOrm.Slip)
	common.Database.Model(&midOrm.Slip{}).Where("id = ?", slipID).Scan(&slip)

	subject := "AOT Limousine [SlipNo : " + slip.SlipNo + "]"
	body := "" //AOT Limousine : " + slipID
	data, err := prepareEmailVariables(slipID)
	if err != nil {
	}

	// attachURL := "https://storage.googleapis.com/fourd-aot/receipt/da92fffe-39c7-456d-b9ac-4de20e3eb949.pdf"
	if _, err := common.SendMailWithTemplate(sender, subject, body, recipient, "invoice", data); err != nil {
	}

	return nil
}

func prepareEmailVariables(id string) (map[string]interface{}, error) {
	type slipDetail struct {
		ForceOriginName      string `json:"force_origin_name"`
		ForceDestinationName string `json:"force_destination_name"`
		VehicleSubTypeName   string `json:"vehicle_sub_type_name"`
		TravelDate           string `json:"travel_date"`
		CreditcardNo         string `json:"creditcard_no"`
		PaymentType          string `json:"payment_type"`
		PaymentDate          string `json:"payment_date"`
		Price                string `json:"price"`
		NetPrice             string `json:"net_price"`
		Discount             string `json:"discount"`
		ReceiptNo            string `json:"receipt_no"`
		SlipNo               string `json:"slip_no"`
	}

	sql := `SELECT s.force_origin_name, s.force_destination_name, vst.name AS vehicle_sub_type_name,pt.name as payment_type,s.credit_card_no as creditcard_no,
                DATE_FORMAT(s.travel_date,'%d/%m/%Y %H:%i') AS travel_date,DATE_FORMAT(s.payment_date,'%d/%m/%Y %H:%i') AS payment_date,
                FORMAT(s.price, 2) AS price,FORMAT(s.net_price, 2) AS net_price,FORMAT(s.discount, 2) AS discount,s.receipt_no,s.slip_no
                FROM slips s
                LEFT JOIN vehicle_sub_types vst ON vst.id = s.slip_vehicle_sub_type_id 
                LEFT JOIN payment_types pt on s.payment_type_id = pt.id
		WHERE s.id = ?`

	var slip slipDetail
	if err := common.Database.Raw(sql, id).Debug().Scan(&slip).Error; err != nil {
		log.Printf("failed to get slip data: %v", err)
		return nil, err
	}

	return map[string]interface{}{
		"force_origin_name":      slip.ForceOriginName,
		"force_destination_name": slip.ForceDestinationName,
		"vehicle_sub_type_name":  slip.VehicleSubTypeName,
		"travel_date":            slip.TravelDate,
		"creditcard_no":          slip.CreditcardNo,
		"payment_type":           slip.PaymentType,
		"payment_date":           slip.PaymentDate,
		"price":                  slip.Price,
		"net_price":              slip.NetPrice,
		"discount":               slip.Discount,
		"receipt_no":             slip.ReceiptNo,
		"slip_no":                slip.SlipNo,
	}, nil
}
