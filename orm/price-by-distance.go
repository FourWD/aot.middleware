package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type PriceByDistance struct { // poi ของ รถแต่ละประเภท
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	FleetClientID string `db:"fleet_client_id" json:"fleet_client_id" gorm:"type:varchar(10);"`

	Distance float64 `db:"distance" json:"distance" gorm:"type:decimal(16,4)"`

	ByVehicleModelID01 float64 `db:"by_vehicle_model_id_01" json:"by_vehicle_model_id_01" gorm:"default:null; column:by_vehicle_model_id_01; type:decimal(10,4); comment:'BMW' "`
	ByVehicleModelID02 float64 `db:"by_vehicle_model_id_02" json:"by_vehicle_model_id_02" gorm:"default:null; column:by_vehicle_model_id_02; type:decimal(10,4); comment:'BENZ' "`
	ByVehicleModelID03 float64 `db:"by_vehicle_model_id_03" json:"by_vehicle_model_id_03" gorm:"default:null; column:by_vehicle_model_id_03; type:decimal(10,4); comment:'Toyota Camry' "`
	ByVehicleModelID04 float64 `db:"by_vehicle_model_id_04" json:"by_vehicle_model_id_04" gorm:"default:null; column:by_vehicle_model_id_04; type:decimal(10,4); comment:'Toyota Commuter' "`
	ByVehicleModelID05 float64 `db:"by_vehicle_model_id_05" json:"by_vehicle_model_id_05" gorm:"default:null; column:by_vehicle_model_id_05; type:decimal(10,4); comment:'Isuzu Mu-x' "`
	ByVehicleModelID06 float64 `db:"by_vehicle_model_id_06" json:"by_vehicle_model_id_06" gorm:"default:null; column:by_vehicle_model_id_06; type:decimal(10,4); comment:'' "`
	ByVehicleModelID07 float64 `db:"by_vehicle_model_id_07" json:"by_vehicle_model_id_07" gorm:"default:null; column:by_vehicle_model_id_07; type:decimal(10,4); comment:'' "`
	ByVehicleModelID08 float64 `db:"by_vehicle_model_id_08" json:"by_vehicle_model_id_08" gorm:"default:null; column:by_vehicle_model_id_08; type:decimal(10,4); comment:'' "`
	ByVehicleModelID09 float64 `db:"by_vehicle_model_id_09" json:"by_vehicle_model_id_09" gorm:"default:null; column:by_vehicle_model_id_09; type:decimal(10,4); comment:'' "`
	ByVehicleModelID10 float64 `db:"by_vehicle_model_id_10" json:"by_vehicle_model_id_10" gorm:"default:null; column:by_vehicle_model_id_10; type:decimal(10,4); comment:'' "`
}
