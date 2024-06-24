package utils

import (
	"fmt"

	"github.com/FourWD/middleware/common"
	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
)

func Print(c *fiber.Ctx) string {
	ID := c.Params("id")
	slip := prepareData(ID)
	filepath := printSlip(slip)

	return filepath

}

func prepareData(id string) SlipPDF {
	slip := new(SlipPDF)
	common.Database.Raw(`SELECT * FROM slips WHERE id = ?`, id).Scan(&slip)
	return *slip
}

func printSlip(slip SlipPDF) string {
	filepath := "C:/Users/dd_na/OneDrive/เดสก์ท็อป/gersver/AOT/"

	fileextention := ".pdf"

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("Tahoma", "B", "fonts/tahoma.ttf")
	pdf.SetFont("Tahoma", "B", 10)

	// pdf.SetFont("Tahoma", "", 12)

	x := 130.00
	y := 5.00
	width := 30.00
	height := 30.00

	pdf.Image("C:/Users/dd_na/OneDrive/เดสก์ท็อป/gersver/AOT/aot.limousine-service/image.png", x, y, width, height, false, "", 0, "")

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

	filedestination := filepath + "Slip-Retail" + slip.ID + fileextention

	// สร้างไฟล์ PDF slip_id (ต่อด้วย id slips 123)

	err := pdf.OutputFileAndClose(filedestination)
	if err != nil {
		return ""
	}

	return filedestination
}
