package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type RentalFuelRateByDay struct { // ค่าเช่าค่าน้ำมันตามระยะทาง
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	Day                     int     `db:"iday" json:"iday" gorm:"type:int"`
	RentalByVehicleTypeID01 float64 `db:"rental_by_vehicle_type_id_01" json:"rental_by_vehicle_type_id_01" gorm:"default:null; column:rental_by_vehicle_type_id_01; type:decimal(10,4); comment:'BMW' "`
	RentalByVehicleTypeID02 float64 `db:"rental_by_vehicle_type_id_02" json:"rental_by_vehicle_type_id_02" gorm:"default:null; column:rental_by_vehicle_type_id_02; type:decimal(10,4); comment:'BENZ' "`
	RentalByVehicleTypeID03 float64 `db:"rental_by_vehicle_type_id_03" json:"rental_by_vehicle_type_id_03" gorm:"default:null; column:rental_by_vehicle_type_id_03; type:decimal(10,4); comment:'Toyota Camry' "`
	RentalByVehicleTypeID04 float64 `db:"rental_by_vehicle_type_id_04" json:"rental_by_vehicle_type_id_04" gorm:"default:null; column:rental_by_vehicle_type_id_04;  type:decimal(10,4); comment:'Toyota Commuter' "`
	RentalByVehicleTypeID05 float64 `db:"rental_by_vehicle_type_id_05" json:"rental_by_vehicle_type_id_05" gorm:"default:null; column:rental_by_vehicle_type_id_05;  type:decimal(10,4); comment:'Isuzu Mu-x' "`
	RentalByVehicleTypeID06 float64 `db:"rental_by_vehicle_type_id_06" json:"rental_by_vehicle_type_id_06" gorm:"default:null; column:rental_by_vehicle_type_id_06;  type:decimal(10,4); comment:'' "`
	RentalByVehicleTypeID07 float64 `db:"rental_by_vehicle_type_id_07" json:"rental_by_vehicle_type_id_07" gorm:"default:null; column:rental_by_vehicle_type_id_07;  type:decimal(10,4); comment:'' "`
	RentalByVehicleTypeID08 float64 `db:"rental_by_vehicle_type_id_08" json:"rental_by_vehicle_type_id_08" gorm:"default:null; column:rental_by_vehicle_type_id_08;  type:decimal(10,4); comment:'' "`
	RentalByVehicleTypeID09 float64 `db:"rental_by_vehicle_type_id_09" json:"rental_by_vehicle_type_id_09" gorm:"default:null; column:rental_by_vehicle_type_id_09;  type:decimal(10,4); comment:'' "`
	RentalByVehicleTypeID10 float64 `db:"rental_by_vehicle_type_id_10" json:"rental_by_vehicle_type_id_10" gorm:"default:null; column:rental_by_vehicle_type_id_10;  type:decimal(10,4); comment:'' "`
}
