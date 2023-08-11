package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type PriceByHour struct { // poi ของ รถแต่ละประเภท
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Hour                   int     `db:"ihour" json:"ihour" gorm:"type:int"`
	PriceByVehicleTypeID01 float64 `db:"price_by_vehicle_type_id_01" json:"price_by_vehicle_type_id_01" gorm:"default:null; type:decimal(10,4); column:price_by_vehicle_type_id_01; comment:'BMW' "`
	PriceByVehicleTypeID02 float64 `db:"price_by_vehicle_type_id_02" json:"price_by_vehicle_type_id_02" gorm:"default:null; type:decimal(10,4); column:price_by_vehicle_type_id_02; comment:'BENZ' "`
	PriceByVehicleTypeID03 float64 `db:"price_by_vehicle_type_id_03" json:"price_by_vehicle_type_id_03" gorm:"default:null; type:decimal(10,4); column:price_by_vehicle_type_id_03; comment:'Toyota Camry' "`
	PriceByVehicleTypeID04 float64 `db:"price_by_vehicle_type_id_04" json:"price_by_vehicle_type_id_04" gorm:"default:null; type:decimal(10,4); column:price_by_vehicle_type_id_04; comment:'Toyota Commuter' "`
	PriceByVehicleTypeID05 float64 `db:"price_by_vehicle_type_id_05" json:"price_by_vehicle_type_id_05" gorm:"default:null; type:decimal(10,4); column:price_by_vehicle_type_id_05; comment:'Isuzu Mu-x' "`
	PriceByVehicleTypeID06 float64 `db:"price_by_vehicle_type_id_06" json:"price_by_vehicle_type_id_06" gorm:"default:null; type:decimal(10,4); column:price_by_vehicle_type_id_06; comment:'' "`
	PriceByVehicleTypeID07 float64 `db:"price_by_vehicle_type_id_07" json:"price_by_vehicle_type_id_07" gorm:"default:null; type:decimal(10,4); column:price_by_vehicle_type_id_07; comment:'' "`
	PriceByVehicleTypeID08 float64 `db:"price_by_vehicle_type_id_08" json:"price_by_vehicle_type_id_08" gorm:"default:null; type:decimal(10,4); column:price_by_vehicle_type_id_08; comment:'' "`
	PriceByVehicleTypeID09 float64 `db:"price_by_vehicle_type_id_09" json:"price_by_vehicle_type_id_09" gorm:"default:null; type:decimal(10,4); column:price_by_vehicle_type_id_09; comment:'' "`
	PriceByVehicleTypeID10 float64 `db:"price_by_vehicle_type_id_10" json:"price_by_vehicle_type_id_10" gorm:"default:null; type:decimal(10,4); column:price_by_vehicle_type_id_10; comment:'' "`
}
