package utils

import (
	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func GetPrice(systemName string, distance float64) orm.PriceByDistance {
	var getprice orm.PriceByDistance

	var sql = `SELECT * FROM price_by_distances
	    WHERE fleet_client_id = ? AND distance >= ? ORDER BY distance LIMIT 1`
	common.Database.Raw(sql, systemName, distance).Scan(&getprice)
	return getprice
}

// func getSlipVehicle() {
// 	var getVehicle orm.PriceByDistance

// 	var sql = `SELECT * FROM price_by_distances
// 	    WHERE fleet_client_id = ? AND distance >= ? ORDER BY distance LIMIT 1`
// 	common.Database.Raw(sql, fleetClientID, distance).Scan(&getprice)
// 	common.Print(sql, "error")
// 	return getprice
// }
