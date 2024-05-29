package calculatorHandler

import (
	"path"
	"strconv"

	"github.com/FourWD/middleware/common"

	"github.com/gofiber/fiber/v2"
)

func GetImage(c *fiber.Ctx) error {
	imagePath := c.Params("imagepath")

	imageFullPath := path.Join("resource/image", imagePath)

	return c.SendFile(imageFullPath)
}
func GetPrice(c *fiber.Ctx) error {

	fleetClientID := c.Params("fleet_client_id")
	distance := c.Params("distance")
	common.Print(fleetClientID, distance)

	cdnimage := "https://storage.googleapis.com/fourd-aot/images/"
	var data Response
	var discount = 0.00
	f, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		return common.FiberError(c, "4000", "errror")
	}

	data.Distance = f

	var tempVehicle []Vehicle
	common.Database.Raw(`SELECT vs.name_en as vehicle_source_name, vm.id as vehicle_model_id, vm.name as vehicle_model_name, vm.name_en as vehicle_model_name_en, vsm.id as vehicle_sub_model_id,
	vsm.name_en as vehicle_sub_model_name, vsm.sub_model_group, vsm.vehicle_source_id ,vsm.image as vehicle_image, vsm.row_order 
	from vehicle_sub_models vsm  
	left join vehicle_sources vs on vsm.vehicle_source_id = vs.id
	left join vehicle_models vm on vsm.vehicle_model_id = vm.id 
	where is_show = 1 ORDER BY vsm.row_order`).Scan(&tempVehicle)

	var getprice = getDistance(fleetClientID, f)

	var totalPrice float64
	var vehicles []Vehicle

	for _, v := range tempVehicle {
		switch v.VehicleModelID {
		case "01":
			v.Price = getprice.ByVehicleModelID01
			v.Discount = discount
			v.NetPrice = getprice.ByVehicleModelID01 - discount
			v.VehicleImage = cdnimage + v.VehicleImage
		case "02":
			v.Price = getprice.ByVehicleModelID02
			v.Discount = discount
			v.NetPrice = getprice.ByVehicleModelID02 - discount
			v.VehicleImage = cdnimage + v.VehicleImage
		case "03":
			v.Price = getprice.ByVehicleModelID03
			v.Discount = discount
			v.NetPrice = getprice.ByVehicleModelID03 - discount
			v.VehicleImage = cdnimage + v.VehicleImage

		case "04":
			v.Price = getprice.ByVehicleModelID04
			v.Discount = discount
			v.NetPrice = getprice.ByVehicleModelID04 - discount
			v.VehicleImage = cdnimage + v.VehicleImage

		case "05":
			v.Price = getprice.ByVehicleModelID05
			v.Discount = discount
			v.NetPrice = getprice.ByVehicleModelID05 - discount
			v.VehicleImage = cdnimage + v.VehicleImage

		}
		totalPrice += v.Price
		vehicles = append(vehicles, Vehicle{
			VehicleModelID:      v.VehicleModelID,
			VehicleSourceName:   v.VehicleSourceName,
			VehicleModelName:    v.VehicleModelName,
			VehicleModelNameEn:  v.VehicleModelNameEn,
			VehicleSubModelName: v.VehicleSubModelName,
			VehicleSubModelID:   v.VehicleSubModelID,
			VehicleSourceID:     v.VehicleSourceID,
			Price:               v.Price,
			Discount:            v.Discount,
			VehicleImage:        v.VehicleImage,
			NetPrice:            v.NetPrice,
		})
	}

	data.Vehicles = vehicles

	return c.JSON(fiber.Map{"status": 1, "message": "success", "data": data})
}
