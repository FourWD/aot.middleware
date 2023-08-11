package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type RentalFuelRateByDistance struct { // ค่าเช่าค่าน้ำมันตามระยะทาง
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	Hour              int     `db:"ihour" json:"ihour" gorm:"type:int"`
	ByVehicleTypeID01 float64 `db:"by_vehicle_type_id_01" json:"by_vehicle_type_id_01" gorm:"default:null; column:by_vehicle_type_id_01; type:decimal(10,4); comment:'BMW' "`
	ByVehicleTypeID02 float64 `db:"by_vehicle_type_id_02" json:"by_vehicle_type_id_02" gorm:"default:null; column:by_vehicle_type_id_02; type:decimal(10,4); comment:'BENZ' "`
	ByVehicleTypeID03 float64 `db:"by_vehicle_type_id_03" json:"by_vehicle_type_id_03" gorm:"default:null; column:by_vehicle_type_id_03; type:decimal(10,4); comment:'Toyota Camry' "`
	ByVehicleTypeID04 float64 `db:"by_vehicle_type_id_04" json:"by_vehicle_type_id_04" gorm:"default:null; column:by_vehicle_type_id_04;  type:decimal(10,4); comment:'Toyota Commuter' "`
	ByVehicleTypeID05 float64 `db:"by_vehicle_type_id_05" json:"by_vehicle_type_id_05" gorm:"default:null; column:by_vehicle_type_id_05;  type:decimal(10,4); comment:'Isuzu Mu-x' "`
	ByVehicleTypeID06 float64 `db:"by_vehicle_type_id_06" json:"by_vehicle_type_id_06" gorm:"default:null; column:by_vehicle_type_id_06;  type:decimal(10,4); comment:'' "`
	ByVehicleTypeID07 float64 `db:"by_vehicle_type_id_07" json:"by_vehicle_type_id_07" gorm:"default:null; column:by_vehicle_type_id_07;  type:decimal(10,4); comment:'' "`
	ByVehicleTypeID08 float64 `db:"by_vehicle_type_id_08" json:"by_vehicle_type_id_08" gorm:"default:null; column:by_vehicle_type_id_08;  type:decimal(10,4); comment:'' "`
	ByVehicleTypeID09 float64 `db:"by_vehicle_type_id_09" json:"by_vehicle_type_id_09" gorm:"default:null; column:by_vehicle_type_id_09;  type:decimal(10,4); comment:'' "`
	ByVehicleTypeID10 float64 `db:"by_vehicle_type_id_10" json:"by_vehicle_type_id_10" gorm:"default:null; column:by_vehicle_type_id_10;  type:decimal(10,4); comment:'' "`
}
