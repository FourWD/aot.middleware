package utils

import (
	"fmt"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
)

func Print(id string) (string, error) {
	slip := prepareData(id)
	filepath, err := printSlip(slip)
	if err != nil {
		return "", fmt.Errorf("failed to generate PDF: %w", err)
	}
	return filepath, nil
}

func prepareData(id string) SlipPDF {
	var slip SlipPDF
	common.Database.Raw(`SELECT * FROM slips WHERE id = ?`, id).Scan(&slip)
	return slip
}

func printSlip(slip SlipPDF) (string, error) {

	filepathStr := "image/pdf/"
	fileextention := ".pdf"
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
	pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

	// err := os.MkdirAll(outputPath, os.ModePerm)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to create directory: %w", err)
	// }

	// Set font before rendering text
	pdf.SetFont("Sarabun", "", 12)

	// Uncomment and use the image path if needed
	// x := 130.00
	// y := 5.00
	// width := 30.00
	// height := 30.00
	// imagePath := "C:/Users/dd_na/OneDrive/เดสก์ท็อป/gersver/AOT/aot.limousine-service/image.png"
	// if _, err := os.Stat(imagePath); os.IsNotExist(err) {
	//     return "", fmt.Errorf("image file does not exist: %s", imagePath)
	// }
	// pdf.Image(imagePath, x, y, width, height, false, "", 0, "")

	pdf.Cell(40, 10, fmt.Sprintf("SlipNo: %s", slip.SlipNo))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("ReceiptNo: %s", slip.ReceiptNo))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("SlipVehicleTypeID: %s", slip.SlipVehicleTypeID))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("SlipVehicleSubModelID: %s", slip.SlipVehicleSubModelID))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("PaymentTypeID: %s", slip.PaymentTypeID))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("OriginPoiID: %s", slip.OriginPoiID))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("ForceOriginName: %s", slip.ForceOriginName))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("DestinationPoiID: %s", slip.DestinationPoiID))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("ForceDestinationName: %s", slip.ForceDestinationName))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("Distance: %f", slip.Distance))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("Distance *2: %f", slip.DistanceRoundTrip))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("Price: %f", slip.Price))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("Discount: %f", slip.Discount))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("NetPrice: %f", slip.NetPrice))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("Litre: %f", slip.Litre))
	pdf.Ln(5)

	filedestination := filepathStr + "Slip-Retail" + slip.ID + fileextention

	// สร้างไฟล์ PDF slip_id (ต่อด้วย id slips 123)

	errd := pdf.OutputFileAndClose(filedestination)
	if errd != nil {
		fmt.Println("Error:", errd)
	}
	return filedestination, nil
}
